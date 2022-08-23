package router

import (
	"github.com/fanke15/g2a-admin/pkg/lib/bolt"
	"github.com/fanke15/g2a-admin/web/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(e *gin.Engine) {
	// 设置默认页面
	e.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/web/login")
	})

	//-----------------------WEB-----------------------
	webPage := e.Group("web")
	{
		webPage.GET("login", middleware.InitAce("tmpl", "index", map[string]interface{}{
			"title":  "用户登录",
			"config": string(bolt.InitBolt().Query("login")),
		}))

		webPage.GET("dashboard", middleware.InitAce("tmpl", "index", map[string]interface{}{
			"config": string(bolt.InitBolt().Query("dashboard")),
		}))
	}

	//-----------------------API-----------------------
	api := e.Group("api")
	{
		conf := api.Group("config")
		{
			conf.GET("dashboard/chart", func(c *gin.Context) {
				c.String(http.StatusOK, string(bolt.InitBolt().Query("dashboardChart")))
			})
			conf.GET("account/info", func(c *gin.Context) {
				c.String(http.StatusOK, string(bolt.InitBolt().Query("accountInfo")))
			})
			conf.GET("account/asset", func(c *gin.Context) {
				c.String(http.StatusOK, string(bolt.InitBolt().Query("accountAsset")))
			})
		}
	}
}
