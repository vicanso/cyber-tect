package controller

import (
	"strconv"
	"time"

	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/router"
	"github.com/vicanso/cyber-tect/service"
	"github.com/vicanso/cyber-tect/util"
	"github.com/vicanso/hes"

	"github.com/oklog/ulid"
)

type (
	commonCtrl struct{}
)

func init() {
	ctrl := commonCtrl{}
	g := router.NewGroup("/common")
	g.GET("/ulid2time/:id", noQuery, ctrl.ulid2time)
	g.GET("/objid2time/:id", noQuery, ctrl.objectID2time)
	// 请求限制最少1秒才响应，避免接口被刷
	g.GET("/ip-location/:ip", noQuery, waitFor(time.Second), ctrl.getLocationByIP)
}

// ulid2time convert ulid to time
func (ctrl commonCtrl) ulid2time(c *cod.Context) (err error) {
	id, e := ulid.ParseStrict(c.Param("id"))
	if e != nil {
		err = hes.NewWithError(e)
		return
	}
	t := ulid.Time(id.Time())
	c.Body = &struct {
		Time string `json:"time"`
	}{
		util.FormatTime(t),
	}
	return
}

// objectID2time convert object id to time
func (ctrl commonCtrl) objectID2time(c *cod.Context) (err error) {
	id := c.Param("id")
	// object id 长度为24
	if len(id) != 24 {
		err = hes.New("object id is invalid")
	}
	seconds, e := strconv.ParseInt(id[0:8], 16, 64)
	if e != nil {
		err = hes.NewWithError(e)
		return
	}
	t := time.Unix(seconds, 0)
	c.Body = &struct {
		Time string `json:"time"`
	}{
		util.FormatTime(t),
	}
	return
}

// getLocationByIP get location by ip address
func (ctrl commonCtrl) getLocationByIP(c *cod.Context) (err error) {
	info, err := service.GetLocationByIP(c.Param("ip"), c)

	if err != nil {
		return
	}
	c.Body = info
	return
}
