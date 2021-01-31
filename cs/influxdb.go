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

package cs

const (
	// MeasurementPerformance 应用性能统计
	MeasurementPerformance = "performance"
	// MeasurementHTTPRequest http request统计
	MeasurementHTTPRequest = "httpRequest"
	// MeasurementRedisStats redis性能统计
	MeasurementRedisStats = "redisStats"
	// MeasurementRedisError redis出错统计
	MeasurementRedisError = "redisError"
	// MeasurementRedisConn redis连接
	MeasurementRedisConn = "redisConn"
	// MeasurementRouterConcurrency 路由并发访问
	MeasurementRouterConcurrency = "routerConcurrency"
	// MeasurementHTTPStats http性能统计
	MeasurementHTTPStats = "httpStats"
	// MeasurementHTTPInstanceStats http instance统计
	MeasurementHTTPInstanceStats = "httpInstanceStats"
	// MeasurementEntStats ent性能统计
	MeasurementEntStats = "entStats"
	// MeasurementEntOP ent的操作记录
	MeasurementEntOP = "entOP"
	// MeasurementHTTPError http响应出错统计
	MeasurementHTTPError = "httpError"
	// MeasurementUserTracker 用户行为记录
	MeasurementUserTracker = "userTracker"
	// MeasurementUserAction 用户行为记录
	// 用于前端记录客户相关的操作，如点击、确认、取消等
	MeasurementUserAction = "userAction"
	// MeasurementUserLogin 用户登录
	MeasurementUserLogin = "userLogin"
	// MeasurementUserAddTrack 添加用户跟踪
	MeasurementUserAddTrack = "userAddTrack"
	// MeasurementException 异常
	MeasurementException = "exception"
)

const (
	// TagCategory 分类
	TagCategory = "category"
	// TagRoute 路由
	TagRoute = "route"
	// TagService 服务名称
	TagService = "service"
	// TagAction 用户的操作action
	TagAction = "action"
	// TagResult 操作结果
	TagResult = "result"
	// TagSchema 数据库的schema
	TagSchema = "schema"
	// TagOP 数据库的操作
	TagOP = "op"
	// TagMethod http method
	TagMethod = "method"
)

// string 类型
const (
	// FieldIP ip
	FieldIP = "ip"
	// FieldURI uri
	FieldURI = "uri"
	// FieldRoute route
	FieldRoute = "route"
	// FieldPath path
	FieldPath = "path"
	// FieldAccount 账号
	FieldAccount = "account"
	// FieldSID session id
	FieldSID = "sid"
	// FieldTID track id
	FieldTID = "tid"
	// FieldQuery url query
	FieldQuery = "query"
	// FieldParams url route params
	FieldParams = "params"
	// FieldForm request body
	FieldForm = "form"
	// FieldError error message
	FieldError = "error"
	// FieldUserAgent user agent
	FieldUserAgent = "userAgent"
	// FieldCountry 国家
	FieldCountry = "country"
	// FieldProvince 省份
	FieldProvince = "province"
	// FieldCity 城市
	FieldCity = "city"
	// FieldISP ISP
	FieldISP = "isp"
	// FieldCategory 分类（注意tag也有分类字段，按需使用）
	FieldCategory = "category"
)

// int 类型
const (
	// FieldProcessing 正在处理请求数
	FieldProcessing = "processing"
	// FieldTotalProcessing 正在处理的总请求数
	FieldTotalProcessing = "totalProcessing"
	// FilePipeProcessing pipe的正在处理请求数
	FilePipeProcessing = "pipeProcessing"
	// FieldUse 耗时
	FieldUse = "use"
	// FieldStatus 状态码
	FieldStatus = "status"
	// FieldDNSUse dns耗时
	FieldDNSUse = "dnsUse"
	// FieldTCPUse tcp耗时
	FieldTCPUse = "tcpUse"
	// FieldTLSUse tls耗时
	FieldTLSUse = "tlsUse"
	// FieldProcessingUse 服务器处理耗时
	FieldProcessingUse = "processingUse"
	// FieldTransferUse 数据传输耗时
	FieldTransferUse = "transferUse"
	// FieldCount 总数
	FieldCount = "count"
	// FieldSize 大小
	FieldSize = "size"
	// FieldHits 命中数量
	FieldHits = "hits"
	// FieldMisses miss数量
	FieldMisses = "misses"
	// FieldTimeouts 超时数量
	FieldTimeouts = "timeouts"
	// FieldTotalConns 总连接
	FieldTotalConns = "totalConns"
	// FieldIdleConns idle连接数
	FieldIdleConns = "idleConns"
	// FieldStaleConns stale连接数
	FieldStaleConns = "staleConns"
	// FieldMaxOpenConns 最大的连接数
	FieldMaxOpenConns = "maxOpenConns"
	// FieldOpenConns 当前连接数
	FieldOpenConns = "openConns"
	// FieldInUseConns 正在使用的连接
	FieldInUseConns = "inUseConns"
	// FieldWaitCount 等待的总数
	FieldWaitCount = "waitCount"
	// FieldWaitDuration 等待的时间
	FieldWaitDuration = "waitDuration"
	// FieldMaxIdleClosed idle close count
	FieldMaxIdleClosed = "maxIdleClosed"
	// FieldMaxIdleTimeClosed idle time close
	FieldMaxIdleTimeClosed = "maxIdleTimeClosed"
	// FieldMaxLifetimeClosed life time close
	FieldMaxLifetimeClosed = "maxLifetimeClosed"
	// FieldGoMaxProcs go max procs
	FieldGoMaxProcs = "goMaxProcs"
	// FieldThreadCount thread count
	FieldThreadCount = "threadCount"
	// FieldMemSys mem sys
	FieldMemSys = "memSys"
	// FieldMemHeapSys mem heap sys
	FieldMemHeapSys = "memHeapSys"
	// FieldMemHeapInuse mem heap inuse
	FieldMemHeapInuse = "memHeapInuse"
	// FieldMemFrees mem frees
	FieldMemFrees = "memFrees"
	// FieldRoutineCount routine count
	FieldRoutineCount = "routineCount"
	// FieldCpuUsage cpu usage
	FieldCpuUsage = "cpuUsage"
	// FieldNumGC num gc
	FieldNumGC = "numGC"
	// FieldPauseNS gc pause ns
	FieldPauseNS = "pauseNS"
	// FieldConnProcessing conn processing
	FieldConnProcessing = "connProcessing"
	// FieldConnAlive conn alive
	FieldConnAlive = "connAlive"
	// FieldConnCreatedCount conn created count
	FieldConnCreatedCount = "connCreatedCount"
	// FieldTotal 总数
	FieldTotal = "total"
)

// bool 类型
const (
	// FieldReused 是否复用
	FieldReused = "reused"
	// FieldException 是否异常
	FieldException = "exception"
)

// map[string]interface{} 类型
const (
	// FieldData 数据
	FieldData = "data"
)
