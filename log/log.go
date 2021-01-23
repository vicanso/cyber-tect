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

// 可通过zap.RegisterSink添加更多的sink实现不同方式的日志传输

package log

import (
	"context"
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/vicanso/cybertect/util"
)

var defaultLogger = mustNewLogger("")

type httpServerLogger struct{}

func (hsl *httpServerLogger) Write(p []byte) (int, error) {
	defaultLogger.Info(string(p),
		zap.String("category", "httpServerLogger"),
	)
	return len(p), nil
}

type redisLogger struct{}

func (rl *redisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.Info(fmt.Sprintf(format, v...),
		zap.String("category", "redisLogger"),
	)
}

// mustNewLogger 初始化logger
func mustNewLogger(outputPath string) *zap.Logger {

	if util.IsDevelopment() {
		c := zap.NewDevelopmentConfig()
		l, err := c.Build(zap.AddStacktrace(zap.ErrorLevel))
		if err != nil {
			panic(err)
		}
		return l
	}
	c := zap.NewProductionConfig()
	if outputPath != "" {
		c.OutputPaths = []string{
			outputPath,
		}
		c.ErrorOutputPaths = []string{
			outputPath,
		}
	}

	// 在一秒钟内, 如果某个级别的日志输出量超过了 Initial, 那么在超过之后, 每 Thereafter 条日志才会输出一条, 其余的日志都将被删除
	c.Sampling.Initial = 1000
	// 如果不希望任何日志丢失，则设置为nil
	// c.Sampling = nil

	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 只针对panic 以上的日志增加stack trace
	l, err := c.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		panic(err)
	}
	return l
}

// SetOutputPath 设置日志的输出目录，在程序初始化时使用
func SetOutputPath(outputPath string) {
	defaultLogger = mustNewLogger(outputPath)
}

// Default 获取默认的logger
func Default() *zap.Logger {
	return defaultLogger
}

// NewHTTPServerLogger create a http server logger
func NewHTTPServerLogger() *log.Logger {
	return log.New(&httpServerLogger{}, "", 0)
}

// NewRedisLogger create a redis logger
func NewRedisLogger() *redisLogger {
	return &redisLogger{}
}
