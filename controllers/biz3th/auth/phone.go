package auth

import (
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/rsp"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/model/auth"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/resource"
)

// GetUserInfo swan.getUserInfo 回调 进行解密 + 账户管理示例
// https://smartprogram.baidu.com/docs/develop/component/formlist/#button/
func GetPhone(ctx *httpserver.Context) interface{} {
	data := &struct {
		Data   string `json:"data"`
		IV     string `json:"iv"`
		OpenID string `json:"open_id"`
	}{}
	if err := ctx.ReqJson(data); err != nil {
		return err
	}

	if data.Data == "" || data.IV == "" || data.OpenID == "" {
		return rsp.ParamIllegal
	}

	phone, err := auth.DecryptUPhoneData(ctx, data.Data, data.IV, data.OpenID, resource.C.SmartApp)
	if err != nil {
		return err
	}

	return phone
}
