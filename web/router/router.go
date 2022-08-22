package router

import (
	"github.com/fanke15/g2a-admin/pkg/basic"
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

		webPage.GET("dashboard", middleware.InitAce("tmpl", "dashboard", map[string]interface{}{
			"title":  "dashboard",
			"config": string(bolt.InitBolt().Query("dashboard")),
		}))
	}

	//-----------------------API-----------------------
	api := e.Group("api")
	{
		acc := api.Group("account")
		{
			acc.POST("home", func(c *gin.Context) {

				a := `
{"type":"page","body":{"type":"property","title":"Information","items":[{"label":"system","content":"Linux"},{"label":"python","content":"3.7.9"},{"label":"program","content":"fastapi-amis-admin"},{"label":"version","content":"0.1.4"},{"label":"license","content":"Apache2.0"}]}}

`

				c.String(200, string(basic.Marshal(a)))
			})
		}

	}

}
