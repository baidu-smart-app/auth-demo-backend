package auth

import (
	authdata "icode.baidu.com/baidu/searchbox/ma-srvdemo/data/auth"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httputil"
)

func Login(ctx *httpserver.Context, code string, conf *Config) (string, error) {
	rsp := &Code2SessionKeyRsp{}

	err := httputil.NewClient().
		SetPostParam("code", code).
		SetPostParam("client_id", conf.AppKey).
		SetPostParam("sk", conf.SecrectKey).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetPath("https://openapi.baidu.com/nalogin/getSessionKeyByCode").
		PostMethod().
		Do().
		RspJson(rsp, func(data interface{}) error {
			if data.(*Code2SessionKeyRsp).ErrorNo != 0 {
				return httputil.BadContent
			}

			return nil
		}).
		Error()
	if err != nil {
		ctx.Warning(err)
		return "", err
	}

	e := authdata.DefaultAuth.SetOpenID2SessionKey(rsp.OpenID, rsp.SessionKey)
	if e != nil {
		ctx.Warning(e)
		return "", e
	}

	return rsp.OpenID, nil
}

type Code2SessionKeyRsp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`

	Error            string `json:"error"`
	ErrorNo          int64  `json:"errno"`
	ErrorDescription string `json:"error_description"`
}
