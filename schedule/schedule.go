package schedule

import (
	"time"

	"github.com/vicanso/cyber-tect/log"
	"github.com/vicanso/cyber-tect/router"
	"github.com/vicanso/cyber-tect/util"

	"go.uber.org/zap"
)

func init() {
	if util.IsDevelopment() {
		return
	}
	go initRouteCountTicker()
	// go initInfluxdbCheckTicker()
	// go initRouterConfigRefreshTicker()
}

func runTicker(ticker *time.Ticker, message string, do func() error, restart func()) {
	defer func() {
		if r := recover(); r != nil {
			err, _ := r.(error)
			log.Default().DPanic(message+" panic",
				zap.Error(err),
			)
		}
		// 如果退出了，重新启动
		go restart()
	}()
	for range ticker.C {
		err := do()
		// TODO 检测不通过时，发送告警
		if err != nil {
			log.Default().Error(message+" fail",
				zap.Error(err),
			)
		}
	}
}

func initRouteCountTicker() {
	// 每5分钟重置route count
	ticker := time.NewTicker(5 * time.Minute)
	runTicker(ticker, "reset route count", func() error {
		router.ResetRouteCount()
		return nil
	}, initRouteCountTicker)
}
