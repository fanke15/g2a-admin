package web

import (
	"embed"

	"github.com/fanke15/g2a-admin/pkg/lib/bolt"

	"github.com/fanke15/g2a-admin/pkg/lib/log"

	"io/fs"
	"path/filepath"

	"strings"
)

const ()

//go:embed static/json
var systemStatic embed.FS

// 页面配置信息缓存 启动时加载
func ListenAmisConfigData() {
	log.Info(">>>>>>>>>>>>>>>>>>>>>>>>>页面配置数据缓存开始！")

	var data = make(map[string][]byte)

	// 读取配置页面数据
	if err := fs.WalkDir(systemStatic, "static/json", func(p string, d fs.DirEntry, e error) error {
		// check if ext is json
		if d.IsDir() {
			return nil
		}
		_, filename := filepath.Split(p)
		fileExt := filepath.Ext(filename)
		if fileExt != ".json" {
			return nil
		}
		// write page data to db
		pageData, e := systemStatic.ReadFile(p)
		if e != nil {
			return e
		}
		data[strings.TrimSuffix(filename, fileExt)] = pageData
		return nil
	}); err != nil {
		log.Error(err.Error(), "func", "ListenAmisConfigData", "msg", "读取页面配置数据失败！")
		return
	}

	if err := bolt.InitBolt().SaveBatch(data); err != nil {
		log.Error(err.Error(), "func", "ListenAmisConfigData", "msg", "批量缓存页面配置数据失败！")
		return
	}
	log.Info(">>>>>>>>>>>>>>>>>>>>>>>>>页面配置数据缓存完成！", "dataLen", len(data))
}

//---------------------------内部私有方法---------------------------//
