package main

import (
	"github.com/fanke15/g2a-admin/app/init"
	"github.com/fanke15/g2a-admin/pkg/lib/bolt"
	"github.com/fanke15/g2a-admin/pkg/lib/conf"
	"github.com/fanke15/g2a-admin/pkg/lib/log"
)

func main() {
	/**启用附加服务**/
	conf.New()
	log.New()
	bolt.New()

	/**启动应用**/
	cmd.New()
}
