package middleware

import (
	"github.com/fanke15/g2a-admin/pkg/basic"
	"github.com/fanke15/g2a-admin/pkg/lib/conf"
	"github.com/fanke15/g2a-admin/pkg/lib/log"
	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

var (
	aceDir = basic.AnySliceToStr(basic.StrNull, conf.New().GetString("project.dir.static"), "ace/")
)

// 模板初始化设置
func InitAce(data interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		tpl, err := ace.Load(basic.AnySliceToStr(basic.StrNull, aceDir, "index"), "", &ace.Options{DynamicReload: true})
		if err != nil {
			log.Error(err.Error(), "func", "InitAce")
			return
		}
		_ = tpl.Execute(c.Writer, data)
	}
}
