package controller

import (
	"os"
	"runtime"
	"time"

	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/global"
	"github.com/vicanso/cyber-tect/router"
)

type (
	systemCtrl struct{}
	// StatusResp status response
	StatusResp struct {
		Status     string `json:"status,omitempty"`
		Uptime     string `json:"uptime,omitempty"`
		StartedAt  string `json:"startedAt,omitempty"`
		GoMaxProcs int    `json:"goMaxProcs,omitempty"`
		Version    string `json:"version,omitempty"`
		Pid        int    `json:"pid,omitempty"`
	}
	// StatsResp stats response
	StatsResp struct {
		Sys             uint64 `json:"sys,omitempty"`
		HeapSys         uint64 `json:"heapSys,omitempty"`
		HeapInuse       uint64 `json:"heapInuse,omitempty"`
		RoutineCount    int    `json:"routineCount,omitempty"`
		ConnectingCount uint32 `json:"connectingCount,omitempty"`
	}
)

var (
	systemStartedAt = time.Now()
)

func init() {
	g := router.NewGroup("/system")
	ctrl := systemCtrl{}
	g.GET("/status", noQuery, ctrl.getStatus)
	g.GET("/stats", noQuery, ctrl.getStats)
	g.GET("/routes", noQuery, ctrl.getRoutes)
	g.GET("/route-counts", noQuery, ctrl.getRouteCounts)
}

// getStatus get application status
func (ctrl systemCtrl) getStatus(c *cod.Context) (err error) {
	status := "running"
	if !global.IsApplicationRunning() {
		status = "pause"
	}
	c.CacheMaxAge("10s")
	c.Body = &StatusResp{
		Status:     status,
		Uptime:     time.Since(systemStartedAt).String(),
		StartedAt:  systemStartedAt.Format(time.RFC3339),
		GoMaxProcs: runtime.GOMAXPROCS(0),
		Version:    runtime.Version(),
		Pid:        os.Getpid(),
	}
	return
}

// getStats get stats info
func (ctrl systemCtrl) getStats(c *cod.Context) (err error) {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)
	var mb uint64 = 1024 * 1024
	c.CacheMaxAge("10s")
	c.Body = &StatsResp{
		Sys:             mem.Sys / mb,
		HeapSys:         mem.HeapSys / mb,
		HeapInuse:       mem.HeapInuse / mb,
		RoutineCount:    runtime.NumGoroutine(),
		ConnectingCount: global.GetConnectingCount(),
	}
	return
}

// getRoutes get the route infos
func (ctrl systemCtrl) getRoutes(c *cod.Context) (err error) {
	c.CacheMaxAge("1m")
	c.Body = &struct {
		Routes []*cod.RouterInfo `json:"routes,omitempty"`
	}{
		c.Cod(nil).Routers,
	}
	return
}

// getRouteCounts get route counts
func (ctrl systemCtrl) getRouteCounts(c *cod.Context) (err error) {
	routeCountInfo := router.GetRouteCount()
	c.CacheMaxAge("1m")
	c.Body = routeCountInfo
	return
}
