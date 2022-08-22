package conf

import (
	"github.com/fanke15/g2a-admin/pkg/basic"
	"github.com/spf13/viper"
	"strings"
)

type ViperClient struct {
	*viper.Viper
}

var (
	vc = &ViperClient{}
)

func New() *ViperClient {
	if vc.Viper == nil {
		conn()
	}
	return vc
}

//---------------------------内部私有方法---------------------------//

// 初始化
func conn() {

	vc = &ViperClient{viper.New()} // 初始化

	vc.SetConfigName("conf")           // 设置配置文件名称
	vc.AddConfigPath("./")             // 设置配置文件所在目录
	vc.SetConfigType(basic.TomlSuffix) // 设置配置文件后缀类型

	if err := vc.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	vc.AutomaticEnv()                                                           // 匹配环境变量
	vc.SetEnvKeyReplacer(strings.NewReplacer(basic.StrDat, basic.StrUnderline)) // 设置配置名称与环境变量匹配分隔符
	return
}
