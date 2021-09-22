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
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
)

type httpDetectorCtrl struct{}

type (
	httpDetectorAddParams struct {
		detectorAddParams

		IPS []string `json:"ips" validate:"omitempty,dive,ip"`
		URL string   `json:"url" validate:"required,xHTTP"`
	}
	httpDetectorListParams struct {
		listParams

		account string
	}
	httpDetectorUpdateParams struct {
		detectorUpdateParams

		account string

		IPS []string `json:"ips" validate:"omitempty,dive,ip"`
		URL string   `json:"url" validate:"omitempty,xHTTP"`
	}
)

type (
	// httpDetectorListResp response of list http
	httpDetectorListResp struct {
		HTTPDetectors []*ent.HTTPDetector `json:"httpDetectors"`
		Count         int                 `json:"count"`
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
}

func getHTTPDetectorClient() *ent.HTTPDetectorClient {
	return helper.EntGetClient().HTTPDetector
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
		Save(ctx)
}

func (listParams *httpDetectorListParams) where(query *ent.HTTPDetectorQuery) {
	if listParams.account != "" {
		query.Where(predicate.HTTPDetector(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(httpdetector.FieldOwners, listParams.account))
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

func (updateParams *httpDetectorUpdateParams) updateByID(ctx context.Context, id int) (*ent.HTTPDetector, error) {
	// 指定账号则判断是否是该配置的owner
	if updateParams.account != "" {
		result, err := getHTTPDetectorClient().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if !util.ContainsString(result.Owners, updateParams.account) {
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

	if len(updateParams.IPS) != 0 {
		updateOne.SetIps(updateParams.IPS)
	}
	if updateParams.URL != "" {
		updateOne.SetURL(updateParams.URL)
	}
	// 允许直接修改owner
	if len(updateParams.Owners) != 0 {
		updateOne.SetOwners(updateParams.Owners)
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
