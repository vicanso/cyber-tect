package controller

import (
	"github.com/vicanso/cod"
	"github.com/vicanso/cyber-tect/router"
	"github.com/vicanso/cyber-tect/umbra"
	"github.com/vicanso/cyber-tect/validate"
)

type (
	cyberCtrl struct{}

	httpStatsParams struct {
		URL string `json:"url" valid:"url"`
		IP  string `json:"ip" valid:"ip,optional"`
	}
)

func init() {
	g := router.NewGroup("/cyber")
	ctrl := cyberCtrl{}
	g.GET("/http", ctrl.getHTTPStats)
}

// getHTTPStats get http stats
func (ctrl cyberCtrl) getHTTPStats(c *cod.Context) (err error) {
	params := &httpStatsParams{}
	err = validate.Do(params, c.Query())
	if err != nil {
		return
	}
	http := umbra.HTTP{
		URL: params.URL,
		IP:  params.IP,
	}
	c.Body = http.GetCheckResult()
	return
}
