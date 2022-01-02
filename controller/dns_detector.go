// Copyright 2021 tree xie
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
	"github.com/vicanso/cybertect/ent/dnsdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
)

type dnsDetectorCtrl struct{}

type (
	dnsDetectorAddParams struct {
		detectorAddParams

		Host    string   `json:"host" validate:"required,hostname"`
		IPS     []string `json:"ips" validate:"required,dive,ip"`
		Servers []string `json:"servers" validate:"required,dive,ip"`
	}

	dnsDetectorListParams struct {
		listParams

		account string
	}

	dnsDetectorUpdateParams struct {
		detectorUpdateParams

		account string

		Host    string   `json:"host" validate:"omitempty,hostname"`
		IPS     []string `json:"ips" validate:"omitempty,dive,ip"`
		Servers []string `json:"servers" validate:"omitempty,dive,ip"`
	}

	dnsDetectorResultListParams struct {
		detectorListResultParams
	}
)

type (
	dnsDetectorListResp struct {
		DNSDetectors []*ent.DNSDetector `json:"dnsDetectors"`
		Count        int                `json:"count"`
	}
	dnsDetectorResultListResp struct {
		DNSDetectorResults []*ent.DNSDetectorResult `json:"dnsDetectorResults"`
		Count              int                      `json:"count"`
	}
)

func init() {
	prefix := "/dns-detectors"
	g := router.NewGroup(
		prefix,
		loadUserSession,
		shouldBeLogin,
	)
	ctrl := dnsDetectorCtrl{}

	g.POST(
		"/v1",
		newTrackerMiddleware(cs.ActionDetectorDNSAdd),
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
		newTrackerMiddleware(cs.ActionDetectorDNSUpdate),
		ctrl.updateByID,
	)
	g.GET(
		"/v1/results",
		ctrl.listResult,
	)
}

func getDNSDetectorClient() *ent.DNSDetectorClient {
	return helper.EntGetClient().DNSDetector
}

func getDNSDetectorResultClient() *ent.DNSDetectorResultClient {
	return helper.EntGetClient().DNSDetectorResult
}

func (addParams *dnsDetectorAddParams) save(ctx context.Context) (*ent.DNSDetector, error) {
	return getDNSDetectorClient().Create().
		SetStatus(addParams.Status).
		SetName(addParams.Name).
		SetOwners(addParams.Owners).
		SetReceivers(addParams.Receivers).
		SetTimeout(addParams.Timeout).
		SetDescription(addParams.Description).
		SetHost(addParams.Host).
		SetIps(addParams.IPS).
		SetServers(addParams.Servers).
		SetInterval(addParams.Interval).
		Save(ctx)
}

func (listParams *dnsDetectorListParams) where(query *ent.DNSDetectorQuery) {
	if listParams.account != "" {
		query.Where(predicate.DNSDetector(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(dnsdetector.FieldOwners, listParams.account))
		}))
	}
}

func (listParams *dnsDetectorListParams) count(ctx context.Context) (int, error) {
	query := getDNSDetectorClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *dnsDetectorListParams) queryAll(ctx context.Context) ([]*ent.DNSDetector, error) {
	query := getDNSDetectorClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (listParams *dnsDetectorResultListParams) where(query *ent.DNSDetectorResultQuery) {
	tasks := listParams.doTaskFilter()
	if len(tasks) != 0 {
		query.Where(dnsdetectorresult.TaskIn(tasks...))
	}
	result := listParams.Result
	if result != 0 {
		query.Where(dnsdetectorresult.Result(schema.DetectorResult(result)))
	}
	ms := listParams.GetDurationMillSecond()
	if ms > 0 {
		query.Where(dnsdetectorresult.MaxDurationGTE(ms))
	}
	startedAt := listParams.StartedAt
	if !startedAt.IsZero() {
		query.Where(dnsdetectorresult.CreatedAtGTE(startedAt))
	}
	endedAt := listParams.EndedAt
	if !endedAt.IsZero() {
		query.Where(dnsdetectorresult.CreatedAtLTE(endedAt))
	}
}

func (listParams *dnsDetectorResultListParams) count(ctx context.Context) (int, error) {
	query := getDNSDetectorResultClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *dnsDetectorResultListParams) queryAll(ctx context.Context) (ent.DNSDetectorResults, error) {
	query := getDNSDetectorResultClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	fields := listParams.GetFields()
	if len(fields) != 0 {
		results := make(ent.DNSDetectorResults, 0)
		err := query.Select(fields...).Scan(ctx, &results)
		return results, err
	}
	return query.All(ctx)
}

func (updateParams *dnsDetectorUpdateParams) updateByID(ctx context.Context, id int) (*ent.DNSDetector, error) {
	account := updateParams.account
	if account != "" {
		result, err := getDNSDetectorClient().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if !util.ContainsString(result.Owners, account) {
			return nil, errInvalidUser
		}
	}
	updateOne := getDNSDetectorClient().UpdateOneID(id)
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

	if updateParams.Host != "" {
		updateOne.SetHost(updateParams.Host)
	}
	if len(updateParams.IPS) != 0 {
		updateOne.SetIps(updateParams.IPS)
	}
	if len(updateParams.Servers) != 0 {
		updateOne.SetServers(updateParams.Servers)
	}

	interval := updateParams.Interval
	if interval != "" {
		if interval == "0s" {
			interval = ""
		}
		updateOne.SetInterval(interval)
	}

	return updateOne.Save(ctx)
}

// 添加dns检测
func (*dnsDetectorCtrl) add(c *elton.Context) error {
	params := dnsDetectorAddParams{}
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

func (*dnsDetectorCtrl) list(c *elton.Context) error {
	params := dnsDetectorListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	us := getUserSession(c)
	// 非管理员，则指定当前账号
	if !us.IsAdmin() {
		params.account = us.MustGetInfo().Account
	}
	resp := dnsDetectorListResp{
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
	resp.DNSDetectors = result
	c.Body = &resp

	return nil
}

func (*dnsDetectorCtrl) updateByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}

	params := dnsDetectorUpdateParams{}
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

func (*dnsDetectorCtrl) findByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	result, err := getDNSDetectorClient().Get(c.Context(), id)
	if err != nil {
		return err
	}
	c.Body = result
	return nil
}

func (*dnsDetectorCtrl) listResult(c *elton.Context) error {
	params := dnsDetectorResultListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	resp := dnsDetectorResultListResp{
		Count: -1,
	}
	params.Tasks, err = getDetectorTasksByReceiver(
		c.Context(),
		detectorCategoryDNS,
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
	detectors, err := getDNSDetectorClient().Query().
		Where(
			dnsdetector.IDIn(idList...),
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

	resp.DNSDetectorResults = result
	c.Body = &resp
	return nil
}
