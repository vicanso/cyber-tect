package controller

import (
	"github.com/vicanso/cyber-tect/log"
	"github.com/vicanso/cyber-tect/middleware"
	"github.com/vicanso/cyber-tect/util"
)

var (
	logger     = log.Default()
	now        = util.NowString
	getTrackID = util.GetTrackID

	noQuery = middleware.NoQuery
	waitFor = middleware.WaitFor
)
