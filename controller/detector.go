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
	"strconv"

	"github.com/vicanso/cybertect/ent/schema"
	"github.com/vicanso/cybertect/ent/user"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
)

type (
	detectorCtrl struct{}

	// detectorAddParams detector add params
	detectorAddParams struct {
		Name        string   `json:"name,omitempty" validate:"required,xDetectorName"`
		Status      int      `json:"status,omitempty" validate:"required,xStatus"`
		Description string   `json:"description,omitempty" validate:"required,xDetectorDesc"`
		Receivers   []string `json:"receivers,omitempty" validate:"required,dive,xUserAccount"`
		Timeout     string   `json:"timeout,omitempty" validate:"required,xDuration"`
	}
	// detectorUpdateParams detector update params
	detectorUpdateParams struct {
		CurrentUser string

		Name        string   `json:"name,omitempty" validate:"omitempty,xDetectorName"`
		Status      int      `json:"status,omitempty" validate:"omitempty,xStatus"`
		Description string   `json:"description,omitempty" validate:"omitempty,xDetectorDesc"`
		Receivers   []string `json:"receivers,omitempty" validate:"omitempty,dive,xUserAccount"`
		Timeout     string   `json:"timeout,omitempty" validate:"omitempty,xDuration"`
	}

	// detectorListHTTPResultParams params of list http result
	detectorListResultParams struct {
		listParams

		Task   string `json:"task,omitempty" validate:"omitempty,xDetectorTaskID"`
		Result string `json:"result,omitempty" validate:"omitempty,xDetectorResult"`
	}
)

const (
	errDetectorCategory = "detector"
)

var errInvalidUser = hes.New("仅能创建者允许修改配置", errDetectorCategory)

func init() {
	g := router.NewGroup("/detectors", loadUserSession, shouldBeLogin)

	ctrl := detectorCtrl{}

	// 查询接收者列表
	g.GET(
		"/v1/receivers",
		ctrl.listReceiver,
	)
}

// GetTaskID get task id
func (params *detectorListResultParams) GetTaskID() int {
	// 参数已校验是数字，因此转换时不判断
	id, _ := strconv.Atoi(params.Task)
	return id
}

// GetResult get result
func (params *detectorListResultParams) GetResult() int8 {
	// 参数已校验，因此转换不判断
	result, _ := strconv.Atoi(params.Result)
	return int8(result)
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
