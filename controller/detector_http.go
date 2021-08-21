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
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
)

type (
	detectorHTTPCtrl struct{}
	// detectorAddHTTPParams params of add http
	detectorAddHTTPParams struct {
		detectorAddParams

		IPS []string `json:"ips,omitempty" validate:"omitempty,dive,ip"`
		URL string   `json:"url,omitempty" validate:"required,xHTTP"`
	}
	// detectorUpdateHTTPParams params of update http
	detectorUpdateHTTPParams struct {
		detectorUpdateParams

		IPS []string `json:"ips,omitempty" validate:"omitempty,dive,ip"`
		URL string   `json:"url,omitempty" validate:"omitempty,xHTTP"`
	}
	// detectorListHTTPParams params of list http
	detectorListHTTPParams struct {
		listParams

		Owner string
	}
	// detectorListHTTPResp response of list http
	detectorListHTTPResp struct {
		HTTPS []*ent.HTTPDetector `json:"https,omitempty"`
		Count int                 `json:"count,omitempty"`
	}

	// detectorListHTTPResultParams params of list http result
	detectorListHTTPResultParams struct {
		detectorListResultParams
	}

	// detectorListHTTPResultResp response of list http result
	detectorListHTTPResultResp struct {
		HTTPResults   []*ent.HTTPDetectorResult `json:"httpResults,omitempty"`
		HTTPDetectors []*ent.HTTPDetector       `json:"httpDetectors,omitempty"`
		Count         int                       `json:"count,omitempty"`
	}

	// detectorFilterHTTPParams params of filter http
	detectorFilterHTTPParams struct {
		Keyword string `json:"keyword,omitempty" validate:"required,xKeyword"`
	}
	// detectorFilterHTTPResp response of filter http
	detectorFilterHTTPResp struct {
		HTTPDetectors []*ent.HTTPDetector `json:"httpDetectors,omitempty"`
	}
)

func init() {
	prefix := "/detectors/v1/https"
	g := router.NewGroup(prefix, loadUserSession, shouldBeLogin)
	nsg := router.NewGroup(prefix)

	ctrl := detectorHTTPCtrl{}

	// 查询http配置
	g.GET(
		"",
		ctrl.list,
	)
	// 添加http配置
	g.POST(
		"",
		newTrackerMiddleware(cs.ActionDetectorHTTPAdd),
		ctrl.add,
	)
	// 更新http配置
	g.PATCH(
		"/{id}",
		newTrackerMiddleware(cs.ActionDetectorHTTPUpdate),
		ctrl.updateByID,
	)

	// HTTP筛选
	nsg.GET(
		"/filter",
		ctrl.filter,
	)

	// 查询http检测结果
	nsg.GET(
		"/results",
		ctrl.listResult,
	)
	// 查询http检测结果详情
	nsg.GET(
		"/results/{id}",
		ctrl.getResult,
	)
}

// save http save
func (params *detectorAddHTTPParams) save(ctx context.Context, owner string) (result *ent.HTTPDetector, err error) {
	return getEntClient().HTTPDetector.Create().
		SetName(params.Name).
		SetStatus(schema.Status(params.Status)).
		SetDescription(params.Description).
		SetReceivers(params.Receivers).
		SetTimeout(params.Timeout).
		SetIps(params.IPS).
		SetURL(params.URL).
		SetOwner(owner).
		Save(ctx)
}

// where http where
func (params *detectorListHTTPParams) where(query *ent.HTTPDetectorQuery) {
	if params.Owner != "" {
		query.Where(httpdetector.Owner(params.Owner))
	}
}

// queryAll query all http detector
func (params *detectorListHTTPParams) queryAll(ctx context.Context) (https []*ent.HTTPDetector, err error) {
	query := getEntClient().HTTPDetector.Query()
	query = query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	params.where(query)

	return query.All(ctx)
}

// count count http detector
func (params *detectorListHTTPParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().HTTPDetector.Query()
	params.where(query)
	return query.Count(ctx)
}

// updateByID update http detector by id
func (params *detectorUpdateHTTPParams) updateByID(ctx context.Context, id int) (result *ent.HTTPDetector, err error) {
	currentHTTP, err := getEntClient().HTTPDetector.Get(ctx, id)
	if err != nil {
		return
	}
	if currentHTTP.Owner != params.CurrentUser {
		err = errInvalidUser.Clone()
		return
	}

	updateOne := getEntClient().HTTPDetector.UpdateOneID(id)

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

	if params.URL != "" {
		updateOne.SetURL(params.URL)
	}

	return updateOne.Save(ctx)
}

// where http detector result where
func (params *detectorListHTTPResultParams) where(query *ent.HTTPDetectorResultQuery) {
	task := params.GetTaskID()
	if task != 0 {
		query.Where(httpdetectorresult.Task(task))
	}
	result := params.GetResult()
	if result != 0 {
		query.Where(httpdetectorresult.Result(result))
	}
	ms := params.GetDurationMillSecond()
	if ms > 0 {
		query.Where(httpdetectorresult.MaxDurationGTE(ms))
	}
	createdAtRange := params.GetCreatedAtGT()
	if !createdAtRange.IsZero() {
		query.Where(httpdetectorresult.CreatedAtGTE(createdAtRange))
	}

}

// queryAll query all http result
func (params *detectorListHTTPResultParams) queryAll(ctx context.Context) (httpResults []*ent.HTTPDetectorResult, err error) {
	query := getEntClient().HTTPDetectorResult.Query()

	query = query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	params.where(query)
	fields := params.GetFields()
	if len(fields) == 0 {
		return query.All(ctx)
	}
	scan := query.Select(fields...)
	return scan.All(ctx)
}

// count count http detector result
func (params *detectorListHTTPResultParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().HTTPDetectorResult.Query()
	params.where(query)
	return query.Count(ctx)
}

// query do filter query
func (params *detectorFilterHTTPParams) query(ctx context.Context) (httpResults []*ent.HTTPDetector, err error) {
	query := httpdetector.NameContains(params.Keyword)
	id, _ := strconv.Atoi(params.Keyword)
	if id != 0 {
		query = httpdetector.Or(query, httpdetector.ID(id))
	}

	return getEntClient().HTTPDetector.Query().
		Where(query).
		Limit(10).
		Select("name", "id").
		All(ctx)
}

// add 添加http配置
func (*detectorHTTPCtrl) add(c *elton.Context) (err error) {
	params := detectorAddHTTPParams{}
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

// list 获取http配置
func (*detectorHTTPCtrl) list(c *elton.Context) (err error) {
	params := detectorListHTTPParams{}
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
	https, err := params.queryAll(c.Context())
	if err != nil {
		return
	}
	c.Body = &detectorListHTTPResp{
		Count: count,
		HTTPS: https,
	}
	return
}

// updateByID 更新http配置
func (*detectorHTTPCtrl) updateByID(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	params := detectorUpdateHTTPParams{}
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

// listResult list http result
func (*detectorHTTPCtrl) listResult(c *elton.Context) (err error) {
	params := detectorListHTTPResultParams{}
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
	detectors, _ := getEntClient().HTTPDetector.Query().
		Where(httpdetector.IDIn(taskIDList...)).
		Select("name", "id").
		All(c.Context())

	c.CacheMaxAge(time.Minute)
	c.Body = &detectorListHTTPResultResp{
		HTTPResults:   results,
		HTTPDetectors: detectors,
		Count:         count,
	}
	return
}

// getResult get
func (*detectorHTTPCtrl) getResult(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	result, err := getEntClient().HTTPDetectorResult.Get(c.Context(), id)
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = result
	return
}

// filter filter detector
func (*detectorHTTPCtrl) filter(c *elton.Context) (err error) {
	params := detectorFilterHTTPParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	results, err := params.query(c.Context())
	if err != nil {
		return
	}
	c.CacheMaxAge(time.Minute)
	c.Body = &detectorFilterHTTPResp{
		HTTPDetectors: results,
	}
	return
}
