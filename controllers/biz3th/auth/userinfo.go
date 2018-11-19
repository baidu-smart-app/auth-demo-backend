package auth

import (
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/rsp"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/model/auth"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/resource"
)

// GetUserInfo swan.getUserInfo 回调 进行解密 + 账户管理示例
// https://smartapp.baidu.com/docs/develop/api/open_userinfo/#getUserInfo/
func GetUserInfo(ctx *httpserver.Context) interface{} {
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

	user, err := auth.DecryptUserData(ctx, data.Data, data.IV, data.OpenID, resource.C.SmartApp)
	if err != nil {
		return err
	}

	return user
}
