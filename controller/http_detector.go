// Copyright 2020 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"context"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/thoas/go-funk"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
)

type httpDetectorCtrl struct{}

type (
	httpDetectorAddParams struct {
		detectorAddParams

		IPS     []string `json:"ips" validate:"required,dive,ip"`
		Proxies []string `json:"proxies" validate:"required,dive,uri" `
		URL     string   `json:"url" validate:"required,xHTTP"`
		// 检测脚本
		Script string `json:"script" validate:"omitempty"`
	}
	httpDetectorListParams struct {
		listParams

		account string
	}
	httpDetectorUpdateParams struct {
		detectorUpdateParams

		account string

		IPS     []string `json:"ips" validate:"omitempty,dive,ip"`
		Proxies []string `json:"proxies" validate:"omitempty,dive,uri" `
		URL     string   `json:"url" validate:"omitempty,xHTTP"`
		// 检测脚本
		Script string `json:"script" validate:"omitempty"`
	}

	httpDetectorResultListParams struct {
		detectorListResultParams
	}
)

type (
	// httpDetectorListResp response of list http
	httpDetectorListResp struct {
		HTTPDetectors []*ent.HTTPDetector `json:"httpDetectors"`
		Count         int                 `json:"count"`
	}
	httpDetectorResultListResp struct {
		HTTPDetectorResults []*ent.HTTPDetectorResult `json:"httpDetectorResults"`
		Count               int                       `json:"count"`
	}
)

func init() {
	prefix := "/http-detectors"
	g := router.NewGroup(
		prefix,
		loadUserSession,
		shouldBeLogin,
	)
	ctrl := httpDetectorCtrl{}

	// 添加配置
	g.POST(
		"/v1",
		newTrackerMiddleware(cs.ActionDetectorHTTPAdd),
		ctrl.add,
	)
	// 查询配置
	g.GET(
		"/v1",
		ctrl.list,
	)
	g.GET(
		"/v1/{id}",
		ctrl.findByID,
	)
	g.PATCH(
		"/v1/{id}",
		newTrackerMiddleware(cs.ActionDetectorHTTPUpdate),
		ctrl.updateByID,
	)
	g.GET(
		"/results/v1",
		ctrl.listResult,
	)
}

func getHTTPDetectorClient() *ent.HTTPDetectorClient {
	return helper.EntGetClient().HTTPDetector
}
func getHTTPDetectorResultClient() *ent.HTTPDetectorResultClient {
	return helper.EntGetClient().HTTPDetectorResult
}

func (addParams *httpDetectorAddParams) save(ctx context.Context) (*ent.HTTPDetector, error) {
	return getHTTPDetectorClient().Create().
		SetStatus(addParams.Status).
		SetName(addParams.Name).
		SetOwners(addParams.Owners).
		SetReceivers(addParams.Receivers).
		SetTimeout(addParams.Timeout).
		SetDescription(addParams.Description).
		SetIps(addParams.IPS).
		SetURL(addParams.URL).
		SetScript(addParams.Script).
		SetProxies(addParams.Proxies).
		Save(ctx)
}

func (listParams *httpDetectorListParams) where(query *ent.HTTPDetectorQuery) {
	account := listParams.account
	if account != "" {
		query.Where(predicate.HTTPDetector(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(httpdetector.FieldOwners, account))
		}))
	}
}

func (listParams *httpDetectorListParams) count(ctx context.Context) (int, error) {
	query := getHTTPDetectorClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *httpDetectorListParams) queryAll(ctx context.Context) ([]*ent.HTTPDetector, error) {
	query := getHTTPDetectorClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

// where http detector result where
func (listParams *httpDetectorResultListParams) where(query *ent.HTTPDetectorResultQuery) {
	tasks := listParams.doTaskFilter()
	if len(tasks) != 0 {
		query.Where(httpdetectorresult.TaskIn(tasks...))
	}
	result := listParams.Result
	if result != 0 {
		query.Where(httpdetectorresult.Result(schema.DetectorResult(result)))
	}
	ms := listParams.GetDurationMillSecond()
	if ms > 0 {
		query.Where(httpdetectorresult.MaxDurationGTE(ms))
	}
	startedAt := listParams.StartedAt
	if !startedAt.IsZero() {
		query.Where(httpdetectorresult.CreatedAtGTE(startedAt))
	}
	endedAt := listParams.EndedAt
	if !endedAt.IsZero() {
		query.Where(httpdetectorresult.CreatedAtLTE(endedAt))
	}
}

func (listParams *httpDetectorResultListParams) count(ctx context.Context) (int, error) {
	query := getHTTPDetectorResultClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *httpDetectorResultListParams) queryAll(ctx context.Context) ([]*ent.HTTPDetectorResult, error) {
	query := getHTTPDetectorResultClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (updateParams *httpDetectorUpdateParams) updateByID(ctx context.Context, id int) (*ent.HTTPDetector, error) {
	// 指定账号则判断是否是该配置的owner
	account := updateParams.account
	if account != "" {
		result, err := getHTTPDetectorClient().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if !util.ContainsString(result.Owners, account) {
			return nil, errInvalidUser
		}
	}
	updateOne := getHTTPDetectorClient().UpdateOneID(id)
	if updateParams.Name != "" {
		updateOne.SetName(updateParams.Name)
	}
	if updateParams.Status != 0 {
		updateOne.SetStatus(updateParams.Status)
	}
	if updateParams.Description != "" {
		updateOne.SetDescription(updateParams.Description)
	}
	if len(updateParams.Receivers) != 0 {
		updateOne.SetReceivers(updateParams.Receivers)
	}
	if updateParams.Timeout != "" {
		updateOne.SetTimeout(updateParams.Timeout)
	}
	// 允许直接修改owner
	if len(updateParams.Owners) != 0 {
		updateOne.SetOwners(updateParams.Owners)
	}

	if len(updateParams.IPS) != 0 {
		updateOne.SetIps(updateParams.IPS)
	}
	if len(updateParams.Proxies) != 0 {
		updateOne.SetProxies(updateParams.Proxies)
	}
	if updateParams.URL != "" {
		updateOne.SetURL(updateParams.URL)
	}
	if updateParams.Script != "" {
		updateOne.SetScript(strings.TrimSpace(updateParams.Script))
	}

	return updateOne.Save(ctx)
}

// 添加http检测
func (*httpDetectorCtrl) add(c *elton.Context) error {
	params := httpDetectorAddParams{}
	err := validateBody(c, &params)
	if err != nil {
		return err
	}
	result, err := params.save(c.Context())
	if err != nil {
		return err
	}
	c.Created(result)
	return nil
}

func (*httpDetectorCtrl) list(c *elton.Context) error {
	params := httpDetectorListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	us := getUserSession(c)
	// 如果非管理员，则指定当前账号
	if !us.IsAdmin() {
		params.account = us.MustGetInfo().Account
	}
	resp := httpDetectorListResp{
		Count: -1,
	}
	if params.ShouldCount() {
		count, err := params.count(c.Context())
		if err != nil {
			return err
		}
		resp.Count = count
	}
	result, err := params.queryAll(c.Context())
	if err != nil {
		return err
	}
	resp.HTTPDetectors = result
	c.Body = &resp
	return nil
}

func (*httpDetectorCtrl) updateByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	params := httpDetectorUpdateParams{}
	err = validateBody(c, &params)
	if err != nil {
		return err
	}
	us := getUserSession(c)
	// 如果非管理员，则指定当前账号
	if !us.IsAdmin() {
		params.account = us.MustGetInfo().Account
	}
	result, err := params.updateByID(c.Context(), id)
	if err != nil {
		return err
	}
	c.Body = result
	return nil
}

func (*httpDetectorCtrl) findByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	result, err := getHTTPDetectorClient().Get(c.Context(), id)
	if err != nil {
		return err
	}
	c.Body = result
	return nil
}

func (*httpDetectorCtrl) listResult(c *elton.Context) error {
	params := httpDetectorResultListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	resp := httpDetectorResultListResp{
		Count: -1,
	}
	params.Tasks, err = getDetectorTasksByReceiver(
		c.Context(),
		detectorCategoryHTTP,
		getUserSession(c),
	)
	if err == errTaskNotFound {
		c.Body = &resp
		return nil
	}
	if err != nil {
		return err
	}

	if params.ShouldCount() {
		count, err := params.count(c.Context())
		if err != nil {
			return err
		}
		resp.Count = count
	}
	result, err := params.queryAll(c.Context())
	if err != nil {
		return err
	}

	// 填充检测任务名称
	idList := make([]int, 0)
	for _, item := range result {
		if funk.ContainsInt(idList, item.Task) {
			continue
		}
		idList = append(idList, item.Task)
	}
	detectors, err := getHTTPDetectorClient().Query().
		Where(
			httpdetector.IDIn(idList...),
		).
		Select("id", "name").
		All(c.Context())
	if err != nil {
		return err
	}
	for _, item := range result {
		for _, d := range detectors {
			if item.Task == d.ID {
				item.TaskName = d.Name
			}
		}
	}

	resp.HTTPDetectorResults = result
	c.Body = &resp
	return nil
}
