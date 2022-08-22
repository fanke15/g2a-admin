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
		c.Redirect(http.StatusPermanentRedirect, "/page/login")
	})

	//-----------------------WEB-----------------------
	webPage := e.Group("page")
	{
		webPage.GET("login", middleware.InitAce("tmpl", "index", map[string]interface{}{
			"title":  "用户登录",
			"config": string(bolt.InitBolt().Query("login")),
		}))
	}

	//-----------------------API-----------------------

}
