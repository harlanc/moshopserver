package services

import (
	"crypto"

	"github.com/astaxie/beego/httplib"
	"github.com/moshopserver/utils"
)

type LoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func Login(code string, fullUserInfo string) {

	config, _ := utils.GetConfig()

	//https://developers.weixin.qq.com/miniprogram/dev/api-backend/auth.code2Session.html
	req := httplib.Get("https://api.weixin.qq.com/sns/jscode2session")
	req.Param("grant_type", "authorization_code")
	req.Param("js_code", code)
	req.Param("secret", config.Wx.Secret)
	req.Param("appid", config.Wx.Appid)

	var res LoginResponse
	req.ToJSON(&res)

	crypto.Hash()

}
