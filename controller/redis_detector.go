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
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/redisdetector"
	"github.com/vicanso/cybertect/ent/redisdetectorresult"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
)

type redisDetectorCtrl struct{}

type (
	redisDetectorAddParams struct {
		detectorAddParams

		Uris []string `json:"uris" validate:"required,dive,uri"`
	}

	redisDetectorListParams struct {
		listParams

		account string
	}

	redisDetectorUpdateParams struct {
		detectorUpdateParams

		account string
		Uris    []string `json:"uris" validate:"omitempty,dive,uri"`
	}

	redisDetectorResultListParams struct {
		detectorListResultParams
	}
)

type (
	redisDetectorListResp struct {
		RedisDetectors ent.RedisDetectors `json:"redisDetectors"`
		Count          int                `json:"count"`
	}
	redisDetectorResultListResp struct {
		RedisDetectorResults ent.RedisDetectorResults `json:"redisDetectorResults"`
		Count                int                      `json:"count"`
	}
)

func init() {
	prefix := "/redis-detectors"
	g := router.NewGroup(
		prefix,
		loadUserSession,
		shouldBeLogin,
	)
	ctrl := redisDetectorCtrl{}

	g.POST(
		"/v1",
		newTrackerMiddleware(cs.ActionDetectorRedisAdd),
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
		newTrackerMiddleware(cs.ActionDetectorRedisUpdate),
		ctrl.updateByID,
	)

	router.NewGroup(prefix).GET(
		"/results/v1",
		ctrl.listResult,
	)
}

func getRedisDetectorClient() *ent.RedisDetectorClient {
	return helper.EntGetClient().RedisDetector
}

func getRedisDetectorResultClient() *ent.RedisDetectorResultClient {
	return helper.EntGetClient().RedisDetectorResult
}

func (addParams *redisDetectorAddParams) save(ctx context.Context) (*ent.RedisDetector, error) {
	return getRedisDetectorClient().Create().
		SetStatus(addParams.Status).
		SetName(addParams.Name).
		SetOwners(addParams.Owners).
		SetReceivers(addParams.Receivers).
		SetTimeout(addParams.Timeout).
		SetDescription(addParams.Description).
		SetUris(addParams.Uris).
		Save(ctx)
}

func (listParams *redisDetectorListParams) where(query *ent.RedisDetectorQuery) {
	account := listParams.account
	if account != "" {
		query.Where(predicate.RedisDetector(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(redisdetector.FieldOwners, account))
		}))
	}
}

func (listParams *redisDetectorListParams) count(ctx context.Context) (int, error) {
	query := getRedisDetectorClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *redisDetectorListParams) queryAll(ctx context.Context) (ent.RedisDetectors, error) {
	query := getRedisDetectorClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (listParams *redisDetectorResultListParams) where(query *ent.RedisDetectorResultQuery) {
	task := listParams.Task
	if task != 0 {
		query.Where(redisdetectorresult.Task(task))
	}
	result := listParams.Result
	if result != 0 {
		query.Where(redisdetectorresult.Result(schema.DetectorResult(result)))
	}
	ms := listParams.GetDurationMillSecond()
	if ms > 0 {
		query.Where(redisdetectorresult.MaxDurationGTE(ms))
	}
	startedAt := listParams.StartedAt
	if !startedAt.IsZero() {
		query.Where(redisdetectorresult.CreatedAtGTE(startedAt))
	}
	endedAt := listParams.EndedAt
	if !endedAt.IsZero() {
		query.Where(redisdetectorresult.CreatedAtLTE(endedAt))
	}
}

func (listParams *redisDetectorResultListParams) count(ctx context.Context) (int, error) {
	query := getRedisDetectorResultClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *redisDetectorResultListParams) queryAll(ctx context.Context) (ent.RedisDetectorResults, error) {
	query := getRedisDetectorResultClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (updateParams *redisDetectorUpdateParams) updateByID(ctx context.Context, id int) (*ent.RedisDetector, error) {
	account := updateParams.account
	if account != "" {
		result, err := getRedisDetectorClient().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if !util.ContainsString(result.Owners, account) {
			return nil, errInvalidUser
		}
	}
	updateOne := getRedisDetectorClient().UpdateOneID(id)
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

	if len(updateParams.Uris) != 0 {
		updateOne.SetUris(updateParams.Uris)
	}

	return updateOne.Save(ctx)
}

func (*redisDetectorCtrl) add(c *elton.Context) error {
	params := redisDetectorAddParams{}
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

func (*redisDetectorCtrl) list(c *elton.Context) error {
	params := redisDetectorListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	us := getUserSession(c)
	if !us.IsAdmin() {
		params.account = us.MustGetInfo().Account
	}
	resp := redisDetectorListResp{
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
	resp.RedisDetectors = result
	c.Body = &resp
	return nil
}

func (*redisDetectorCtrl) updateByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	params := redisDetectorUpdateParams{}
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

func (*redisDetectorCtrl) findByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	result, err := getRedisDetectorClient().Get(c.Context(), id)
	if err != nil {
		return err
	}
	c.Body = result
	return nil
}

func (*redisDetectorCtrl) listResult(c *elton.Context) error {
	params := redisDetectorResultListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	resp := redisDetectorResultListResp{
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

	// 填充检测任务名称
	idList := make([]int, 0)
	for _, item := range result {
		if funk.ContainsInt(idList, item.Task) {
			continue
		}
		idList = append(idList, item.Task)
	}
	detectors, err := getRedisDetectorClient().Query().
		Where(
			redisdetector.IDIn(idList...),
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
	resp.RedisDetectorResults = result
	c.Body = &resp
	return nil
}
