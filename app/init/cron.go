package cmd

import (
	"github.com/fanke15/g2a-admin/pkg/lib/cron"
	"github.com/fanke15/g2a-admin/pkg/lib/log"
	"github.com/fanke15/g2a-admin/web"
	"github.com/urfave/cli"
	"time"
)

type (
	Cron struct{}
)

// 设置终端输出
func (c *Cron) Command() cli.Command {
	return cli.Command{
		Name:    "cron",
		Aliases: []string{"-r"},
		Usage:   "cron start",
		Subcommands: []cli.Command{
			{
				Name:   "test",
				Usage:  "示例定时任务",
				Action: c.Test,
			},
			{
				Name:   "start",
				Usage:  "执行定时任务监控",
				Action: c.RunCorn,
			},
		},
	}
}

//---------------------------内部私有方法---------------------------//
func (c *Cron) Test(cl *cli.Context) {
	log.Info("cron test start--------------->")

	time.Sleep(30 * time.Minute)
}

func (c *Cron) RunCorn(cl *cli.Context) {
	log.Info("cron start--------------->")

	web.ListenAmisConfigData() // 启动时加载

	//阻塞执行
	cron.New().Start()
}
