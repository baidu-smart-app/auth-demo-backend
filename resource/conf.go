package resource

import (
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/conf"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/log"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/model/auth"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/model/pay"
)

// Config 全局配置
type Config struct {
	Server            *httpserver.Config
	Log               *log.Config
	Pay               *pay.Config
	SmartApp          *auth.Config
	SelfRsaPrivKey    []byte
	PlatformRsaPubKey []byte
}

// C 全局配置单例
var C = &Config{
	Server:   &httpserver.Config{},
	Log:      &log.Config{},
	Pay:      &pay.Config{},
	SmartApp: &auth.Config{},
}

var ZERO = struct{}{}

// Init 加载全局配置
func Init(file2data map[string]interface{}) error {
	for file, data := range file2data {
		if err := conf.LoadJSON(file, data); err != nil {
			return err
		}
	}

	return nil
}
