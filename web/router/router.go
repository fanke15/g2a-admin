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
		c.Redirect(http.StatusPermanentRedirect, "/web/dashboard")
	})

	//-----------------------WEB-----------------------
	webPage := e.Group("web")
	{
		webPage.GET("login", middleware.InitAce(map[string]interface{}{
			"title": "用户登录",
			"conf":  string(bolt.InitBolt().Query("login")),
		}))

		webPage.GET("dashboard", middleware.InitAce(map[string]interface{}{
			"conf": string(bolt.InitBolt().Query("dashboard")),
		}))
	}

	//-----------------------CONF-----------------------
	pageConf := e.Group("conf")
	{
		pageConf.GET("dashboard/chart", func(c *gin.Context) {
			c.String(http.StatusOK, string(bolt.InitBolt().Query("dashboard_chart")))
		})
		pageConf.GET("account/info", func(c *gin.Context) {
			c.String(http.StatusOK, string(bolt.InitBolt().Query("account_info")))
		})
		pageConf.GET("account/asset", func(c *gin.Context) {
			c.String(http.StatusOK, string(bolt.InitBolt().Query("account_asset")))
		})
	}

	//-----------------------API-----------------------
	api := e.Group("api")
	{
		api.POST("login")
	}

}
