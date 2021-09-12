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
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/vicanso/elton"
	M "github.com/vicanso/elton/middleware"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/helper"
	"github.com/vicanso/cybertect/interceptor"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/middleware"
	"github.com/vicanso/cybertect/schema"
	"github.com/vicanso/cybertect/service"
	"github.com/vicanso/cybertect/session"
	"github.com/vicanso/cybertect/util"
	"github.com/vicanso/hes"
)

type listParams = helper.EntListParams

var (
	// getEntClient = helper.EntGetClient
	now = util.NowString

	getUserSession = session.NewUserSession
	// 加载用户session
	loadUserSession = elton.Compose(session.New(), sessionHandle)
	// 判断用户是否登录
	shouldBeLogin = checkLoginMiddleware
	// 判断用户是否未登录
	shouldBeAnonymous = checkAnonymousMiddleware
	// 判断用户是否admin权限
	shouldBeAdmin = newCheckRolesMiddleware([]string{
		schema.UserRoleSu,
		schema.UserRoleAdmin,
	})
	// shouldBeSu 判断用户是否su权限
	shouldBeSu = newCheckRolesMiddleware([]string{
		schema.UserRoleSu,
	})

	// 创建新的并发控制中间件
	newConcurrentLimit = middleware.NewConcurrentLimit
	// 创建IP限制中间件
	newIPLimit = middleware.NewIPLimit
	// 创建出错限制中间件
	newErrorLimit = middleware.NewErrorLimit
	// noCacheIfRequestNoCache 请求参数指定no cache，则设置no-cache
	noCacheIfRequestNoCache = middleware.NewNoCacheWithCondition("cacheControl", "no-cache")

	// 图形验证码校验
	captchaValidate = newMagicalCaptchaValidate()
	// GetInfluxDB 仅提供基础服务
	GetInfluxDB = helper.GetInfluxDB
	// 获取influx service
	GetInfluxSrv = service.GetInfluxSrv
	// 文件服务
	fileSrv = &service.FileSrv{}
)

type (
	trackerExtraParams struct {
		// 步骤（tag)
		Step string
	}
)

func getUserClient() *ent.UserClient {
	return helper.EntGetClient().User
}

func getUserLoginClient() *ent.UserLoginClient {
	return helper.EntGetClient().UserLogin
}

func getConfigurationClient() *ent.ConfigurationClient {
	return helper.EntGetClient().Configuration
}

func newMagicalCaptchaValidate() elton.Handler {
	magicValue := ""
	if !util.IsProduction() {
		magicValue = "0145"
	}
	return middleware.ValidateCaptcha(magicValue)
}

// isLogin 判断是否登录状态
func isLogin(c *elton.Context) bool {
	us := session.NewUserSession(c)
	return us.IsLogin()
}

func validateLogin(c *elton.Context) error {
	if !isLogin(c) {
		return hes.New("请先登录", errUserCategory)
	}
	return nil
}

// checkLoginMiddleware 校验是否登录中间件
func checkLoginMiddleware(c *elton.Context) error {
	err := validateLogin(c)
	if err != nil {
		return err
	}
	return c.Next()
}

// checkAnonymousMiddleware 判断是匿名状态
func checkAnonymousMiddleware(c *elton.Context) error {
	if isLogin(c) {
		return hes.New("已是登录状态，请先退出登录", errUserCategory)
	}
	return c.Next()
}

// newCheckRolesMiddleware 创建用户角色校验中间件
func newCheckRolesMiddleware(validRoles []string) elton.Handler {
	return func(c *elton.Context) error {
		err := validateLogin(c)
		if err != nil {
			return err
		}
		us := session.NewUserSession(c)
		userInfo, err := us.GetInfo()
		if err != nil {
			return err
		}
		valid := util.ContainsAny(validRoles, userInfo.Roles)
		if valid {
			return c.Next()
		}
		return hes.NewWithStatusCode("禁止使用该功能", http.StatusForbidden, errUserCategory)
	}
}

// newTrackerMiddleware 初始化用户行为跟踪中间件
func newTrackerMiddleware(action string, params ...trackerExtraParams) elton.Handler {
	marshalString := func(data interface{}) string {
		buf, _ := json.Marshal(data)
		return string(buf)
	}
	var extraParams *trackerExtraParams
	if len(params) != 0 {
		extraParams = &params[0]
	}
	return M.NewTracker(M.TrackerConfig{
		Mask:      cs.MaskRegExp,
		MaxLength: 30,
		OnTrack: func(info *M.TrackerInfo, c *elton.Context) {
			account := ""
			tid := util.GetDeviceID(c.Context())
			us := session.NewUserSession(c)
			if us != nil && us.IsLogin() {
				account = us.MustGetInfo().Account
			}
			ip := c.RealIP()
			sid := util.GetSessionID(c)

			fields := map[string]interface{}{
				cs.FieldAccount: account,
				cs.FieldIP:      ip,
				cs.FieldSID:     sid,
				cs.FieldTID:     tid,
			}
			if len(info.Query) != 0 {
				fields[cs.FieldQuery] = marshalString(info.Query)
			}
			if len(info.Params) != 0 {
				fields[cs.FieldParams] = marshalString(info.Params)
			}
			if len(info.Form) != 0 {
				fields[cs.FieldForm] = marshalString(info.Form)
			}
			if info.Err != nil {
				fields[cs.FieldError] = info.Err.Error()
			}
			currentStep := ""
			if extraParams != nil {
				currentStep = extraParams.Step
			}
			event := log.Info(c.Context()).
				Str("category", "tracker").
				Str("action", action).
				Str("ip", ip).
				Str("sid", sid).
				Int("result", info.Result)
			if currentStep != "" {
				event = event.Str("step", currentStep)
			}
			if len(info.Query) != 0 {
				event = event.Dict("query", log.Struct(info.Query))
			}
			if len(info.Params) != 0 {
				event = event.Dict("params", log.Struct(info.Params))
			}
			if len(info.Form) != 0 {
				event = event.Dict("form", zerolog.
					Dict().
					Fields(info.Form))
			}
			event.Err(info.Err).
				Msg("")
			tags := map[string]string{
				cs.TagAction: action,
				cs.TagResult: strconv.Itoa(info.Result),
			}
			if currentStep != "" {
				tags["step"] = currentStep
			}
			GetInfluxSrv().Write(cs.MeasurementUserTracker, tags, fields)
		},
	})
}

// getIDFromParams get id form context params
func getIDFromParams(c *elton.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		he := hes.Wrap(err)
		he.Category = "parseInt"
		return 0, he
	}
	return id, nil
}

// sessionHandle session的相关处理
func sessionHandle(c *elton.Context) error {
	interData, _ := interceptor.GetSessionData()

	us := session.NewUserSession(c)
	account := ""
	if us.IsLogin() {
		account = us.MustGetInfo().Account
	}

	// 设置账号信息
	c.WithContext(util.SetAccount(c.Context(), account))

	// 如果无配置，则直接跳过
	if interData == nil {
		return c.Next()
	}

	// 如果配置该账号允许
	if account != "" && util.ContainsString(interData.AllowAccounts, account) {
		return c.Next()
	}
	// 如果路由配置允许
	if util.ContainsString(interData.AllowRoutes, c.Route) {
		return c.Next()
	}

	// 如果有配置拦截信息，则以出错返回
	he := hes.New(interData.Message)
	he.Category = "sessionInterceptorMiddleware"
	return he
}

// isIntranet 判断是否内网访问
func isIntranet(c *elton.Context) error {
	if elton.IsIntranet(c.ClientIP()) {
		return c.Next()
	}
	return hes.NewWithStatusCode("Forbidden", 403)
}
