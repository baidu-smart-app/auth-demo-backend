package auth

import (
	"encoding/json"

	authdata "auth-demo-backend/data/auth"
	"auth-demo-backend/lib/httpserver"
	"auth-demo-backend/lib/rsp"
)

// DecryptUserData 解密用户数据
// https://smartprogram.baidu.com/docs/develop/api/open_login/#%E7%94%A8%E6%88%B7%E6%95%B0%E6%8D%AE%E7%9A%84%E7%AD%BE%E5%90%8D%E9%AA%8C%E8%AF%81%E5%92%8C%E5%8A%A0%E8%A7%A3%E5%AF%86
func DecryptUPhoneData(ctx *httpserver.Context, data, iv, openID string, conf *Config) (phone *Phone, err error) {
	sessionKey, err := authdata.DefaultAuth.GetSessionKeyByOpenID(openID)
	if err != nil {
		ctx.Warning(err)
		return
	}

	if sessionKey == "" {
		err = rsp.AuthNoSessionKey
		ctx.Warning(err)
		return
	}

	content, err := Decrypt(data, sessionKey, iv, conf.AppKey)
	if err != nil {
		ctx.Warning(err)
		return
	}

	phone = &Phone{}
	err = json.Unmarshal([]byte(content), phone)
	if err != nil {
		ctx.Warning(err)
		return
	}

	return
}

type Phone map[string]interface{}
