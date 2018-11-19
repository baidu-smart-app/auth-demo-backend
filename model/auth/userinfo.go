package auth

import (
	"encoding/json"

	authdata "icode.baidu.com/baidu/searchbox/ma-srvdemo/data/auth"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/httpserver"
	"icode.baidu.com/baidu/searchbox/ma-srvdemo/lib/rsp"
)

// DecryptUserData 解密用户数据
// https://smartprogram.baidu.com/docs/develop/api/open_login/#%E7%94%A8%E6%88%B7%E6%95%B0%E6%8D%AE%E7%9A%84%E7%AD%BE%E5%90%8D%E9%AA%8C%E8%AF%81%E5%92%8C%E5%8A%A0%E8%A7%A3%E5%AF%86
func DecryptUserData(ctx *httpserver.Context, data, iv, openID string, conf *Config) (user *User, err error) {
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

	user = &User{}
	err = json.Unmarshal([]byte(content), user)
	if err != nil {
		ctx.Warning(err)
		return
	}

	return
}

// {"openid":"open_id","nickname":"baidu_user","headimgurl":"url of image","sex":1}
type User struct {
	OpenID     string `json:"openid"`
	NickName   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
	Sex        int8   `json:"sex"`
}
