// Copyright 2019 tree xie
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
	"time"

	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/service"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
)

type (
	configurationCtrl      struct{}
	addConfigurationParams struct {
		Name      string     `json:"name,omitempty" validate:"xConfigName,required"`
		Category  string     `json:"category,omitempty" validate:"xConfigCategory"`
		Status    int        `json:"status,omitempty" validate:"xConfigStatus,required"`
		Data      string     `json:"data,omitempty" validate:"xConfigData,required"`
		BeginDate *time.Time `json:"beginDate,omitempty" validate:"-"`
		EndDate   *time.Time `json:"endDate,omitempty" validate:"-"`
	}
	updateConfigurationParams struct {
		Status    int        `json:"status,omitempty" validate:"xConfigStatus"`
		Category  string     `json:"category,omitempty" validate:"xConfigCategory"`
		Data      string     `json:"data,omitempty" validate:"xConfigData"`
		BeginDate *time.Time `json:"beginDate" validate:"-"`
		EndDate   *time.Time `json:"endDate" validate:"-"`
	}
	listConfigurationParmas struct {
		Name     string `json:"name,omitempty" validate:"xConfigName"`
		Category string `json:"category,omitempty" validate:"xConfigCategory"`
	}
)

func init() {
	// TODO 增加用户权限判断
	g := router.NewGroup("/configurations", loadUserSession)
	ctrl := configurationCtrl{}

	g.GET(
		"/v1",
		shouldBeAdmin,
		ctrl.list,
	)
	g.GET(
		"/v1/available",
		shouldBeAdmin,
		ctrl.listAvailable,
	)
	g.GET(
		"/v1/unavailable",
		shouldBeAdmin,
		ctrl.listUnavailable,
	)

	g.POST(
		"/v1",
		newTracker(cs.ActionConfigurationAdd),
		shouldBeAdmin,
		ctrl.add,
	)
	g.PATCH(
		"/v1/{configID}",
		newTracker(cs.ActionConfigurationUpdate),
		shouldBeAdmin,
		ctrl.update,
	)
	g.DELETE(
		"/v1/{configID}",
		newTracker(cs.ActionConfigurationDelete),
		shouldBeAdmin,
		ctrl.delete,
	)
}

// list configuration
func (ctrl configurationCtrl) list(c *elton.Context) (err error) {
	params := &listConfigurationParmas{}
	err = validate.Do(params, c.Query())
	if err != nil {
		return
	}
	result, err := configSrv.List(service.ConfigurationQueryParmas{
		Name:     params.Name,
		Category: params.Category,
	})
	if err != nil {
		return
	}
	c.Body = map[string]interface{}{
		"configs": result,
	}
	return
}

// listAvailable list available config
func (ctrl configurationCtrl) listAvailable(c *elton.Context) (err error) {
	result, err := configSrv.Available()
	if err != nil {
		return
	}
	c.Body = map[string]interface{}{
		"configs": result,
	}
	return
}

// listUnavailable list unavailable config
func (ctrl configurationCtrl) listUnavailable(c *elton.Context) (err error) {
	result, err := configSrv.Unavailable()
	if err != nil {
		return
	}
	c.Body = map[string]interface{}{
		"configs": result,
	}
	return
}

// add configuration
func (ctrl configurationCtrl) add(c *elton.Context) (err error) {
	params := &addConfigurationParams{}
	err = validate.Do(params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	conf := &service.Configuration{
		Name:      params.Name,
		Category:  params.Category,
		Status:    params.Status,
		Data:      params.Data,
		Owner:     us.GetAccount(),
		BeginDate: params.BeginDate,
		EndDate:   params.EndDate,
	}
	err = configSrv.Add(conf)
	if err != nil {
		return
	}
	c.Created(conf)
	return
}

// update configuration
func (ctrl configurationCtrl) update(c *elton.Context) (err error) {
	id, err := strconv.Atoi(c.Param("configID"))
	if err != nil {
		return
	}
	params := &updateConfigurationParams{}
	err = validate.Do(params, c.RequestBody)
	if err != nil {
		return
	}
	err = configSrv.Update(service.Configuration{
		ID: uint(id),
	}, service.Configuration{
		Status:    params.Status,
		Data:      params.Data,
		Category:  params.Category,
		BeginDate: params.BeginDate,
		EndDate:   params.EndDate,
	})
	if err != nil {
		return
	}

	c.NoContent()
	return
}

// delete configuration
func (ctrl configurationCtrl) delete(c *elton.Context) (err error) {
	id, err := strconv.Atoi(c.Param("configID"))
	if err != nil {
		return
	}
	err = configSrv.DeleteByID(uint(id))
	if err != nil {
		return
	}
	c.NoContent()
	return
}
