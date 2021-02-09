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
	"github.com/vicanso/cybertect/ent/pingdetector"
	"github.com/vicanso/cybertect/ent/pingdetectorresult"
	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
)

type (
	detectorPingCtrl struct{}
	// detectorAddPingParams params of add ping
	detectorAddPingParams struct {
		detectorAddParams

		IPS []string `json:"ips,omitempty" validate:"required,dive,ip"`
	}
	// detectorUpdatePingParams params of update ping
	detectorUpdatePingParams struct {
		detectorUpdateParams

		IPS []string `json:"ips,omitempty" validate:"omitempty,dive,ip"`
	}
	// detectorListPingParams params of list ping
	detectorListPingParams struct {
		listParams

		Owner string
	}
	// detectorListPingResp response of list ping
	detectorListPingResp struct {
		Pings []*ent.PingDetector `json:"pings,omitempty"`
		Count int                 `json:"count,omitempty"`
	}

	// detectorListPingResultParams params of list ping result
	detectorListPingResultParams struct {
		detectorListResultParams
	}
	// detectorListPingResultResp response of list ping result
	detectorListPingResultResp struct {
		PingResults   []*ent.PingDetectorResult `json:"pingResults,omitempty"`
		PingDetectors []*ent.PingDetector       `json:"pingDetectors,omitempty"`
		Count         int                       `json:"count,omitempty"`
	}

	// detectorFilterPingParams params of filter ping
	detectorFilterPingParams struct {
		Keyword string `json:"keyword,omitempty" validate:"required,xKeyword"`
	}
	// detectorFilterPingResp response of filter ping
	detectorFilterPingResp struct {
		PingDetectors []*ent.PingDetector `json:"pingDetectors,omitempty"`
	}
)

func init() {
	prefix := "/detectors/v1/pings"
	g := router.NewGroup(prefix, loadUserSession, shouldBeLogin)
	nsg := router.NewGroup(prefix)

	ctrl := detectorPingCtrl{}
	// 查询ping配置
	g.GET(
		"",
		ctrl.list,
	)
	// 添加ping配置
	g.POST(
		"",
		newTrackerMiddleware(cs.ActionDetectorPingAdd),
		ctrl.add,
	)
	// 更新ping配置
	g.PATCH(
		"/{id}",
		newTrackerMiddleware(cs.ActionDetectorPingUpdate),
		ctrl.updateByID,
	)

	// 筛选ping配置
	nsg.GET(
		"/filter",
		ctrl.filter,
	)

	// 查询ping检测结果
	nsg.GET(
		"/results",
		ctrl.listResult,
	)
	// 查询ping检测结果详情
	nsg.GET(
		"/results/{id}",
		ctrl.getResult,
	)
}

// save ping save
func (params *detectorAddPingParams) save(ctx context.Context, owner string) (result *ent.PingDetector, err error) {
	return getEntClient().PingDetector.Create().
		SetName(params.Name).
		SetStatus(schema.Status(params.Status)).
		SetDescription(params.Description).
		SetReceivers(params.Receivers).
		SetTimeout(params.Timeout).
		SetOwner(owner).
		SetIps(params.IPS).
		Save(ctx)
}

// where ping where
func (params *detectorListPingParams) where(query *ent.PingDetectorQuery) {
	if params.Owner != "" {
		query.Where(pingdetector.OwnerEQ(params.Owner))
	}
}

// queryAll query all ping detector
func (params *detectorListPingParams) queryAll(ctx context.Context) (pings []*ent.PingDetector, err error) {
	query := getEntClient().PingDetector.Query()

	query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	params.where(query)

	return query.All(ctx)
}

// count count ping detector
func (params *detectorListPingParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().PingDetector.Query()
	params.where(query)
	return query.Count(ctx)
}

// updateByID update ping detector by id
func (params *detectorUpdatePingParams) updateByID(ctx context.Context, id int) (result *ent.PingDetector, err error) {
	currentPing, err := getEntClient().PingDetector.Get(ctx, id)
	if err != nil {
		return
	}
	if currentPing.Owner != params.CurrentUser {
		err = errInvalidUser.Clone()
		return
	}

	updateOne := getEntClient().PingDetector.UpdateOneID(id)

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

	if len(params.IPS) != 0 {
		updateOne.SetIps(params.IPS)
	}

	return updateOne.Save(ctx)
}

// where ping detector result where
func (params *detectorListPingResultParams) where(query *ent.PingDetectorResultQuery) {
	task := params.GetTaskID()
	if task != 0 {
		query.Where(pingdetectorresult.Task(task))
	}

	result := params.GetResult()
	if result != 0 {
		query.Where(pingdetectorresult.Result(result))
	}

	ms := params.GetDurationMillSecond()
	if ms > 0 {
		query.Where(pingdetectorresult.MaxDurationGTE(ms))
	}
}

// queryAll query all ping result
func (params *detectorListPingResultParams) queryAll(ctx context.Context) (pingResults []*ent.PingDetectorResult, err error) {
	query := getEntClient().PingDetectorResult.Query()

	query = query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	params.where(query)
	fields := params.GetFields()
	if len(fields) == 0 {
		return query.All(ctx)
	}
	scan := query.Select(fields[0], fields[1:]...)
	return scan.All(ctx)
}

// count count ping detector result
func (params *detectorListPingResultParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().PingDetectorResult.Query()
	params.where(query)
	return query.Count(ctx)
}

// query do filter query
func (params *detectorFilterPingParams) query(ctx context.Context) (pingResults []*ent.PingDetector, err error) {
	query := pingdetector.NameContains(params.Keyword)
	id, _ := strconv.Atoi(params.Keyword)
	if id != 0 {
		query = pingdetector.Or(query, pingdetector.ID(id))
	}
	return getEntClient().PingDetector.Query().
		Where(query).
		Limit(10).
		Select("name", "id").
		All(ctx)
}

// add 添加Ping记录
func (*detectorPingCtrl) add(c *elton.Context) (err error) {
	params := detectorAddPingParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	result, err := params.save(c.Context(), us.MustGetInfo().Account)
	if err != nil {
		return
	}
	c.Body = result
	return
}

// list 获取Ping配置
func (*detectorPingCtrl) list(c *elton.Context) (err error) {
	params := detectorListPingParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	us := getUserSession(c)
	params.Owner = us.MustGetInfo().Account
	count := -1
	if params.ShouldCount() {
		count, err = params.count(c.Context())
		if err != nil {
			return
		}
	}
	pings, err := params.queryAll(c.Context())
	if err != nil {
		return
	}
	c.Body = &detectorListPingResp{
		Count: count,
		Pings: pings,
	}
	return
}

// updateByID 更新Ping配置
func (*detectorPingCtrl) updateByID(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	params := detectorUpdatePingParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	params.CurrentUser = us.MustGetInfo().Account
	result, err := params.updateByID(c.Context(), id)
	if err != nil {
		return
	}
	c.Body = result
	return
}

// listResult list ping result
func (*detectorPingCtrl) listResult(c *elton.Context) (err error) {
	params := detectorListPingResultParams{}
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
	// 如果获取失败，忽略（因为仅用于获取任务名称
	detectors, _ := getEntClient().PingDetector.Query().
		Where(pingdetector.IDIn(taskIDList...)).
		Select("name", "id").
		All(c.Context())

	c.CacheMaxAge(time.Minute)
	c.Body = &detectorListPingResultResp{
		PingResults:   results,
		PingDetectors: detectors,
		Count:         count,
	}
	return
}

// getResult get ping result
func (*detectorPingCtrl) getResult(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	result, err := getEntClient().PingDetectorResult.Get(c.Context(), id)
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = result
	return
}

// filter filter detector
func (*detectorPingCtrl) filter(c *elton.Context) (err error) {
	params := detectorFilterPingParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	results, err := params.query(c.Context())
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = &detectorFilterPingResp{
		PingDetectors: results,
	}
	return
}
