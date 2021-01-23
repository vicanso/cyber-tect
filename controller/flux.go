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

// flux查询influxdb相关数据

package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
)

type (
	fluxCtrl struct{}

	// fluxListParams flux查询参数
	fluxListParams struct {
		Measurement string    `json:"measurement,omitempty"`
		Begin       time.Time `json:"begin,omitempty" validate:"required"`
		End         time.Time `json:"end,omitempty" validate:"required"`
		Account     string    `json:"account,omitempty" validate:"omitempty,xUserAccount"`
		Limit       string    `json:"limit,omitempty" validate:"required,xLargerLimit"`
		Exception   string    `json:"exception,omitempty" validate:"omitempty,xBoolean"`
		// 用户行为类型筛选
		Action   string `json:"action,omitempty" validate:"omitempty,xTag"`
		Result   string `json:"result,omitempty" validate:"omitempty,xTag"`
		Category string `json:"category,omitempty" validate:"omitempty,xTag"`
	}
	// fluxListTagValuesParams flux tag values查询参数
	fluxListTagValuesParams struct {
		Measurement string `json:"measurement,omitempty" validate:"required,xMeasurement"`
		Tag         string `json:"tag,omitempty" validate:"required,xTag"`
	}
)

func init() {
	sessionConfig = config.GetSessionConfig()
	g := router.NewGroup("/fluxes", loadUserSession)

	ctrl := fluxCtrl{}
	// 查询用户tracker
	g.GET(
		"/v1/trackers",
		shouldBeAdmin,
		ctrl.listTracker,
	)
	// 查询http出错
	g.GET(
		"/v1/http-errors",
		shouldBeAdmin,
		ctrl.listHTTPError,
	)
	// 获取tag的值
	g.GET(
		"/v1/tag-values/{measurement}/{tag}",
		shouldBeAdmin,
		ctrl.listTagValue,
	)
}

// Query get flux query string
func (params *fluxListParams) Query() string {
	start := util.FormatTime(params.Begin.UTC())
	stop := util.FormatTime(params.End.UTC())
	query := fmt.Sprintf(`
		|> range(start: %s, stop: %s)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> sort(columns:["_time"], desc: true)
		|> limit(n:%s)
		|> pivot(
			rowKey:["_time"],
			columnKey: ["_field"],
			valueColumn: "_value"
		)
		`,
		start,
		stop,
		params.Measurement,
		params.Limit,
	)
	// 用户行为类型
	if params.Action != "" {
		query += fmt.Sprintf(`|> filter(fn: (r) => r.action == "%s")`, params.Action)
	}
	// 结果
	if params.Result != "" {
		query += fmt.Sprintf(`|> filter(fn: (r) => r.result == "%s")`, params.Result)
	}
	if params.Category != "" {
		query += fmt.Sprintf(`|> filter(fn: (r) => r.category == "%s")`, params.Category)
	}
	// 账号
	if params.Account != "" {
		query += fmt.Sprintf(`|> filter(fn: (r) => r.account == "%s")`, params.Account)
	}
	// 异常
	if params.Exception != "" {
		value := "true"
		if params.Exception == "0" {
			value = "false"
		}
		query += fmt.Sprintf(`|> filter(fn: (r) => r.exception == %s)`, value)
	}
	return query
}

func (params *fluxListParams) Do(ctx context.Context) (items []map[string]interface{}, err error) {
	items, err = getInfluxSrv().Query(ctx, params.Query())
	if err != nil {
		return
	}
	// 清除不需要字段
	for _, item := range items {
		delete(item, "_measurement")
		delete(item, "_start")
		delete(item, "_stop")
		delete(item, "table")
	}
	return
}

// listValue get the values of tag
func (ctrl fluxCtrl) listTagValue(c *elton.Context) (err error) {
	params := fluxListTagValuesParams{}
	err = validate.Do(&params, c.Params.ToMap())
	if err != nil {
		return
	}
	values, err := getInfluxSrv().ListTagValue(c.Context(), params.Measurement, params.Tag)
	if err != nil {
		return
	}
	c.Body = map[string][]string{
		"values": values,
	}
	return
}

// listHTTPError list http error
func (ctrl fluxCtrl) listHTTPError(c *elton.Context) (err error) {
	params := fluxListParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	params.Measurement = cs.MeasurementHTTPError
	result, err := params.Do(c.Context())
	c.Body = map[string]interface{}{
		"httpErrors": result,
	}
	return
}

// listTracker list user tracker
func (ctrl fluxCtrl) listTracker(c *elton.Context) (err error) {
	params := fluxListParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}
	params.Measurement = cs.MeasurementUserTracker

	result, err := params.Do(c.Context())
	c.Body = map[string]interface{}{
		"trackers": result,
	}
	return
}
