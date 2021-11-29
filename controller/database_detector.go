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
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/thoas/go-funk"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/databasedetector"
	"github.com/vicanso/cybertect/ent/databasedetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/elton"
)

type databaseDetectorCtrl struct{}

type (
	databaseDetectorAddParams struct {
		detectorAddParams

		Uris []string `json:"uris" validate:"required,dive,xDetectorDatabaseURI"`
		// tls cert pem block
		CertPem string `json:"certPem" validate:"omitempty,xDetectorTLSPemData"`
		// tls key pem block
		KeyPem string `json:"keyPem" validate:"omitempty,xDetectorTLSPemData"`
	}

	databaseDetectorListParams struct {
		listParams

		account string
	}

	databaseDetectorUpdateParams struct {
		detectorUpdateParams

		account string
		Uris    []string `json:"uris" validate:"omitempty,dive,xDetectorDatabaseURI"`

		// tls cert pem block
		CertPem string `json:"certPem" validate:"omitempty,xDetectorTLSPemData"`
		// tls key pem block
		KeyPem string `json:"keyPem" validate:"omitempty,xDetectorTLSPemData"`
	}

	databaseDetectorResultListParams struct {
		detectorListResultParams
	}
)

type (
	redisDetectorListResp struct {
		DatabaseDetectors ent.DatabaseDetectors `json:"databaseDetectors"`
		Count             int                   `json:"count"`
	}
	databasedetectorresultListResp struct {
		DatabaseDetectorResults ent.DatabaseDetectorResults `json:"databaseDetectorResults"`
		Count                   int                         `json:"count"`
	}
)

func init() {
	prefix := "/database-detectors"
	g := router.NewGroup(
		prefix,
		loadUserSession,
		shouldBeLogin,
	)
	ctrl := databaseDetectorCtrl{}

	g.POST(
		"/v1",
		newTrackerMiddleware(cs.ActionDetectorDatabaseAdd),
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
		newTrackerMiddleware(cs.ActionDetectorDatabaseUpdate),
		ctrl.updateByID,
	)

	g.GET(
		"/results/v1",
		ctrl.listResult,
	)
}

func getDatabaseDetectorClient() *ent.DatabaseDetectorClient {
	return helper.EntGetClient().DatabaseDetector
}

func getDatabaseDetectorResultClient() *ent.DatabaseDetectorResultClient {
	return helper.EntGetClient().DatabaseDetectorResult
}

func (addParams *databaseDetectorAddParams) save(ctx context.Context) (*ent.DatabaseDetector, error) {
	return getDatabaseDetectorClient().Create().
		SetStatus(addParams.Status).
		SetName(addParams.Name).
		SetOwners(addParams.Owners).
		SetReceivers(addParams.Receivers).
		SetTimeout(addParams.Timeout).
		SetDescription(addParams.Description).
		SetUris(addParams.Uris).
		SetCertPem(addParams.CertPem).
		SetKeyPem(addParams.KeyPem).
		Save(ctx)
}

func (listParams *databaseDetectorListParams) where(query *ent.DatabaseDetectorQuery) {
	account := listParams.account
	if account != "" {
		query.Where(predicate.DatabaseDetector(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(databasedetector.FieldOwners, account))
		}))
	}
}

func (listParams *databaseDetectorListParams) count(ctx context.Context) (int, error) {
	query := getDatabaseDetectorClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *databaseDetectorListParams) queryAll(ctx context.Context) (ent.DatabaseDetectors, error) {
	query := getDatabaseDetectorClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	return query.All(ctx)
}

func (listParams *databaseDetectorResultListParams) where(query *ent.DatabaseDetectorResultQuery) {
	tasks := listParams.doTaskFilter()
	if len(tasks) != 0 {
		query.Where(databasedetectorresult.TaskIn(tasks...))
	}
	result := listParams.Result
	if result != 0 {
		query.Where(databasedetectorresult.Result(schema.DetectorResult(result)))
	}
	ms := listParams.GetDurationMillSecond()
	if ms > 0 {
		query.Where(databasedetectorresult.MaxDurationGTE(ms))
	}
	startedAt := listParams.StartedAt
	if !startedAt.IsZero() {
		query.Where(databasedetectorresult.CreatedAtGTE(startedAt))
	}
	endedAt := listParams.EndedAt
	if !endedAt.IsZero() {
		query.Where(databasedetectorresult.CreatedAtLTE(endedAt))
	}
}

func (listParams *databaseDetectorResultListParams) count(ctx context.Context) (int, error) {
	query := getDatabaseDetectorResultClient().Query()
	listParams.where(query)
	return query.Count(ctx)
}

func (listParams *databaseDetectorResultListParams) queryAll(ctx context.Context) (ent.DatabaseDetectorResults, error) {
	query := getDatabaseDetectorResultClient().Query()
	query = query.Limit(listParams.GetLimit()).
		Offset(listParams.GetOffset()).
		Order(listParams.GetOrders()...)
	listParams.where(query)
	fields := listParams.GetFields()
	if len(fields) != 0 {
		results := make(ent.DatabaseDetectorResults, 0)
		err := query.Select(fields...).Scan(ctx, &results)
		return results, err
	}
	return query.All(ctx)
}

func (updateParams *databaseDetectorUpdateParams) updateByID(ctx context.Context, id int) (*ent.DatabaseDetector, error) {
	account := updateParams.account
	if account != "" {
		result, err := getDatabaseDetectorClient().Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if !util.ContainsString(result.Owners, account) {
			return nil, errInvalidUser
		}
	}
	updateOne := getDatabaseDetectorClient().UpdateOneID(id)
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
	if updateParams.CertPem != "" {
		// trim space，如果是空格，则会清除该字段
		updateOne.SetCertPem(strings.TrimSpace(updateParams.CertPem))
	}
	if updateParams.KeyPem != "" {
		// trim space，如果是空格，则会清除该字段
		updateOne.SetKeyPem(strings.TrimSpace(updateParams.KeyPem))
	}

	return updateOne.Save(ctx)
}

func (*databaseDetectorCtrl) add(c *elton.Context) error {
	params := databaseDetectorAddParams{}
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

func (*databaseDetectorCtrl) list(c *elton.Context) error {
	params := databaseDetectorListParams{}
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
	resp.DatabaseDetectors = result
	c.Body = &resp
	return nil
}

func (*databaseDetectorCtrl) updateByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	params := databaseDetectorUpdateParams{}
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

func (*databaseDetectorCtrl) findByID(c *elton.Context) error {
	id, err := getIDFromParams(c)
	if err != nil {
		return err
	}
	result, err := getDatabaseDetectorClient().Get(c.Context(), id)
	if err != nil {
		return err
	}
	c.Body = result
	return nil
}

func (*databaseDetectorCtrl) listResult(c *elton.Context) error {
	params := databaseDetectorResultListParams{}
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	resp := databasedetectorresultListResp{
		Count: -1,
	}
	params.Tasks, err = getDetectorTasksByReceiver(
		c.Context(),
		detectorCategoryDatabase,
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
	detectors, err := getDatabaseDetectorClient().Query().
		Where(
			databasedetector.IDIn(idList...),
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
	resp.DatabaseDetectorResults = result
	c.Body = &resp
	return nil
}
