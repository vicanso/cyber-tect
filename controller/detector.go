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
	"fmt"
	"strconv"

	"github.com/vicanso/cybertect/cs"

	"github.com/lib/pq"
	"github.com/vicanso/cybertect/detector"
	"github.com/vicanso/cybertect/router"
	"github.com/vicanso/cybertect/validate"
	"github.com/vicanso/elton"
	"github.com/vicanso/hes"
)

var (
	dnsSrv  = new(detector.DNSSrv)
	tcpSrv  = new(detector.TCPSrv)
	pingSrv = new(detector.PingSrv)
	httpSrv = new(detector.HTTPSrv)
)

type detectorCtrl struct{}

type (
	addDetectorParams struct {
		Timeout     string        `json:"timeout,omitempty" validate:"xDuration"`
		Status      int           `json:"status,omitempty" validate:"xDetectorStatus,required"`
		Description string        `json:"description,omitempty" validate:"xDetectorDescription"`
		Receivers   pq.Int64Array `json:"receivers,omitempty" validate:"-"`
	}
	updateDetectorParams struct {
		Timeout     string        `json:"timeout,omitempty" validate:"xDuration"`
		Status      int           `json:"status,omitempty" validate:"xDetectorStatus"`
		Description string        `json:"description,omitempty" validate:"xDetectorDescription"`
		Receivers   pq.Int64Array `json:"receivers,omitempty" validate:"-"`
	}
	addDNSParams struct {
		addDetectorParams

		Server   string `json:"server,omitempty" validate:"xDNSServer,required"`
		Hostname string `json:"hostname,omitempty" validate:"xDNSHostname,required"`
	}
	updateDNSParams struct {
		updateDetectorParams

		Server   string `json:"server,omitempty" validate:"xDNSServer"`
		Hostname string `json:"hostname,omitempty" validate:"xDNSHostname"`
	}
	addTCPParams struct {
		addDetectorParams

		Network string `json:"network,omitempty" validate:"xTCPNetwork,required"`
		IP      string `json:"ip,omitempty" validate:"ip,required"`
		Port    int    `json:"port,omitempty" validate:"port,required"`
	}
	updateTCPParams struct {
		updateDetectorParams

		Network string `json:"network,omitempty" validate:"xTCPNetwork"`
		IP      string `json:"ip,omitempty" validate:"ip"`
		Port    int    `json:"port,omitempty" validate:"port"`
	}
	addPingParams struct {
		addDetectorParams

		Network string `json:"network,omitempty" validate:"xTCPNetwork,required"`
		IP      string `json:"ip,omitempty" validate:"ip,required"`
	}
	updatePingParams struct {
		updateDetectorParams

		Network string `json:"network,omitempty" validate:"xTCPNetwork"`
		IP      string `json:"ip,omitempty" validate:"ip"`
	}
	addHTTPParams struct {
		addDetectorParams

		URL string `json:"url,omitempty" validate:"url,required"`
		IP  string `json:"ip,omitempty" validate:"ip"`
	}
	updateHTTPParams struct {
		updateDetectorParams

		URL string `json:"url,omitempty" validate:"url"`
		IP  string `json:"ip,omitempty" validate:"ip"`
	}
)

const (
	catDNS  = "dns"
	catTCP  = "tcp"
	catPing = "ping"
	catHTTP = "http"
)

func init() {
	ctrl := detectorCtrl{}
	g := router.NewGroup("/detectors", loadUserSession, shouldLogined)

	g.POST(
		"/v1/{category}",
		newTracker(cs.ActionDetectorAdd),
		ctrl.add,
	)
	g.PATCH(
		"/v1/{category}/{id}",
		newTracker(cs.ActionDetectorUpdate),
		ctrl.update,
	)
}

// addDNSDetector 添加dns detector
func addDNSDetector(c *elton.Context) (data *detector.DNS, err error) {
	params := addDNSParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	data = &detector.DNS{
		Owner:       us.GetAccount(),
		Server:      params.Server,
		Hostname:    params.Hostname,
		Timeout:     params.Timeout,
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
	}
	err = dnsSrv.Add(data)
	return
}

// updateDNSDetector 更新dns detector
func updateDNSDetector(id uint, c *elton.Context) (err error) {
	us := getUserSession(c)
	err = dnsSrv.ValidateOwner(id, us.GetAccount())
	if err != nil {
		return
	}
	params := updateDNSParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	err = dnsSrv.UpdateByID(id, &detector.DNS{
		Server:      params.Server,
		Hostname:    params.Hostname,
		Timeout:     params.Timeout,
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
	})
	return
}

// addTCPDetector add tcp detector
func addTCPDetector(c *elton.Context) (data *detector.TCP, err error) {
	params := addTCPParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	data = &detector.TCP{
		Owner:       us.GetAccount(),
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		Network:     params.Network,
		IP:          params.IP,
		Port:        params.Port,
	}
	err = tcpSrv.Add(data)
	return
}

// updateTCPDetector update tcp detector
func updateTCPDetector(id uint, c *elton.Context) (err error) {
	us := getUserSession(c)
	err = tcpSrv.ValidateOwner(id, us.GetAccount())
	if err != nil {
		return
	}
	params := updateTCPParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	err = tcpSrv.UpdateByID(id, &detector.TCP{
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		Network:     params.Network,
		IP:          params.IP,
		Port:        params.Port,
	})
	return
}

// addPingDetector add ping detector
func addPingDetector(c *elton.Context) (data *detector.Ping, err error) {
	params := addPingParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	data = &detector.Ping{
		Owner:       us.GetAccount(),
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		Network:     params.Network,
		IP:          params.IP,
	}
	err = pingSrv.Add(data)
	return
}

// updatePingDetector update ping detector
func updatePingDetector(id uint, c *elton.Context) (err error) {
	us := getUserSession(c)
	err = pingSrv.ValidateOwner(id, us.GetAccount())
	if err != nil {
		return
	}
	params := updatePingParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	err = pingSrv.UpdateByID(id, &detector.Ping{
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		Network:     params.Network,
		IP:          params.IP,
	})
	return
}

// addHTTPDetector add http detector
func addHTTPDetector(c *elton.Context) (data *detector.HTTP, err error) {
	params := addHTTPParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	data = &detector.HTTP{
		Owner:       us.GetAccount(),
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		IP:          params.IP,
		URL:         params.URL,
	}
	err = httpSrv.Add(data)
	return
}

// updateHTTPDetector update http detector
func updateHTTPDetector(id uint, c *elton.Context) (err error) {
	us := getUserSession(c)
	err = httpSrv.ValidateOwner(id, us.GetAccount())
	if err != nil {
		return
	}
	params := updateHTTPParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	updateParams := &detector.HTTP{
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		IP:          params.IP,
		URL:         params.URL,
	}
	needToRemoveIP := false
	// 如果是0.0.0.0表示删除
	if updateParams.IP == "0.0.0.0" {
		updateParams.IP = ""
		needToRemoveIP = true
	}
	err = httpSrv.UpdateByID(id, updateParams)
	if err != nil {
		return
	}
	// TODO 后续再确认有无其它方法，暂时使用此方式（并未保证事务）
	if needToRemoveIP {
		err = httpSrv.UpdateByID(id, "ip", "")
	}
	return
}

func (ctrl detectorCtrl) add(c *elton.Context) (err error) {
	cat := c.Param("category")
	var data interface{}
	switch cat {
	case catDNS:
		data, err = addDNSDetector(c)
	case catTCP:
		data, err = addTCPDetector(c)
	case catPing:
		data, err = addPingDetector(c)
	case catHTTP:
		data, err = addHTTPDetector(c)
	default:
		err = hes.New(fmt.Sprintf("Not support category:%s", cat))
	}

	if err != nil {
		return
	}
	c.Created(data)
	return
}

func (ctrl detectorCtrl) update(c *elton.Context) (err error) {
	cat := c.Param("category")
	v, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	id := uint(v)
	switch cat {
	case catDNS:
		err = updateDNSDetector(id, c)
	case catTCP:
		err = updateTCPDetector(id, c)
	case catPing:
		err = updatePingDetector(id, c)
	case catHTTP:
		err = updateHTTPDetector(id, c)
	default:
		err = hes.New(fmt.Sprintf("Not support category:%s", cat))
	}

	if err != nil {
		return
	}
	c.NoContent()
	return
}
