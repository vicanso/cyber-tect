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
	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/ent/tcpdetector"
	"github.com/vicanso/cybertect/ent/tcpdetectorresult"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
)

type (
	detectorTCPCtrl struct{}
	// detectorAddTCPParams params of add tcp
	detectorAddTCPParams struct {
		detectorAddParams

		Addrs []string `json:"addrs,omitempty" validate:"required,dive,hostname_port"`
	}
	// detectorUpdateTCPParams params of update tcp
	detectorUpdateTCPParams struct {
		detectorUpdateParams

		Addrs []string `json:"addrs,omitempty" validate:"omitempty,dive,hostname_port"`
	}
	// detectorListTCPParams params of list tcp
	detectorListTCPParams struct {
		listParams

		Owner string
	}
	// detectorListTCPResp response of list tcp
	detectorListTCPResp struct {
		Tcps  []*ent.TCPDetector `json:"tcps,omitempty"`
		Count int                `json:"count,omitempty"`
	}

	// detectorListTCPResultParams params of list tcp result
	detectorListTCPResultParams struct {
		detectorListResultParams
	}
	// detectorListTCPResultResp response of list tcp result
	detectorListTCPResultResp struct {
		TCPResults   []*ent.TCPDetectorResult `json:"tcpResults,omitempty"`
		TCPDetectors []*ent.TCPDetector       `json:"tcpDetectors,omitempty"`
		Count        int                      `json:"count,omitempty"`
	}

	// detectorFilterTCPParams params of filter tcp
	detectorFilterTCPParams struct {
		Keyword string `json:"keyword,omitempty" validate:"required,xKeyword"`
	}
	// detectorFilterTCPResp response of filter tcp
	detectorFilterTCPResp struct {
		TCPDetectors []*ent.TCPDetector `json:"tcpDetectors,omitempty"`
	}
)

func init() {
	prefix := "/detectors/v1/tcps"
	g := router.NewGroup(prefix, loadUserSession, shouldBeLogin)
	nsg := router.NewGroup(prefix)

	ctrl := detectorTCPCtrl{}
	// 查询tcp配置
	g.GET(
		"",
		ctrl.list,
	)
	// 添加tcp配置
	g.POST(
		"",
		newTrackerMiddleware(cs.ActionDetectorTCPAdd),
		ctrl.add,
	)
	// 更新tcp配置
	g.PATCH(
		"/{id}",
		newTrackerMiddleware(cs.ActionDetectorTCPUpdate),
		ctrl.updateByID,
	)

	// TCP 筛选
	nsg.GET(
		"/filter",
		ctrl.filter,
	)

	// 查询tcp检测结果
	nsg.GET(
		"/results",
		ctrl.listResult,
	)
	// 查询tcp检测结果详情
	nsg.GET(
		"/results/{id}",
		ctrl.getResult,
	)
}

// save tcp save
func (params *detectorAddTCPParams) save(ctx context.Context, owner string) (result *ent.TCPDetector, err error) {
	return getEntClient().TCPDetector.Create().
		SetName(params.Name).
		SetStatus(schema.Status(params.Status)).
		SetDescription(params.Description).
		SetReceivers(params.Receivers).
		SetTimeout(params.Timeout).
		SetAddrs(params.Addrs).
		SetOwner(owner).
		Save(ctx)
}

// where tcp where
func (params *detectorListTCPParams) where(query *ent.TCPDetectorQuery) {
	if params.Owner != "" {
		query.Where(tcpdetector.OwnerEQ(params.Owner))
	}
}

// queryAll query all tcp detector
func (params *detectorListTCPParams) queryAll(ctx context.Context) (tcps []*ent.TCPDetector, err error) {
	query := getEntClient().TCPDetector.Query()

	query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	params.where(query)
	return query.All(ctx)
}

// count count tcp detector
func (params *detectorListTCPParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().TCPDetector.Query()
	params.where(query)
	return query.Count(ctx)
}

// updateByID update tcp detector by id
func (params *detectorUpdateTCPParams) updateByID(ctx context.Context, id int) (result *ent.TCPDetector, err error) {
	currentTCP, err := getEntClient().TCPDetector.Get(ctx, id)
	if err != nil {
		return
	}
	if currentTCP.Owner != params.CurrentUser {
		err = errInvalidUser.Clone()
		return
	}

	updateOne := getEntClient().TCPDetector.UpdateOneID(id)

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

	if len(params.Addrs) != 0 {
		updateOne.SetAddrs(params.Addrs)
	}

	return updateOne.Save(ctx)
}

// where tcp detector result where
func (params *detectorListTCPResultParams) where(query *ent.TCPDetectorResultQuery) {
	task := params.GetTaskID()
	if task != 0 {
		query.Where(tcpdetectorresult.Task(task))
	}

	result := params.GetResult()
	if result != 0 {
		query.Where(tcpdetectorresult.Result(result))
	}

	ms := params.GetDurationMillSecond()
	if ms > 0 {
		query.Where(tcpdetectorresult.MaxDurationGTE(ms))
	}
}

// query query all tcp result
func (params *detectorListTCPResultParams) queryAll(ctx context.Context) (tcpResults []*ent.TCPDetectorResult, err error) {
	query := getEntClient().TCPDetectorResult.Query()
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

// count count tcp detector result
func (params *detectorListTCPResultParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().TCPDetectorResult.Query()
	params.where(query)
	return query.Count(ctx)
}

// query do filter query
func (params *detectorFilterTCPParams) query(ctx context.Context) (tcpResults []*ent.TCPDetector, err error) {
	query := tcpdetector.NameContains(params.Keyword)
	id, _ := strconv.Atoi(params.Keyword)
	if id != 0 {
		query = tcpdetector.Or(query, tcpdetector.ID(id))
	}
	return getEntClient().TCPDetector.Query().
		Where(query).
		Limit(10).
		Select("name", "id").
		All(ctx)
}

// add 添加TCP记录
func (*detectorTCPCtrl) add(c *elton.Context) (err error) {
	params := detectorAddTCPParams{}
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

// list 获取TCP配置
func (*detectorTCPCtrl) list(c *elton.Context) (err error) {
	params := detectorListTCPParams{}
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
	tcps, err := params.queryAll(c.Context())
	if err != nil {
		return
	}
	c.Body = &detectorListTCPResp{
		Count: count,
		Tcps:  tcps,
	}
	return
}

// updateByID 更新TCP配置
func (*detectorTCPCtrl) updateByID(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	params := detectorUpdateTCPParams{}
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

// listResult list tcp result
func (*detectorTCPCtrl) listResult(c *elton.Context) (err error) {
	params := detectorListTCPResultParams{}
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
	detectors, _ := getEntClient().TCPDetector.Query().
		Where(tcpdetector.IDIn(taskIDList...)).
		Select("name", "id").
		All(c.Context())

	c.CacheMaxAge(time.Minute)
	c.Body = &detectorListTCPResultResp{
		TCPResults:   results,
		TCPDetectors: detectors,
		Count:        count,
	}
	return
}

// getResult get tcp result
func (*detectorTCPCtrl) getResult(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	result, err := getEntClient().TCPDetectorResult.Get(c.Context(), id)
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = result
	return
}

// filter filter detector
func (*detectorTCPCtrl) filter(c *elton.Context) (err error) {
	params := detectorFilterTCPParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	results, err := params.query(c.Context())
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = &detectorFilterTCPResp{
		TCPDetectors: results,
	}
	return
}
