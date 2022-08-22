package cmd

import (
	"github.com/fanke15/g2a-admin/pkg/basic"
	"github.com/fanke15/g2a-admin/pkg/lib/conf"
	"github.com/fanke15/g2a-admin/web/router"
	"github.com/fanke15/g2a-admin/web/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"time"
)

type Api struct {
}

func (a *Api) Command() cli.Command {
	return cli.Command{
		Name:    "api",
		Aliases: []string{"api"},
		Usage:   "api start",
		Subcommands: []cli.Command{
			{
				Name:   "start",
				Usage:  "开启运行api服务",
				Action: a.RunApi,
			},
		},
	}
}

//---------------------------内部私有方法---------------------------//

func (a *Api) RunApi(c *cli.Context) {
	engine := gin.New()
	engine.Use(middleware.Cors())

	engine.Static("assets", conf.New().GetString("project.dir.static"))
	engine.StaticFile("favicon.ico", basic.AnySliceToStr(basic.StrNull, conf.New().GetString("project.dir.static"), "img/favicon.ico"))

	router.Router(engine)
	var port = basic.AnySliceToStr(basic.StrNull, basic.StrColon, conf.New().GetString("project.port"))
	if err := engine.Run(port); err != nil {
		panic(err)
	}

	time.Sleep(basic.One * time.Second)
}
