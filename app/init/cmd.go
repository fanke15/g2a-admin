package cmd

import (
	"github.com/fanke15/g2a-admin/pkg/lib/conf"
	"github.com/urfave/cli"
	"os"
	"sort"
)

type (
	Cmd interface {
		Command() cli.Command
	}
	All struct{}
)

func New() {
	app := cli.NewApp()
	app.Name = conf.New().GetString("project.name")
	app.Usage = conf.New().GetString("project.description")

	app.Commands = []cli.Command{
		buildCommand(new(All)),
		buildCommand(new(Api)),
		buildCommand(new(Cron)),
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func (a *All) Command() cli.Command {
	return cli.Command{
		Name:    "all",
		Aliases: []string{"all"},
		Usage:   "all start",
		Subcommands: []cli.Command{
			{
				Name:   "start",
				Usage:  "启动所有服务",
				Action: a.RunAll,
			},
		},
	}
}

//---------------------------内部私有方法---------------------------//

func (a *All) RunAll(c *cli.Context) {
	(&Cron{}).RunCorn(c)

	(&Api{}).RunApi(c)
}

func buildCommand(cmd Cmd) cli.Command {
	return cmd.Command()
}
