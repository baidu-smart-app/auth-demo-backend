package auth

import (
	"auth-demo-backend/lib/httpserver"
	"auth-demo-backend/lib/rsp"
	"auth-demo-backend/model/auth"
	"auth-demo-backend/resource"
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
