package router

import (
	"sync"
	"sync/atomic"

	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/config"
	"github.com/vicanso/cyber-tect/util"
)

const (
	apiPrefixKey = "apiPrefix"
)

var (
	// groupList 路由组列表
	groupList = make([]*cod.Group, 0)

	// routeCounter the route counter info
	routeCounter     = &RouteCounter{}
	routeCounterLock = new(sync.RWMutex)
)

type (
	// RouteCounter route counter
	RouteCounter struct {
		CreatedAt string
		Counts    map[string]*uint32
	}
)

// NewGroup new router group
func NewGroup(path string, handlerList ...cod.Handler) *cod.Group {
	// 如果配置文件中有配置路由
	path = config.GetString(apiPrefixKey) + path
	g := cod.NewGroup(path, handlerList...)
	groupList = append(groupList, g)
	return g
}

// GetGroups get groups
func GetGroups() []*cod.Group {
	return groupList
}

// InitRouteCounter init route counter
func InitRouteCounter(routeInfos []*cod.RouterInfo) {
	routeCounterLock.Lock()
	defer routeCounterLock.Unlock()
	routeCounter.CreatedAt = util.NowString()
	routeCounter.Counts = make(map[string]*uint32)
	counts := routeCounter.Counts
	for _, info := range routeInfos {
		key := info.Method + " " + info.Path
		var v uint32
		counts[key] = &v
	}
}

// AddRouteCount add the route's count
func AddRouteCount(method, path string) {
	if method == "" || path == "" {
		return
	}
	key := method + " " + path
	v := routeCounter.Counts[key]
	if v == nil {
		return
	}
	atomic.AddUint32(v, 1)
}

// ResetRouteCount reset the route count
func ResetRouteCount() {
	for _, v := range routeCounter.Counts {
		atomic.StoreUint32(v, 0)
	}
	routeCounterLock.Lock()
	defer routeCounterLock.Unlock()
	routeCounter.CreatedAt = util.NowString()
}

// GetRouteCount get the route count
func GetRouteCount() map[string]interface{} {
	routeCounterLock.RLock()
	defer routeCounterLock.RUnlock()
	m := make(map[string]uint32)
	for k, v := range routeCounter.Counts {
		m[k] = *v
	}
	data := make(map[string]interface{})
	data["createdAt"] = routeCounter.CreatedAt
	data["counts"] = m
	return data
}
