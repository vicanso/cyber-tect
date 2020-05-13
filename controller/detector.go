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
	"strings"
	"time"

	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/helper"

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
	queryDetectorParams struct {
		Limit  string `json:"limit,omitempty" validate:"xLimit"`
		Offset string `json:"offset,omitempty" validate:"xOffset"`
		Order  string `json:"order,omitempty"`
	}
	queryDetectorResultParams struct {
		Limit    string `json:"limit,omitempty" validate:"xLimit"`
		Offset   string `json:"offset,omitempty" validate:"xOffset"`
		Order    string `json:"order,omitempty"`
		Task     string `json:"task,omitempty" validate:"omitempty,xDetectorTask"`
		Result   string `json:"result,omitempty" validate:"omitempty,xDetectorResult"`
		Duration string `json:"duration,omitempty" validate:"omitempty,xDuration"`
	}
	addDetectorParams struct {
		Timeout     string         `json:"timeout,omitempty" validate:"omitempty,xDuration"`
		Status      int            `json:"status,omitempty" validate:"xDetectorStatus,required"`
		Description string         `json:"description,omitempty" validate:"xDetectorDescription"`
		Receivers   pq.StringArray `json:"receivers,omitempty" validate:"required"`
	}
	updateDetectorParams struct {
		Timeout     string         `json:"timeout,omitempty" validate:"omitempty,xDuration"`
		Status      int            `json:"status,omitempty" validate:"xDetectorStatus"`
		Description string         `json:"description,omitempty" validate:"xDetectorDescription"`
		Receivers   pq.StringArray `json:"receivers,omitempty" validate:"-"`
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

		Network string `json:"network,omitempty" validate:"xTCPNetwork"`
		IP      string `json:"ip,omitempty" validate:"ip,required"`
		Port    int    `json:"port,omitempty" validate:"required"`
	}
	updateTCPParams struct {
		updateDetectorParams

		Network string `json:"network,omitempty" validate:"xTCPNetwork"`
		IP      string `json:"ip,omitempty" validate:"isdefault|ip"`
		Port    int    `json:"port,omitempty"`
	}
	addPingParams struct {
		addDetectorParams

		Network string `json:"network,omitempty" validate:"xTCPNetwork"`
		IP      string `json:"ip,omitempty" validate:"ip,required"`
	}
	updatePingParams struct {
		updateDetectorParams

		Network string `json:"network,omitempty" validate:"xTCPNetwork"`
		IP      string `json:"ip,omitempty" validate:"isdefault|ip"`
	}
	addHTTPParams struct {
		addDetectorParams

		URL string `json:"url,omitempty" validate:"url,required"`
		IP  string `json:"ip,omitempty" validate:"isdefault|ip"`
	}
	updateHTTPParams struct {
		updateDetectorParams

		URL string `json:"url,omitempty" validate:"isdefault|url"`
		IP  string `json:"ip,omitempty" validate:"isdefault|ip"`
	}
)

const (
	catDNS  = "dns"
	catTCP  = "tcp"
	catPing = "ping"
	catHTTP = "http"
)

var (
	addDetectors         map[string]func(*elton.Context) (interface{}, error)
	updateDetectors      map[string]func(uint, *elton.Context) error
	getDetectors         map[string]func(uint) (interface{}, error)
	countDetectors       map[string]func(queryDetectorParams) (int, error)
	listDetectors        map[string]func(helper.PGQueryParams) (interface{}, error)
	countDetectorResults map[string]func(queryDetectorResultParams) (int, error)
	listDetectorResults  map[string]func(queryDetectorResultParams) (interface{}, error)
)

func init() {
	ctrl := detectorCtrl{}
	addDetectors = map[string]func(*elton.Context) (interface{}, error){
		catDNS:  ctrl.addDNS,
		catTCP:  ctrl.addTCP,
		catPing: ctrl.addPing,
		catHTTP: ctrl.addHTTP,
	}
	updateDetectors = map[string]func(uint, *elton.Context) error{
		catDNS:  ctrl.updateDNS,
		catTCP:  ctrl.updateTCP,
		catPing: ctrl.updatePing,
		catHTTP: ctrl.updateHTTP,
	}
	getDetectors = map[string]func(uint) (interface{}, error){
		catDNS:  ctrl.findDNS,
		catTCP:  ctrl.findTCP,
		catPing: ctrl.findPing,
		catHTTP: ctrl.findHTTP,
	}
	countDetectors = map[string]func(queryDetectorParams) (int, error){
		catDNS:  ctrl.countDNS,
		catTCP:  ctrl.countTCP,
		catPing: ctrl.countPing,
		catHTTP: ctrl.countHTTP,
	}
	listDetectors = map[string]func(helper.PGQueryParams) (interface{}, error){
		catDNS:  ctrl.listDNS,
		catTCP:  ctrl.listTCP,
		catPing: ctrl.listPing,
		catHTTP: ctrl.listHTTP,
	}

	countDetectorResults = map[string]func(queryDetectorResultParams) (int, error){
		catDNS:  ctrl.countDNSResult,
		catTCP:  ctrl.countTCPResult,
		catPing: ctrl.countPingResult,
		catHTTP: ctrl.countHTTPResult,
	}
	listDetectorResults = map[string]func(queryDetectorResultParams) (interface{}, error){
		catDNS:  ctrl.listDNSResult,
		catTCP:  ctrl.listTCPResult,
		catPing: ctrl.listPingResult,
		catHTTP: ctrl.listHTTPResult,
	}
	g := router.NewGroup("/detectors", loadUserSession, shouldLogined)

	g.POST(
		"/v1/{category}",
		newTracker(cs.ActionDetectorAdd),
		ctrl.checkCategory,
		ctrl.add,
	)
	g.GET(
		"/v1/{category}",
		ctrl.checkCategory,
		ctrl.list,
	)
	g.PATCH(
		"/v1/{category}/{id}",
		ctrl.checkCategory,
		newTracker(cs.ActionDetectorUpdate),
		ctrl.update,
	)
	g.GET(
		"/v1/{category}/{id}",
		ctrl.checkCategory,
		ctrl.getByID,
	)

	g.GET(
		"/v1/results/{category}",
		ctrl.checkCategory,
		ctrl.listResult,
	)
}

func (params *queryDetectorResultParams) toConditions() (conditions []interface{}) {
	queryList := make([]string, 0)
	args := make([]interface{}, 0)

	if params.Task != "" {
		queryList = append(queryList, "task = ?")
		args = append(args, params.Task)
	}

	if params.Result != "" {
		queryList = append(queryList, "result = ?")
		args = append(args, params.Result)
	}
	if params.Duration != "" {
		d, _ := time.ParseDuration(params.Duration)
		if d.Milliseconds() != 0 {
			queryList = append(queryList, "duration >= ?")
			args = append(args, d.Milliseconds())
		}
	}

	conditions = make([]interface{}, 0)
	if len(queryList) != 0 {
		conditions = append(conditions, strings.Join(queryList, " AND "))
		conditions = append(conditions, args...)
	}

	return
}

func (params *queryDetectorResultParams) toQueryParams() (queryParams helper.PGQueryParams) {
	queryParams.Limit, _ = strconv.Atoi(params.Limit)
	queryParams.Offset, _ = strconv.Atoi(params.Offset)
	queryParams.Order = params.Order
	return queryParams
}

// addDNS 添加dns detector
func (ctrl detectorCtrl) addDNS(c *elton.Context) (data interface{}, err error) {
	params := addDNSParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	d := &detector.DNS{
		Owner:       us.GetAccount(),
		Server:      params.Server,
		Hostname:    params.Hostname,
		Timeout:     params.Timeout,
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
	}
	err = dnsSrv.Add(d)
	if err != nil {
		return
	}
	data = d
	return
}

// updateDNS 更新dns detector
func (ctrl detectorCtrl) updateDNS(id uint, c *elton.Context) (err error) {
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

// findDNS find dns detector
func (ctrl detectorCtrl) findDNS(id uint) (data interface{}, err error) {
	d, err := dnsSrv.FindByID(id)
	if err != nil {
		return nil, err
	}
	data = d
	return
}

// countDNS count dns detector
func (ctrl detectorCtrl) countDNS(params queryDetectorParams) (count int, err error) {
	return dnsSrv.Count()
}

// listDNS list dns detector
func (ctrl detectorCtrl) listDNS(queryParams helper.PGQueryParams) (data interface{}, err error) {
	detectors, err := dnsSrv.List(queryParams)
	if err != nil {
		return
	}
	data = detectors
	return
}

// addTCP add tcp detector
func (ctrl detectorCtrl) addTCP(c *elton.Context) (data interface{}, err error) {
	params := addTCPParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	t := &detector.TCP{
		Owner:       us.GetAccount(),
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		Network:     params.Network,
		IP:          params.IP,
		Port:        params.Port,
	}
	err = tcpSrv.Add(t)
	if err != nil {
		return
	}
	data = t
	return
}

// updateTCP update tcp detector
func (ctrl detectorCtrl) updateTCP(id uint, c *elton.Context) (err error) {
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

// findTCP find tcp detector
func (ctrl detectorCtrl) findTCP(id uint) (data interface{}, err error) {
	t, err := tcpSrv.FindByID(id)
	if err != nil {
		return nil, err
	}
	data = t
	return
}

// countTCP count tcp detector
func (ctrl detectorCtrl) countTCP(params queryDetectorParams) (count int, err error) {
	return tcpSrv.Count()
}

// listTCP list tcp detector
func (ctrl detectorCtrl) listTCP(queryParams helper.PGQueryParams) (data interface{}, err error) {
	detectors, err := tcpSrv.List(queryParams)
	if err != nil {
		return
	}
	data = detectors
	return
}

// addPing add ping detector
func (ctrl detectorCtrl) addPing(c *elton.Context) (data interface{}, err error) {
	params := addPingParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	p := &detector.Ping{
		Owner:       us.GetAccount(),
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		Network:     params.Network,
		IP:          params.IP,
	}
	err = pingSrv.Add(p)
	if err != nil {
		return
	}
	data = p
	return
}

// updatePing update ping detector
func (ctrl detectorCtrl) updatePing(id uint, c *elton.Context) (err error) {
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

// findPing find ping detector
func (ctrl detectorCtrl) findPing(id uint) (data interface{}, err error) {
	p, err := pingSrv.FindByID(id)
	if err != nil {
		return nil, err
	}
	data = p
	return
}

// countPing count ping detector
func (ctrl detectorCtrl) countPing(params queryDetectorParams) (count int, err error) {
	return pingSrv.Count()
}

// listPing list ping detector
func (ctrl detectorCtrl) listPing(queryParams helper.PGQueryParams) (data interface{}, err error) {
	detectors, err := pingSrv.List(queryParams)
	if err != nil {
		return
	}
	data = detectors
	return
}

// addHTTP add http detector
func (ctrl detectorCtrl) addHTTP(c *elton.Context) (data interface{}, err error) {
	params := addHTTPParams{}
	err = validate.Do(&params, c.RequestBody)
	if err != nil {
		return
	}
	us := getUserSession(c)
	h := &detector.HTTP{
		Owner:       us.GetAccount(),
		Status:      params.Status,
		Description: params.Description,
		Receivers:   params.Receivers,
		Timeout:     params.Timeout,
		IP:          params.IP,
		URL:         params.URL,
	}
	err = httpSrv.Add(h)
	if err != nil {
		return
	}
	data = h
	return
}

// updateHTTP update http detector
func (ctrl detectorCtrl) updateHTTP(id uint, c *elton.Context) (err error) {
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

// findHTTP find http detector
func (ctrl detectorCtrl) findHTTP(id uint) (data interface{}, err error) {
	h, err := httpSrv.FindByID(id)
	if err != nil {
		return nil, err
	}
	data = h
	return
}

// countHTTP count http detector
func (ctrl detectorCtrl) countHTTP(params queryDetectorParams) (count int, err error) {
	return httpSrv.Count()
}

// listHTTP list http detector
func (ctrl detectorCtrl) listHTTP(queryParams helper.PGQueryParams) (data interface{}, err error) {
	detectors, err := httpSrv.List(queryParams)
	if err != nil {
		return
	}
	data = detectors
	return
}

func (ctrl detectorCtrl) checkCategory(c *elton.Context) (err error) {
	cat := c.Param("category")
	switch cat {
	case catDNS:
	case catTCP:
	case catPing:
	case catHTTP:
	default:
		err = hes.New(fmt.Sprintf("Not support category:%s", cat))
	}
	if err != nil {
		return
	}
	return c.Next()
}

func (ctrl detectorCtrl) add(c *elton.Context) (err error) {
	fn := addDetectors[c.Param("category")]
	data, err := fn(c)

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
	fn := updateDetectors[cat]
	err = fn(id, c)
	if err != nil {
		return
	}
	c.NoContent()
	return
}

func (ctrl detectorCtrl) list(c *elton.Context) (err error) {
	cat := c.Param("category")

	params := queryDetectorParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}

	queryParams := helper.PGQueryParams{}
	queryParams.Limit, _ = strconv.Atoi(params.Limit)
	queryParams.Offset, _ = strconv.Atoi(params.Offset)
	queryParams.Order = params.Order

	count := -1
	if queryParams.Offset == 0 {
		count, err = countDetectors[cat](params)
		if err != nil {
			return
		}
	}

	data, err := listDetectors[cat](queryParams)
	if err != nil {
		return
	}
	c.Body = map[string]interface{}{
		"count":     count,
		"detectors": data,
	}
	return
}

func (ctrl detectorCtrl) getByID(c *elton.Context) (err error) {
	cat := c.Param("category")
	v, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	fn := getDetectors[cat]
	data, err := fn(uint(v))
	if err != nil {
		return
	}
	c.Body = data
	return
}

func (ctrl detectorCtrl) listHTTPResult(params queryDetectorResultParams) (data interface{}, err error) {
	results, err := httpSrv.ListResult(params.toQueryParams(), params.toConditions()...)
	if err != nil {
		return
	}
	data = results
	return
}
func (ctrl detectorCtrl) countHTTPResult(params queryDetectorResultParams) (count int, err error) {
	return httpSrv.CountResult(params.toConditions()...)
}

func (ctrl detectorCtrl) listDNSResult(params queryDetectorResultParams) (data interface{}, err error) {
	results, err := dnsSrv.ListResult(params.toQueryParams(), params.toConditions()...)
	if err != nil {
		return
	}
	data = results
	return
}
func (ctrl detectorCtrl) countDNSResult(params queryDetectorResultParams) (count int, err error) {
	return dnsSrv.CountResult(params.toConditions()...)
}

func (ctrl detectorCtrl) listPingResult(params queryDetectorResultParams) (data interface{}, err error) {
	results, err := pingSrv.ListResult(params.toQueryParams(), params.toConditions()...)
	if err != nil {
		return
	}
	data = results
	return
}
func (ctrl detectorCtrl) countPingResult(params queryDetectorResultParams) (count int, err error) {
	return pingSrv.CountResult(params.toConditions()...)
}

func (ctrl detectorCtrl) listTCPResult(params queryDetectorResultParams) (data interface{}, err error) {
	results, err := tcpSrv.ListResult(params.toQueryParams(), params.toConditions()...)
	if err != nil {
		return
	}
	data = results
	return
}
func (ctrl detectorCtrl) countTCPResult(params queryDetectorResultParams) (count int, err error) {
	return tcpSrv.CountResult(params.toConditions()...)
}

func (ctrl detectorCtrl) listResult(c *elton.Context) (err error) {
	cat := c.Param("category")

	params := queryDetectorResultParams{}
	err = validate.Do(&params, c.Query())
	if err != nil {
		return
	}

	offset, _ := strconv.Atoi(params.Offset)

	count := -1
	if offset == 0 {
		count, err = countDetectorResults[cat](params)
		if err != nil {
			return
		}
	}

	fn := listDetectorResults[cat]
	data, err := fn(params)
	if err != nil {
		return
	}
	c.Body = map[string]interface{}{
		"count":   count,
		"results": data,
	}
	return
}
