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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/thoas/go-funk"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/dnsdetector"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/tcpdetector"
	"github.com/vicanso/cybertect/ent/tcpdetectorresult"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
)

type tcpDetectorCtrl struct{}

type (
	tcpDetectorAddParams struct {
		detectorAddParams

		Addrs []string `json:"addrs" validate:"required,dive,xHostPort"`
	}
	tcpDetectorListParams struct {
		listParams

		account string
	}
	tcpDetectorUpdateParams struct {
		detectorUpdateParams

		account string
		Addrs   []string `json:"addrs" validate:"omitempty,dive,xHostPort"`
	}

	tcpDetectorResultListParams struct {
		detectorListResultParams
	}
)

type (
	tcpDetectorListResp struct {
		TCPDetectors []*ent.TCPDetector `json:"tcpDetectors"`
		Count        int                `json:"count"`
	}
	tcpDetectorResultListResp struct {
		TCPDetectorResults []*ent.TCPDetectorResult `json:"tcpDetectorResults"`
		Count              int                      `json:"count"`
	}
)

func init() {
	prefix := "/tcp-detectors"
	g := router.NewGroup(
		prefix,
		loadUserSession,
		shouldBeLogin,
	)
	ctrl := tcpDetectorCtrl{}

	// 添加配置
	g.POST(
		"/v1",
		newTrackerMiddleware(cs.ActionDetectorTCPAdd),
		ctrl.add,
	)
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
		newTrackerMiddleware(cs.ActionDetectorTCPUpdate),
		ctrl.updateByID,
	)
	g.GET(
		"/results/v1",
		ctrl.listResult,
	)
}

func getTCPDetectorClient() *ent.TCPDetectorClient {
	return helper.EntGetClient().TCPDetector
}

func getTCPDetectorResultClient() *ent.TCPDetectorResultClient {
	return helper.EntGetClient().TCPDetectorResult
}

func (addParams *tcpDetectorAddParams) save(ctx context.Context) (*ent.TCPDetector, error) {
	return getTCPDetectorClient().Create().
		SetStatus(addParams.Status).
		SetName(addParams.Name).
		SetOwners(addParams.Owners).
		SetReceivers(addParams.Receivers).
		SetTimeout(addParams.Timeout).
		SetDescription(addParams.Description).
		SetAddrs(addParams.Addrs).
		Save(ctx)
}

func (listParams *tcpDetectorListParams) where(query *ent.TCPDetectorQuery) {
	account := listParams.account
	if account != "" {
		query.Where(predicate.TCPDetector(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(dnsdetector.FieldOwners, account))
		}))
	}
}

func (listParams *tcpDetectorListParams) count(ctx context.Context) (int, error) {
	query := getTCPDetectorClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *tcpDetectorListParams) queryAll(ctx context.Context) ([]*ent.TCPDetector, error) {
	query := getTCPDetectorClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (listParams *tcpDetectorResultListParams) where(query *ent.TCPDetectorResultQuery) {
	tasks := listParams.Tasks
	if len(tasks) != 0 {
		query.Where(tcpdetectorresult.TaskIn(tasks...))
	}
	result := listParams.Result
	if result != 0 {
		query.Where(tcpdetectorresult.Result(schema.DetectorResult(result)))
	}
	ms := listParams.GetDurationMillSecond()
	if ms > 0 {
		query.Where(tcpdetectorresult.MaxDurationGTE(ms))
	}
	startedAt := listParams.StartedAt
	if !startedAt.IsZero() {
		query.Where(tcpdetectorresult.CreatedAtGTE(startedAt))
	}
	endedAt := listParams.EndedAt
	if !endedAt.IsZero() {
		query.Where(tcpdetectorresult.CreatedAtLTE(endedAt))
	}
}

func (listParams *tcpDetectorResultListParams) count(ctx context.Context) (int, error) {
	query := getTCPDetectorResultClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *tcpDetectorResultListParams) queryAll(ctx context.Context) ([]*ent.TCPDetectorResult, error) {
	query := getTCPDetectorResultClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (updateParams *tcpDetectorUpdateParams) updateByID(ctx context.Context, id int) (*ent.TCPDetector, error) {
	account := updateParams.account
	if account != "" {
		result, err := getTCPDetectorClient().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if !util.ContainsString(result.Owners, account) {
			return nil, errInvalidUser
		}
	}
	updateOne := getTCPDetectorClient().UpdateOneID(id)
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

	if len(updateParams.Addrs) != 0 {
		updateOne.SetAddrs(updateParams.Addrs)
	}
	return updateOne.Save(ctx)
}

func (*tcpDetectorCtrl) add(c *elton.Context) error {
	params := tcpDetectorAddParams{}
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

func (*tcpDetectorCtrl) list(c *elton.Context) error {
	params := tcpDetectorListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	us := getUserSession(c)
	// 非管理员，只能查询所拥有的配置
	if !us.IsAdmin() {
		params.account = us.MustGetInfo().Account
	}
	resp := tcpDetectorListResp{
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
	resp.TCPDetectors = result
	c.Body = &resp
	return nil
}

func (*tcpDetectorCtrl) updateByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	params := tcpDetectorUpdateParams{}
	err = validateBody(c, &params)
	if err != nil {
		return err
	}
	us := getUserSession(c)
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

func (*tcpDetectorCtrl) findByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	result, err := getTCPDetectorClient().Get(c.Context(), id)
	if err != nil {
		return err
	}
	c.Body = result
	return nil
}

func (*tcpDetectorCtrl) listResult(c *elton.Context) error {
	params := tcpDetectorResultListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	resp := tcpDetectorResultListResp{
		Count: -1,
	}
	params.Tasks, err = getDetectorTasksByReceiver(
		c.Context(),
		detectorCategoryTCP,
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
	detectors, err := getTCPDetectorClient().Query().
		Where(
			tcpdetector.IDIn(idList...),
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

	resp.TCPDetectorResults = result
	c.Body = &resp

	return nil
}
