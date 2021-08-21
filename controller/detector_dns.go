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
	"strconv"
	"time"

	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/dnsdetector"
	"github.com/vicanso/cybertect/ent/dnsdetectorresult"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
)

type (
	detectorDNSCtrl struct{}
	// detectorAddDNSParams params of add dns
	detectorAddDNSParams struct {
		detectorAddParams

		Host    string   `json:"host,omitempty" validate:"required,hostname"`
		IPS     []string `json:"ips,omitempty" validate:"required,dive,ip"`
		Servers []string `json:"servers,omitempty" validate:"required,dive,ip"`
	}
	// detectorUpdateDNSParams params of update dns
	detectorUpdateDNSParams struct {
		detectorUpdateParams

		Host    string   `json:"host,omitempty" validate:"omitempty,hostname"`
		IPS     []string `json:"ips,omitempty" validate:"omitempty,dive,ip"`
		Servers []string `json:"servers,omitempty" validate:"omitempty,dive,ip"`
	}
	// detectorListDNSParams params of list dns
	detectorListDNSParams struct {
		listParams

		Owner string
	}
	// detectorDNSListResp response of list dns
	detectorListDNSResp struct {
		DNSES []*ent.DNSDetector `json:"dnses,omitempty"`
		Count int                `json:"count,omitempty"`
	}

	// detectorListDNSResultParams params of list dns result
	detectorListDNSResultParams struct {
		detectorListResultParams
	}
	// detectorListDNSResultResp response of list dns result
	detectorListDNSResultResp struct {
		DNSResults   []*ent.DNSDetectorResult `json:"dnsResults,omitempty"`
		DNSDetectors []*ent.DNSDetector       `json:"dnsDetectors,omitempty"`
		Count        int                      `json:"count,omitempty"`
	}

	// detectorFilterDNSParams params of filter dns
	detectorFilterDNSParams struct {
		Keyword string `json:"keyword,omitempty" validate:"required,xKeyword"`
	}
	// detectorFilterDNSResp response of filter dns
	detectorFilterDNSResp struct {
		DNSDetectors []*ent.DNSDetector `json:"dnsDetectors,omitempty"`
	}
)

func init() {
	prefix := "/detectors/v1/dnses"
	g := router.NewGroup(prefix, loadUserSession, shouldBeLogin)
	nsg := router.NewGroup(prefix)

	ctrl := detectorDNSCtrl{}

	// 查询dns配置
	g.GET(
		"",
		ctrl.list,
	)
	// 添加dns配置
	g.POST(
		"",
		newTrackerMiddleware(cs.ActionDetectorDNSAdd),
		ctrl.add,
	)
	// 更新dns配置
	g.PATCH(
		"/{id}",
		newTrackerMiddleware(cs.ActionDetectorDNSUpdate),
		ctrl.updateByID,
	)

	// dns检测配置筛选
	nsg.GET(
		"/filter",
		ctrl.filter,
	)

	// 查询dns检测结果
	nsg.GET(
		"/results",
		ctrl.listResult,
	)
	// 查询dns检测结果详情
	nsg.GET(
		"/results/{id}",
		ctrl.getResult,
	)
}

// save dns save
func (params *detectorAddDNSParams) save(ctx context.Context, owner string) (result *ent.DNSDetector, err error) {
	return getEntClient().DNSDetector.Create().
		SetName(params.Name).
		SetStatus(schema.Status(params.Status)).
		SetDescription(params.Description).
		SetReceivers(params.Receivers).
		SetTimeout(params.Timeout).
		SetHost(params.Host).
		SetIps(params.IPS).
		SetServers(params.Servers).
		SetOwner(owner).
		Save(ctx)
}

// where dns where
func (params *detectorListDNSParams) where(query *ent.DNSDetectorQuery) {
	if params.Owner != "" {
		query.Where(dnsdetector.OwnerEQ(params.Owner))
	}
}

// queryAll query all dns detector
func (params *detectorListDNSParams) queryAll(ctx context.Context) (dnses []*ent.DNSDetector, err error) {
	query := getEntClient().DNSDetector.Query()
	query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	params.where(query)
	return query.All(ctx)
}

// count count dns detector
func (params *detectorListDNSParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().DNSDetector.Query()
	params.where(query)
	return query.Count(ctx)
}

// updateByID update dns detector by id
func (params *detectorUpdateDNSParams) updateByID(ctx context.Context, id int) (result *ent.DNSDetector, err error) {
	currentDNS, err := getEntClient().DNSDetector.Get(ctx, id)
	if err != nil {
		return
	}
	if currentDNS.Owner != params.CurrentUser {
		err = errInvalidUser.Clone()
		return
	}

	updateOne := getEntClient().DNSDetector.UpdateOneID(id)

	if params.Name != "" {
		updateOne.SetName(params.Name)
	}
	if params.Status != 0 {
		updateOne.SetStatus(schema.Status(params.Status))
	}
	if params.Description != "" {
		updateOne.SetDescription(params.Description)
	}
	if len(params.Receivers) != 0 {
		updateOne.SetReceivers(params.Receivers)
	}
	if params.Timeout != "" {
		updateOne.SetTimeout(params.Timeout)
	}

	if params.Host != "" {
		updateOne.SetHost(params.Host)
	}
	if len(params.IPS) != 0 {
		updateOne.SetIps(params.IPS)
	}
	if len(params.Servers) != 0 {
		updateOne.SetServers(params.Servers)
	}
	return updateOne.Save(ctx)
}

// where dns detector result where
func (params *detectorListDNSResultParams) where(query *ent.DNSDetectorResultQuery) {
	task := params.GetTaskID()
	if task != 0 {
		query.Where(dnsdetectorresult.Task(task))
	}

	result := params.GetResult()
	if result != 0 {
		query.Where(dnsdetectorresult.Result(result))
	}

	ms := params.GetDurationMillSecond()
	if ms > 0 {
		query.Where(dnsdetectorresult.MaxDurationGTE(ms))
	}
	createdAtRange := params.GetCreatedAtGT()
	if !createdAtRange.IsZero() {
		query.Where(dnsdetectorresult.CreatedAtGTE(createdAtRange))
	}
}

// queryAll query all dns result
func (params *detectorListDNSResultParams) queryAll(ctx context.Context) (dnsResults []*ent.DNSDetectorResult, err error) {
	query := getEntClient().DNSDetectorResult.Query()
	query = query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	params.where(query)
	fields := params.GetFields()
	if len(fields) == 0 {
		return query.All(ctx)
	}
	newFiles := make([]string, 0)
	if len(fields) > 1 {
		newFiles = fields[1:]
	}
	scan := query.Select(fields[0], newFiles...)
	return scan.All(ctx)
}

// count count count dns detector result
func (params *detectorListDNSResultParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().DNSDetectorResult.Query()
	params.where(query)
	return query.Count(ctx)
}

// query do filter query
func (params *detectorFilterDNSParams) query(ctx context.Context) (dnsResults []*ent.DNSDetector, err error) {

	query := dnsdetector.NameContains(params.Keyword)
	id, _ := strconv.Atoi(params.Keyword)
	if id != 0 {
		query = dnsdetector.Or(query, dnsdetector.ID(id))
	}

	return getEntClient().DNSDetector.Query().
		Where(query).
		Limit(10).
		Select("name", "id").
		All(ctx)
}

// add 添加DNS配置
func (*detectorDNSCtrl) add(c *elton.Context) (err error) {
	params := detectorAddDNSParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	result, err := params.save(c.Context(), us.GetInfo().Account)
	if err != nil {
		return
	}
	c.Body = result
	return
}

// list 获取dns配置
func (*detectorDNSCtrl) list(c *elton.Context) (err error) {
	params := detectorListDNSParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	us := getUserSession(c)
	params.Owner = us.GetInfo().Account
	count := -1
	if params.ShouldCount() {
		count, err = params.count(c.Context())
		if err != nil {
			return
		}
	}
	dnses, err := params.queryAll(c.Context())
	if err != nil {
		return
	}
	c.Body = &detectorListDNSResp{
		Count: count,
		DNSES: dnses,
	}
	return
}

// updateByID 更新dns配置
func (*detectorDNSCtrl) updateByID(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	params := detectorUpdateDNSParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	params.CurrentUser = us.GetInfo().Account
	result, err := params.updateByID(c.Context(), id)
	if err != nil {
		return
	}
	c.Body = result
	return
}

// listResult list dns result
func (*detectorDNSCtrl) listResult(c *elton.Context) (err error) {
	params := detectorListDNSResultParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	count := -1
	if params.ShouldCount() {
		count, err = params.count(c.Context())
		if err != nil {
			return
		}
	}
	results, err := params.queryAll(c.Context())
	if err != nil {
		return
	}

	// 根据任务ID获取任务名称
	taskIDList := make([]int, 0)
	ids := map[int]bool{}
	for _, item := range results {
		_, exists := ids[item.Task]
		if exists {
			continue
		}
		taskIDList = append(taskIDList, item.Task)
		ids[item.Task] = true
	}
	// 如果获取失败，忽略（因为仅用于获取任务名称）
	detectors, _ := getEntClient().DNSDetector.Query().
		Where(dnsdetector.IDIn(taskIDList...)).
		Select("name", "id").
		All(c.Context())

	c.CacheMaxAge(time.Minute)
	c.Body = &detectorListDNSResultResp{
		DNSResults:   results,
		DNSDetectors: detectors,
		Count:        count,
	}
	return
}

// getResult get detail of result
func (*detectorDNSCtrl) getResult(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	result, err := getEntClient().DNSDetectorResult.Get(c.Context(), id)
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = result
	return
}

// filter filter detector
func (*detectorDNSCtrl) filter(c *elton.Context) (err error) {
	params := detectorFilterDNSParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	results, err := params.query(c.Context())
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = &detectorFilterDNSResp{
		DNSDetectors: results,
	}
	return
}
