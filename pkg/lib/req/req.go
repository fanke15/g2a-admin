package req

import (
	"github.com/fanke15/g2a-admin/pkg/basic"
	"github.com/imroc/req/v3"
	"os"
	"time"
)

var (
	reqClient = &req.Client{}
)

func init() {
	opt := &req.DumpOptions{
		Output:         os.Stdout,
		RequestHeader:  false,
		ResponseBody:   false,
		RequestBody:    false,
		ResponseHeader: false,
		Async:          false,
	}
	reqClient = req.C().SetTimeout(basic.One * time.Minute).SetCommonDumpOptions(opt).EnableDumpAll()
}

func InitReq() *req.Request {
	return reqClient.R()
}
