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
	"github.com/vicanso/cybertect/ent/pingdetector"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
)

type pingDetectorCtrl struct{}

type (
	pingDetectorAddParams struct {
		detectorAddParams

		IPS []string `json:"ips" validate:"required,dive,ip"`
	}
	pingDetectorListParams struct {
		listParams

		account string
	}
	pingDetectorUpdateParams struct {
		detectorUpdateParams

		account string
		IPS     []string `json:"ips" validate:"omitempty,dive,ip"`
	}
)

type (
	pingDetectorListResp struct {
		PingDetectors []*ent.PingDetector `json:"pingDetectors"`
		Count         int                 `json:"count"`
	}
)

func init() {
	g := router.NewGroup(
		"/ping-detectors",
		loadUserSession,
		shouldBeLogin,
	)
	ctrl := pingDetectorCtrl{}

	g.POST(
		"/v1",
		newTrackerMiddleware(cs.ActionDetectorPingAdd),
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
		newTrackerMiddleware(cs.ActionDetectorPingUpdate),
		ctrl.updateByID,
	)
}

func getPingDetectorClient() *ent.PingDetectorClient {
	return helper.EntGetClient().PingDetector
}

func (addParams *pingDetectorAddParams) save(ctx context.Context) (*ent.PingDetector, error) {
	return getPingDetectorClient().Create().
		SetStatus(addParams.Status).
		SetName(addParams.Name).
		SetOwners(addParams.Owners).
		SetReceivers(addParams.Receivers).
		SetTimeout(addParams.Timeout).
		SetDescription(addParams.Description).
		SetIps(addParams.IPS).
		Save(ctx)
}

func (listParams *pingDetectorListParams) where(query *ent.PingDetectorQuery) {
	account := listParams.account
	if account != "" {
		query.Where(predicate.PingDetector(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(pingdetector.FieldOwners, account))
		}))
	}
}

func (listParams *pingDetectorListParams) count(ctx context.Context) (int, error) {
	query := getPingDetectorClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *pingDetectorListParams) queryAll(ctx context.Context) ([]*ent.PingDetector, error) {
	query := getPingDetectorClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (updateParams *pingDetectorUpdateParams) updateByID(ctx context.Context, id int) (*ent.PingDetector, error) {
	account := updateParams.account
	if account != "" {
		result, err := getPingDetectorClient().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if !util.ContainsString(result.Owners, account) {
			return nil, errInvalidUser
		}
	}
	updateOne := getPingDetectorClient().UpdateOneID(id)
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

	return updateOne.Save(ctx)
}

func (*pingDetectorCtrl) add(c *elton.Context) error {
	params := pingDetectorAddParams{}
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

func (*pingDetectorCtrl) list(c *elton.Context) error {
	params := pingDetectorListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	us := getUserSession(c)
	if !us.IsAdmin() {
		params.account = us.MustGetInfo().Account
	}
	resp := pingDetectorListResp{
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
	resp.PingDetectors = result
	c.Body = &resp
	return nil
}

func (*pingDetectorCtrl) updateByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	params := pingDetectorUpdateParams{}
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

func (*pingDetectorCtrl) findByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	result, err := getPingDetectorClient().Get(c.Context(), id)
	if err != nil {
		return err
	}
	c.Body = result
	return nil
}
