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

// 应用相关配置，包括IP拦截、路由mock、路由并发限制等配置信息

package controller

import (
	"context"
	"time"

	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/configuration"
	confSchema "github.com/vicanso/cybertect/ent/configuration"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/service"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
)

type (
	configurationCtrl struct{}

	// configurationListResp 配置列表响应
	configurationListResp struct {
		Configurations []*ent.Configuration `json:"configurations,omitempty"`
		Count          int                  `json:"count,omitempty"`
	}

	// configurationAddParams 添加配置参数
	configurationAddParams struct {
		Name      string              `json:"name,omitempty" validate:"required,xConfigurationName"`
		Category  confSchema.Category `json:"category,omitempty" validate:"required,xConfigurationCategory"`
		Status    schema.Status       `json:"status,omitempty" validate:"required,xStatus"`
		Data      string              `json:"data,omitempty" validate:"required,xConfigurationData"`
		StartedAt time.Time           `json:"startedAt,omitempty"`
		EndedAt   time.Time           `json:"endedAt,omitempty"`
	}
	// configurationUpdateParams 更新配置参数
	configurationUpdateParams struct {
		Name      string              `json:"name,omitempty" validate:"omitempty,xConfigurationName"`
		Status    schema.Status       `json:"status,omitempty" validate:"omitempty,xStatus"`
		Category  confSchema.Category `json:"category,omitempty" validate:"omitempty,xConfigurationCategory"`
		Data      string              `json:"data,omitempty" validate:"omitempty,xConfigurationData"`
		StartedAt time.Time           `json:"startedAt,omitempty"`
		EndedAt   time.Time           `json:"endedAt,omitempty"`
	}

	// configurationListParmas 配置查询参数
	configurationListParmas struct {
		listParams

		Name     string              `json:"name,omitempty" validate:"omitempty,xConfigurationName"`
		Category confSchema.Category `json:"category,omitempty" validate:"omitempty,xConfigurationCategory"`
	}
)

const (
	errConfigurationCategory = "configuration"
)

func init() {
	g := router.NewGroup("/configurations", loadUserSession, shouldBeSu)
	ctrl := configurationCtrl{}

	// 查询配置
	g.GET(
		"/v1",
		ctrl.list,
	)

	// 添加配置
	g.POST(
		"/v1",
		newTrackerMiddleware(cs.ActionConfigurationAdd),
		ctrl.add,
	)

	// 获取当前有效配置
	g.GET(
		"/v1/current-valid",
		ctrl.getCurrentValid,
	)

	// 更新配置
	g.PATCH(
		"/v1/{id}",
		newTrackerMiddleware(cs.ActionConfigurationUpdate),
		ctrl.update,
	)

	// 查询单个配置
	g.GET(
		"/v1/{id}",
		ctrl.findByID,
	)
}

// validateBeforeSave 保存前校验
func (params *configurationAddParams) validateBeforeSave(ctx context.Context) (err error) {
	exists, err := getEntClient().Configuration.Query().
		Where(configuration.Name(params.Name)).
		Exist(ctx)
	if err != nil {
		return
	}
	if exists {
		err = hes.New("该配置已存在", errConfigurationCategory)
		return
	}
	return
}

// save 保存配置
func (params *configurationAddParams) save(ctx context.Context, owner string) (configuration *ent.Configuration, err error) {
	err = params.validateBeforeSave(ctx)
	if err != nil {
		return
	}
	return getEntClient().Configuration.Create().
		SetName(params.Name).
		SetStatus(params.Status).
		SetCategory(params.Category).
		SetData(params.Data).
		SetOwner(owner).
		SetStartedAt(params.StartedAt).
		SetEndedAt(params.EndedAt).
		Save(ctx)
}

// where 将查询条件中的参数转换为对应的where条件
func (params *configurationListParmas) where(query *ent.ConfigurationQuery) *ent.ConfigurationQuery {
	if params.Name != "" {
		query = query.Where(configuration.Name(params.Name))
	}
	if params.Category != "" {
		query = query.Where(configuration.CategoryEQ(params.Category))
	}
	return query
}

// queryAll 查询配置列表
func (params *configurationListParmas) queryAll(ctx context.Context) (configurations []*ent.Configuration, err error) {
	query := getEntClient().Configuration.Query()

	query = query.Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Order(params.GetOrders()...)
	query = params.where(query)

	return query.All(ctx)
}

// count 计算总数
func (params *configurationListParmas) count(ctx context.Context) (count int, err error) {
	query := getEntClient().Configuration.Query()

	query = params.where(query)

	return query.Count(ctx)
}

// update 更新配置信息
func (params *configurationUpdateParams) updateOneID(ctx context.Context, id int) (configuration *ent.Configuration, err error) {
	updateOne := getEntClient().Configuration.
		UpdateOneID(id)
	if !params.StartedAt.IsZero() {
		updateOne = updateOne.SetStartedAt(params.StartedAt)
	}
	if !params.EndedAt.IsZero() {
		updateOne = updateOne.SetEndedAt(params.EndedAt)
	}
	if params.Name != "" {
		updateOne = updateOne.SetName(params.Name)
	}

	if params.Status != 0 {
		updateOne = updateOne.SetStatus(params.Status)
	}
	if params.Category != "" {
		updateOne = updateOne.SetCategory(params.Category)
	}
	if params.Data != "" {
		updateOne = updateOne.SetData(params.Data)
	}
	return updateOne.Save(ctx)
}

// add 添加配置
func (*configurationCtrl) add(c *elton.Context) (err error) {
	params := configurationAddParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	configuration, err := params.save(c.Context(), us.GetInfo().Account)
	if err != nil {
		return
	}
	c.Created(configuration)
	return
}

// list 查询配置列表
func (*configurationCtrl) list(c *elton.Context) (err error) {
	params := configurationListParmas{}
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
	configurations, err := params.queryAll(c.Context())
	if err != nil {
		return
	}
	c.Body = &configurationListResp{
		Count:          count,
		Configurations: configurations,
	}
	return
}

// update 更新配置信息
func (*configurationCtrl) update(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	params := configurationUpdateParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	configuration, err := params.updateOneID(c.Context(), id)
	if err != nil {
		return
	}

	c.Body = configuration
	return
}

// findByID 通过id查询
func (*configurationCtrl) findByID(c *elton.Context) (err error) {
	id, err := getIDFromParams(c)
	if err != nil {
		return
	}
	configuration, err := getEntClient().Configuration.Get(c.Context(), id)
	if err != nil {
		return
	}
	c.Body = configuration
	return
}

// getCurrentValid 获取当前有效配置
func (*configurationCtrl) getCurrentValid(c *elton.Context) (err error) {
	c.Body = service.GetCurrentValidConfiguration()
	return
}
