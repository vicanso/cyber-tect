package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/config"
	_ "github.com/vicanso/cyber-tect/controller"
	"github.com/vicanso/cyber-tect/global"
	"github.com/vicanso/cyber-tect/log"
	"github.com/vicanso/cyber-tect/middleware"
	"github.com/vicanso/cyber-tect/router"
	_ "github.com/vicanso/cyber-tect/schedule"
	"github.com/vicanso/cyber-tect/util"
	"github.com/vicanso/hes"

	"go.uber.org/zap"

	humanize "github.com/dustin/go-humanize"
	bodyparser "github.com/vicanso/cod-body-parser"
	compress "github.com/vicanso/cod-compress"
	errorHandler "github.com/vicanso/cod-error-handler"
	etag "github.com/vicanso/cod-etag"
	fresh "github.com/vicanso/cod-fresh"
	recover "github.com/vicanso/cod-recover"
	responder "github.com/vicanso/cod-responder"
	stats "github.com/vicanso/cod-stats"
)

func main() {
	logger := log.Default()
	listen := config.GetListen()

	d := cod.New()

	d.Keys = config.GetStringSlice("keys")

	// 如果出错则会触发此回调（在 ErrorHandler 中会将出错转换为相应的http响应，此类情况不会触发）
	d.OnError(func(c *cod.Context, err error) {
		// 可以针对实际场景输出更多的日志信息
		logger.DPanic("exception",
			zap.String("uri", c.Request.RequestURI),
			zap.Error(err),
		)
	})

	d.Use(recover.New())

	d.Use(middleware.NewEntry())

	// 接口响应统计，在项目中可写入数据库方便统计
	d.Use(stats.New(stats.Config{
		OnStats: func(statsInfo *stats.Info, c *cod.Context) {
			// 增加从session中获取当前账号
			logger.Info("access log",
				zap.String("cid", statsInfo.CID),
				zap.String("track", util.GetTrackID(c)),
				zap.String("ip", statsInfo.IP),
				zap.String("method", statsInfo.Method),
				zap.String("uri", statsInfo.URI),
				zap.Int("status", statsInfo.Status),
				zap.String("consuming", statsInfo.Consuming.String()),
				zap.String("size", humanize.Bytes(uint64(statsInfo.Size))),
			)
		},
	}))

	d.Use(errorHandler.NewDefault())

	d.Use(middleware.NewLimiter())

	d.Use(bodyparser.NewDefault())

	d.Use(fresh.NewDefault())
	d.Use(etag.NewDefault())
	d.Use(compress.NewDefault())

	d.Use(responder.NewDefault())

	// health check
	d.GET("/ping", func(c *cod.Context) (err error) {
		if !global.IsApplicationRunning() {
			err = hes.NewWithStatusCode("application is not running", http.StatusServiceUnavailable)
			return
		}
		c.Body = "pong"
		return
	})

	groups := router.GetGroups()
	for _, g := range groups {
		d.AddGroup(g)
	}

	router.InitRouteCounter(d.Routers)

	logger.Info("server is starting",
		zap.String("listen", listen),
	)

	// 设置应用状态为运行中
	global.StartApplication()
	if util.IsDevelopment() {
		err := d.ListenAndServe(listen)
		if err != nil {
			panic(err)
		}
		return
	}

	errCh := make(chan error)
	go func() {
		err := d.ListenAndServe(listen)
		if err != nil {
			errCh <- err
		}
	}()

	closeCh := make(chan os.Signal)
	signal.Notify(closeCh, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		logger.Error("server is closed by error",
			zap.Error(err),
		)
	case sign := <-closeCh:
		logger.Info("server will be closed by signal")
		d.GracefulClose(10 * time.Second)
		logger.Info("server is closed by signal",
			zap.String("sign", sign.String()),
		)
	}
}
