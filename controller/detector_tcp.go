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
	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/ent/tcpdetector"
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
)

func init() {
	g := router.NewGroup("/detectors/v1/tcps", loadUserSession, shouldBeLogin)

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

// add 添加TCP记录
func (*detectorTCPCtrl) add(c *elton.Context) (err error) {
	params := detectorAddTCPParams{}
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

// list 获取TCP配置
func (*detectorTCPCtrl) list(c *elton.Context) (err error) {
	params := detectorListTCPParams{}
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
	params.CurrentUser = us.MustGetInfo().Account
	result, err := params.updateByID(c.Context(), id)
	if err != nil {
		return
	}
	c.Body = result
	return
}
