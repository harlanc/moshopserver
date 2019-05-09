package services

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/astaxie/beego/httplib"
	"github.com/moshopserver/utils"
)

type WXLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//https://developers.weixin.qq.com/miniprogram/dev/api/wx.getUserInfo.html

type Watermark struct {
	AppID     string `json:"appid"`
	TimeStamp int    `json:"timestamp"`
}

type WXUserInfo struct {
	OpenID    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	AvatarUrl string    `json:"avatarUrl"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	UnionID   string    `json:"unionId"`
	Watermark Watermark `json:"watermark"`
}

type ResUserInfo struct {
	UserInfo      WXUserInfo `json:"userInfo"`
	RawData       string     `json:"rawData"`
	Signature     string     `json:"signature"`
	EncryptedData string     `json:"encryptedData"`
	IV            string     `json:"iv"`
}

func Login(code string, fullUserInfo string) *WXUserInfo {

	var resUserInfo ResUserInfo

	err := json.Unmarshal([]byte(fullUserInfo), &resUserInfo)

	if err != nil {
		return nil
	}

	config, _ := utils.GetConfig()

	//https://developers.weixin.qq.com/miniprogram/dev/api-backend/auth.code2Session.html
	req := httplib.Get("https://api.weixin.qq.com/sns/jscode2session")
	req.Param("grant_type", "authorization_code")
	req.Param("js_code", code)
	req.Param("secret", config.Wx.Secret)
	req.Param("appid", config.Wx.Appid)

	var res WXLoginResponse
	req.ToJSON(&res)

	s := sha1.New()
	s.Write([]byte(resUserInfo.RawData + res.SessionKey))
	sha1hash := hex.EncodeToString(s.Sum(nil))

	if resUserInfo.Signature != sha1hash {
		return nil
	}
	userinfo := DecryptUserInfoData(res.SessionKey, resUserInfo.EncryptedData, resUserInfo.IV)

	return &userinfo

}

func DecryptUserInfoData(sessionKey string, encryptedData string, iv string) WXUserInfo {

	sk, _ := base64.StdEncoding.DecodeString(sessionKey)
	ed, _ := base64.StdEncoding.DecodeString(encryptedData)
	i, _ := base64.StdEncoding.DecodeString(iv)

	decryptedData, err := utils.AesCBCDecrypt(ed, sk, i)

	if err != nil {

	}

	var wxuserinfo WXUserInfo
	err = json.Unmarshal([]byte(decryptedData), &wxuserinfo)
	if err != nil {

	}
	return wxuserinfo
}
