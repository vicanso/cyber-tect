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
	"time"

	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/user"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
)

type detectorCtrl struct{}

const errDetectorCategory = "detector"

var errInvalidUser = hes.New("无修改该配置的权限", errDetectorCategory)

type (
	// detectorAddParams detector add params
	detectorAddParams struct {
		// 配置名称
		Name string `json:"name" validate:"required,xDetectorName"`
		// 配置状态
		Status schema.Status `json:"status" validate:"required,xStatus"`
		// 配置描述
		Description string `json:"description" validate:"required,xDetectorDesc"`
		// 告警接收者
		Receivers []string `json:"receivers" validate:"required,dive,xUserAccount"`
		// 拥有者
		Owners []string `json:"owners" validate:"required,dive,xUserAccount"`
		// 超时设置
		Timeout string `json:"timeout" validate:"required,xDuration"`
	}

	// detectorUpdateParams detector update params
	detectorUpdateParams struct {
		Name        string        `json:"name" validate:"omitempty,xDetectorName"`
		Status      schema.Status `json:"status" validate:"omitempty,xStatus"`
		Description string        `json:"description" validate:"omitempty,xDetectorDesc"`
		Receivers   []string      `json:"receivers" validate:"omitempty,dive,xUserAccount"`
		Timeout     string        `json:"timeout" validate:"omitempty,xDuration"`
		// 拥有者
		Owners []string `json:"owners" validate:"omitempty,dive,xUserAccount"`
	}

	detectorListUserParams struct {
		listParams

		// 关键字搜索
		// pattern: xKeyword
		Keyword string `json:"keyword" validate:"omitempty,xKeyword"`
	}

	// detectorListHTTPResultParams params of list http result
	detectorListResultParams struct {
		listParams `json:"listParams"`

		// 任务列表，通过当前登录账号获取
		tasks []int

		Result    int8      `json:"result" validate:"omitempty,xDetectorResult"`
		Duration  string    `json:"duration" validate:"omitempty,xDuration"`
		StartedAt time.Time `json:"startedAt"`
		EndedAt   time.Time `json:"endedAt"`
	}
)

func init() {
	ctrl := detectorCtrl{}
	g := router.NewGroup(
		"/detectors",
		loadUserSession,
		shouldBeLogin,
	)

	// 用户查询
	g.GET("/users/v1", ctrl.listUser)
}

// GetDurationMs
func (params *detectorListResultParams) GetDurationMillSecond() int {
	if params.Duration == "" {
		return 0
	}
	d, _ := time.ParseDuration(params.Duration)
	return int(d.Milliseconds())
}

func (listUserParams *detectorListUserParams) queryAll(ctx context.Context) ([]*ent.User, error) {
	query := getUserClient().Query()

	query = query.Limit(listUserParams.GetLimit()).
		Offset(listUserParams.GetOffset()).
		Order(listUserParams.GetOrders()...)
	if listUserParams.Keyword != "" {
		query = query.Where(user.AccountContains(listUserParams.Keyword))
	}

	return query.All(ctx)
}

func (*detectorCtrl) listUser(c *elton.Context) error {
	params := detectorListUserParams{}
	params.Fields = "account"
	err := validateQuery(c, &params)
	if err != nil {
		return err
	}
	users, err := params.queryAll(c.Context())
	if err != nil {
		return err
	}
	accounts := make([]string, len(users))
	for index, u := range users {
		accounts[index] = u.Account
	}
	c.Body = map[string][]string{
		"accounts": accounts,
	}

	return nil
}
