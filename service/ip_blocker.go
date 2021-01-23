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

package service

import (
	"sync"

	"github.com/vicanso/ips"
)

var (
	ipBlocker = &IPBlocker{
		mutex: sync.RWMutex{},
		IPS:   ips.New(),
	}
)

type (
	// IPBlocker IP拦截器

	IPBlocker struct {
		mutex sync.RWMutex
		IPS   *ips.IPS
	}
)

// ResetIPBlocker 重置IP拦截器的IP列表
func ResetIPBlocker(ipList []string) {
	// blocker 有读写锁，因此ips可以使用无锁
	list := ips.NewWithoutMutex()
	for _, value := range ipList {
		_ = list.Add(value)
	}
	ipBlocker.mutex.Lock()
	defer ipBlocker.mutex.Unlock()
	ipBlocker.IPS = list
}

// IsBlockIP 判断该IP是否有需要拦截
func IsBlockIP(ip string) bool {
	ipBlocker.mutex.RLock()
	defer ipBlocker.mutex.RUnlock()
	blocked := ipBlocker.IPS.Contains(ip)
	return blocked
}

// GetIPBlockList 获取block的ip地址列表
func GetIPBlockList() []string {
	ipBlocker.mutex.RLock()
	defer ipBlocker.mutex.RUnlock()
	ipList := make([]string, 0)
	for _, item := range ipBlocker.IPS.IPList {
		ipList = append(ipList, item.String())
	}
	for _, item := range ipBlocker.IPS.IPNetList {
		ipList = append(ipList, item.String())
	}
	return ipList
}
