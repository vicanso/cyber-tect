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

	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/http"
	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/ent/user"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
)

type (
	detectorCtrl struct{}

	// detectorAddHTTPParams params of add http
	detectorAddHTTPParams struct {
		Name        string   `json:"name,omitempty" validate:"required,xDetectorName"`
		Status      int      `json:"status,omitempty" validate:"required,xStatus"`
		Description string   `json:"description,omitempty" validate:"required,xDetectorDesc"`
		IPS         []string `json:"ips,omitempty" validate:"omitempty,dive,ip"`
		Receivers   []string `json:"receivers,omitempty" validate:"required,dive,xUserAccount"`
		URL         string   `json:"url,omitempty" validate:"required,xHTTP"`
		Timeout     string   `json:"timeout,omitempty" validate:"required,xDuration"`
	}
	// detectorUpdateHTTPParams params of update http
	detectorUpdateHTTPParams struct {
		CurrentUser string

		Name        string   `json:"name,omitempty" validate:"omitempty,xDetectorName"`
		Status      int      `json:"status,omitempty" validate:"omitempty,xStatus"`
		Description string   `json:"description,omitempty" validate:"omitempty,xDetectorDesc"`
		IPS         []string `json:"ips,omitempty" validate:"omitempty,dive,ip"`
		Receivers   []string `json:"receivers,omitempty" validate:"omitempty,dive,xUserAccount"`
		URL         string   `json:"url,omitempty" validate:"omitempty,xHTTP"`
		Timeout     string   `json:"timeout,omitempty" validate:"omitempty,xDuration"`
	}
	// detectorHTTPListParams params of list http
	detectorHTTPListParams struct {
		listParams

		Owner string
	}
	// detectorHTTPListResp response of list http
	detectorHTTPListResp struct {
		HTTPS []*ent.HTTP `json:"https,omitempty"`
		Count int         `json:"count,omitempty"`
	}
)

const (
	errDetectorCategory = "detector"
)

func init() {
	g := router.NewGroup("/detectors", loadUserSession, shouldBeLogin)

	ctrl := detectorCtrl{}

	// 查询接收者列表
	g.GET(
		"/v1/receivers",
		ctrl.listReceiver,
	)

	// 查询http配置
	g.GET(
		"/v1/https",
		ctrl.listHTTP,
	)
	// 添加http配置
	g.POST(
		"/v1/https",
		newTrackerMiddleware(cs.ActionDetectorHTTPAdd),
		ctrl.addHTTP,
	)
	// 更新http配置
	g.PATCH(
		"/v1/https/{id}",
		newTrackerMiddleware(cs.ActionDetectorHTTPUpdate),
		ctrl.updateHTTPByID,
	)
}

// save http save
func (params *detectorAddHTTPParams) save(ctx context.Context, owner string) (result *ent.HTTP, err error) {
	return getEntClient().HTTP.Create().
		SetName(params.Name).
		SetStatus(schema.Status(params.Status)).
		SetDescription(params.Description).
		SetIps(params.IPS).
		SetURL(params.URL).
		SetReceivers(params.Receivers).
		SetTimeout(params.Timeout).
		SetOwner(owner).
		Save(ctx)
}

// where http where
func (params *detectorHTTPListParams) where(query *ent.HTTPQuery) *ent.HTTPQuery {
	if params.Owner != "" {
		query = query.Where(http.OwnerEQ(params.Owner))
	}
	return query
}

// queryAll query all http detector
func (params *detectorHTTPListParams) queryAll(ctx context.Context) (https []*ent.HTTP, err error) {
	query := getEntClient().HTTP.Query()
	query = query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	query = params.where(query)

	return query.All(ctx)
}

// count count http detector
func (params *detectorHTTPListParams) count(ctx context.Context) (count int, err error) {
	query := getEntClient().HTTP.Query()
	query = params.where(query)
	return query.Count(ctx)
}

// updateByID update http detector by id
func (params *detectorUpdateHTTPParams) updateByID(ctx context.Context, id int) (result *ent.HTTP, err error) {
	currentHTTP, err := getEntClient().HTTP.Get(ctx, id)
	if err != nil {
		return
	}
	if currentHTTP.Owner != params.CurrentUser {
		err = hes.New("仅能创建者允许修改配置", errDetectorCategory)
		return
	}

	updateOne := getEntClient().HTTP.UpdateOneID(id)

	if params.Name != "" {
		updateOne = updateOne.SetName(params.Name)
	}
	if params.Status != 0 {
		updateOne = updateOne.SetStatus(schema.Status(params.Status))
	}
	if params.Description != "" {
		updateOne = updateOne.SetDescription(params.Description)
	}
	if len(params.IPS) != 0 {
		updateOne = updateOne.SetIps(params.IPS)
	}
	if len(params.Receivers) != 0 {
		updateOne = updateOne.SetReceivers(params.Receivers)
	}
	if params.URL != "" {
		updateOne = updateOne.SetURL(params.URL)
	}
	if params.Timeout != "" {
		updateOne = updateOne.SetTimeout(params.Timeout)
	}
	return updateOne.Save(ctx)
}

// addHTTP 添加http配置
func (*detectorCtrl) addHTTP(c *elton.Context) (err error) {
	params := detectorAddHTTPParams{}
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

// listHTTP 获取http配置
func (*detectorCtrl) listHTTP(c *elton.Context) (err error) {
	params := detectorHTTPListParams{}
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
	https, err := params.queryAll(c.Context())
	if err != nil {
		return
	}
	c.Body = &detectorHTTPListResp{
		Count: count,
		HTTPS: https,
	}
	return
}

// updateHTTPByID 更新http配置
func (*detectorCtrl) updateHTTPByID(c *elton.Context) (err error) {
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
	params.CurrentUser = us.MustGetInfo().Account
	result, err := params.updateByID(c.Context(), id)
	if err != nil {
		return
	}
	c.Body = result
	return
}

// listReceiver 获取接收者列表
func (*detectorCtrl) listReceiver(c *elton.Context) (err error) {

	users, err := getEntClient().User.Query().
		Where(user.StatusEQ(schema.StatusEnabled)).
		Select("account").
		All(c.Context())
	if err != nil {
		return
	}
	receivers := make([]string, len(users))
	for index, user := range users {
		receivers[index] = user.Account
	}
	c.Body = map[string][]string{
		"receivers": receivers,
	}
	return
}
