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

package helper

import (
	"context"
	"database/sql"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/iancoleman/strcase"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/vicanso/cybertect/config"
	"github.com/vicanso/cybertect/cs"
	"github.com/vicanso/cybertect/ent"
	"github.com/vicanso/cybertect/ent/hook"
	"github.com/vicanso/cybertect/ent/migrate"
	"github.com/vicanso/cybertect/log"
	"github.com/vicanso/cybertect/util"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

var (
	defaultEntDriver, defaultEntClient = mustNewEntClient()
)
var (
	initSchemaOnce sync.Once

	maskRegExp = regexp.MustCompile(`(?i)password`)
)

// processingKeyAll 记录所有表的正在处理请求
const processingKeyAll = "All"

// entProcessingStats ent的处理请求统计
type entProcessingStats struct {
	data map[string]*atomic.Int32
}

// EntEntListParams 公共的列表查询参数
type EntListParams struct {
	Limit  string `json:"limit,omitempty" validate:"required,xLimit"`
	Offset string `json:"offset,omitempty" validate:"omitempty,xOffset"`
	Fields string `json:"fields,omitempty" validate:"omitempty,xFields"`
	Order  string `json:"order,omitempty" validate:"omitempty,xOrder"`
	// IgnoreCount 忽略计算总数
	IgnoreCount string `json:"ignoreCount,omitempty"`
}

var currentEntProcessingStats = new(entProcessingStats)

// mustNewEntClient 初始化客户端与driver
func mustNewEntClient() (*entsql.Driver, *ent.Client) {
	postgresConfig := config.GetPostgresConfig()

	maskURI := postgresConfig.URI
	urlInfo, _ := url.Parse(maskURI)
	if urlInfo != nil {
		pass, ok := urlInfo.User.Password()
		if ok {
			maskURI = strings.ReplaceAll(maskURI, pass, "***")
		}
	}
	log.Default().Info("connect postgres",
		zap.String("uri", maskURI),
	)
	db, err := sql.Open("pgx", postgresConfig.URI)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	driver := entsql.OpenDB(dialect.Postgres, db)
	c := ent.NewClient(ent.Driver(driver))

	ctx := context.Background()
	if err := c.Schema.Create(ctx); err != nil {
		panic(err)
	}
	initSchemaHooks(c)
	return driver, c
}

// GetLimit 获取limit的值
func (params *EntListParams) GetLimit() int {
	limit, _ := strconv.Atoi(params.Limit)
	// 保证limit必须大于0
	if limit <= 0 {
		limit = 10
	}
	return limit
}

// GetOffset 获取offset的值
func (params *EntListParams) GetOffset() int {
	offset, _ := strconv.Atoi(params.Offset)
	return offset
}

// GetOrders 获取排序的函数列表
func (params *EntListParams) GetOrders() []ent.OrderFunc {
	if params.Order == "" {
		return nil
	}
	arr := strings.Split(params.Order, ",")
	funcs := make([]ent.OrderFunc, len(arr))
	for index, item := range arr {
		if item[0] == '-' {
			funcs[index] = ent.Desc(strcase.ToSnake(item[1:]))
		} else {
			funcs[index] = ent.Asc(strcase.ToSnake(item))
		}
	}
	return funcs
}

// GetFields 获取选择的字段
func (params *EntListParams) GetFields() []string {
	if params.Fields == "" {
		return nil
	}
	arr := strings.Split(params.Fields, ",")
	result := make([]string, len(arr))
	for index, item := range arr {
		result[index] = strcase.ToSnake(item)
	}
	return result
}

// ShouldCount 判断是否需要计算总数
func (params *EntListParams) ShouldCount() bool {
	return params.IgnoreCount == "" && params.GetOffset() == 0
}

// init 初始化统计
func (stats *entProcessingStats) init(schemas []string) {
	data := make(map[string]*atomic.Int32)
	data[processingKeyAll] = atomic.NewInt32(0)
	for _, schema := range schemas {
		data[schema] = atomic.NewInt32(0)
	}
	stats.data = data
}

// inc 处理数+1
func (stats *entProcessingStats) inc(schema string) (int32, int32) {
	total := stats.data[processingKeyAll].Inc()
	p, ok := stats.data[schema]
	if !ok {
		return total, 0
	}
	return total, p.Inc()
}

// desc 处理数-1
func (stats *entProcessingStats) dec(schema string) (int32, int32) {
	total := stats.data[processingKeyAll].Dec()
	p, ok := stats.data[schema]
	if !ok {
		return total, 0
	}
	return total, p.Dec()
}

// initSchemaHooks 初始化相关的hooks
func initSchemaHooks(c *ent.Client) {
	schemas := make([]string, len(migrate.Tables))
	for index, table := range migrate.Tables {
		name := strcase.ToCamel(table.Name)
		// 去除最后的复数s
		schemas[index] = name[:len(name)-1]
	}
	currentEntProcessingStats.init(schemas)
	ignoredNameList := []string{
		"updated_at",
		"created_at",
	}
	isIgnored := func(name string) bool {
		for _, item := range ignoredNameList {
			if item == name {
				return true
			}
		}
		return false
	}
	// 禁止删除数据
	c.Use(hook.Reject(ent.OpDelete | ent.OpDeleteOne))
	// 数据库操作统计
	c.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			schemaType := m.Type()
			totalProcessing, processing := currentEntProcessingStats.inc(schemaType)
			defer currentEntProcessingStats.dec(schemaType)
			op := m.Op().String()

			startedAt := time.Now()
			result := cs.ResultSuccess
			message := ""
			mutateResult, err := next.Mutate(ctx, m)
			// 如果失败，则记录出错信息
			if err != nil {
				result = cs.ResultFail
				message = err.Error()
			}
			data := make(map[string]interface{})
			for _, name := range m.Fields() {
				if isIgnored(name) {
					continue
				}
				value, ok := m.Field(name)
				if !ok {
					continue
				}
				valueType := reflect.TypeOf(value)
				maxString := 50
				switch valueType.Kind() {
				case reflect.String:
					str, ok := value.(string)
					// 如果更新过长，则截断
					if ok {
						value = util.CutRune(str, maxString)
					}
				}

				if maskRegExp.MatchString(name) {
					data[name] = "***"
				} else {
					data[name] = value
				}
			}

			d := time.Since(startedAt)
			log.Default().Info("ent stats",
				zap.String("schema", schemaType),
				zap.String("op", op),
				zap.Int("result", result),
				zap.Int32("processing", processing),
				zap.Int32("totalProcessing", totalProcessing),
				zap.String("use", d.String()),
				zap.Any("data", data),
				zap.String("message", message),
			)
			fields := map[string]interface{}{
				cs.FieldProcessing:      int(processing),
				cs.FieldTotalProcessing: int(totalProcessing),
				cs.FieldUse:             int(d.Milliseconds()),
				cs.FieldData:            data,
			}
			if message != "" {
				fields[cs.FieldError] = message
			}
			tags := map[string]string{
				cs.TagSchema: schemaType,
				cs.TagOP:     op,
				cs.TagResult: strconv.Itoa(result),
			}
			GetInfluxSrv().Write(cs.MeasurementEntOP, tags, fields)
			return mutateResult, err
		})
	})
}

// EntGetStats get ent stats
func EntGetStats() map[string]interface{} {
	info := defaultEntDriver.DB().Stats()
	stats := map[string]interface{}{
		cs.FieldMaxOpenConns:      info.MaxOpenConnections,
		cs.FieldOpenConns:         info.OpenConnections,
		cs.FieldInUseConns:        info.InUse,
		cs.FieldIdleConns:         info.Idle,
		cs.FieldWaitCount:         int(info.WaitCount),
		cs.FieldWaitDuration:      int(info.WaitDuration.Milliseconds()),
		cs.FieldMaxIdleClosed:     int(info.MaxIdleClosed),
		cs.FieldMaxIdleTimeClosed: int(info.MaxIdleTimeClosed),
		cs.FieldMaxLifetimeClosed: int(info.MaxLifetimeClosed),
	}
	for name, p := range currentEntProcessingStats.data {
		stats[strcase.ToLowerCamel(name)] = p.Load()
	}
	return stats
}

// EntGetClient get ent client
func EntGetClient() *ent.Client {
	return defaultEntClient
}

// EntPing ent driver ping
func EntPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return defaultEntDriver.DB().PingContext(ctx)
}

// EntInitSchema 初始化schema
func EntInitSchema() (err error) {
	// 只执行一次schema初始化以及hook
	initSchemaOnce.Do(func() {
		err = defaultEntClient.Schema.Create(context.Background())
	})
	return
}
