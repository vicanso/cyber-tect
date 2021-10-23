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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/thoas/go-funk"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/databasedetector"
	"github.com/vicanso/cybertect/ent/dnsdetector"
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/pingdetector"
	"github.com/vicanso/cybertect/ent/tcpdetector"
	"github.com/vicanso/cybertect/ent/user"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/session"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
)

type detectorCtrl struct{}

const errDetectorCategory = "detector"

const (
	detectorCategoryDatabase = "database"
	detectorCategoryDNS      = "dns"
	detectorCategoryTCP      = "tcp"
	detectorCategoryPing     = "ping"
	detectorCategoryHTTP     = "http"
)

var errInvalidUser = hes.New("无修改该配置的权限", errDetectorCategory)
var errTaskNotFound = hes.New("task not found")

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
		Tasks     []int
		Result    int8      `json:"result" validate:"omitempty,xDetectorResult"`
		Duration  string    `json:"duration" validate:"omitempty,xDuration"`
		StartedAt time.Time `json:"startedAt"`
		EndedAt   time.Time `json:"endedAt"`
	}
	getResultSummaryParams struct {
		StartedAt time.Time `json:"startedAt"`
	}
)

type (
	detectorResultSummary struct {
		Category string                `json:"category"`
		Result   schema.DetectorResult `json:"result"`
		Count    int                   `json:"count"`
	}
	getResultSummaryResp struct {
		Summaries []*detectorResultSummary `json:"summaries" validate:"required"`
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

	// 获取结果汇总
	g.GET("/result-summaries/v1", ctrl.getResultSummary)
}

// GetDurationMs
func (params *detectorListResultParams) GetDurationMillSecond() int {
	if params.Duration == "" {
		return 0
	}
	d, _ := time.ParseDuration(params.Duration)
	return int(d.Milliseconds())
}

func getDetectorTasksByReceiver(ctx context.Context, category string, us *session.UserSession) ([]int, error) {
	// 超级用户不限制
	if us.IsAdmin() {
		return nil, nil
	}
	receiver := us.MustGetInfo().Account
	newFilter := func(field string) func(*sql.Selector) {
		return func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(field, receiver))
		}
	}
	fields := "id"
	var arr interface{}
	var err error
	switch category {
	case detectorCategoryDatabase:
		arr, err = getDatabaseDetectorClient().Query().
			Where(newFilter(databasedetector.FieldReceivers)).
			Select(fields).
			All(ctx)
	case detectorCategoryDNS:
		arr, err = getDNSDetectorClient().Query().
			Where(newFilter(dnsdetector.FieldReceivers)).
			Select(fields).
			All(ctx)
	case detectorCategoryHTTP:
		arr, err = getHTTPDetectorClient().Query().
			Where(newFilter(httpdetector.FieldReceivers)).
			Select(fields).
			All(ctx)
	case detectorCategoryPing:
		arr, err = getPingDetectorClient().Query().
			Where(newFilter(pingdetector.FieldReceivers)).
			Select(fields).
			All(ctx)
	case detectorCategoryTCP:
		arr, err = getTCPDetectorClient().Query().
			Where(newFilter(tcpdetector.FieldReceivers)).
			Select(fields).
			All(ctx)

	default:
		return nil, hes.New(category + "类型错误")
	}
	if err != nil {
		return nil, err
	}
	tasks := make([]int, 0)
	funk.ForEach(arr, func(item interface{}) {
		id := -1
		switch data := item.(type) {
		case *ent.DatabaseDetector:
			id = data.ID
		case *ent.HTTPDetector:
			id = data.ID
		case *ent.DNSDetector:
			id = data.ID
		case *ent.PingDetector:
			id = data.ID
		case *ent.TCPDetector:
			id = data.ID
		}
		tasks = append(tasks, id)
	})
	if len(tasks) == 0 {
		return nil, errTaskNotFound
	}
	return tasks, nil
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

func (*detectorCtrl) getResultSummary(c *elton.Context) error {
	queryParams := getResultSummaryParams{}
	err := validateQuery(c, &queryParams)
	if err != nil {
		return err
	}
	startedAt := queryParams.StartedAt
	if time.Since(startedAt) > 10*24*time.Hour {
		return hes.New("汇总时间过长")
	}
	us := getUserSession(c)
	categories := []string{
		detectorCategoryDatabase,
		detectorCategoryHTTP,
		detectorCategoryTCP,
		detectorCategoryPing,
		detectorCategoryDNS,
	}
	summaries := make([]*detectorResultSummary, 0)
	for _, category := range categories {
		tasks, err := getDetectorTasksByReceiver(
			c.Context(),
			category,
			us,
		)
		if err == errTaskNotFound {
			continue
		}
		if err != nil {
			return err
		}
		for _, v := range []schema.DetectorResult{
			schema.DetectorResultSuccess,
			schema.DetectorResultFail,
		} {
			result := int8(v)
			var count int
			switch category {
			case detectorCategoryDatabase:
				params := databaseDetectorResultListParams{}
				params.Tasks = tasks
				params.StartedAt = startedAt
				params.Result = result
				count, err = params.count(c.Context())
				if err != nil {
					return err
				}
			case detectorCategoryDNS:
				params := dnsDetectorResultListParams{}
				params.Tasks = tasks
				params.StartedAt = startedAt
				params.Result = result
				count, err = params.count(c.Context())
				if err != nil {
					return err
				}
			case detectorCategoryHTTP:
				params := httpDetectorResultListParams{}
				params.Tasks = tasks
				params.StartedAt = startedAt
				params.Result = result
				count, err = params.count(c.Context())
				if err != nil {
					return err
				}
			case detectorCategoryPing:
				params := pingDetectorResultListParams{}
				params.Tasks = tasks
				params.StartedAt = startedAt
				params.Result = result
				count, err = params.count(c.Context())
				if err != nil {
					return err
				}
			case detectorCategoryTCP:
				params := tcpDetectorResultListParams{}
				params.Tasks = tasks
				params.StartedAt = startedAt
				params.Result = result
				count, err = params.count(c.Context())
				if err != nil {
					return err
				}
			}
			summaries = append(summaries, &detectorResultSummary{
				Category: category,
				Result:   v,
				Count:    count,
			})
		}
	}

	c.Body = &getResultSummaryResp{
		Summaries: summaries,
	}

	return nil
}
