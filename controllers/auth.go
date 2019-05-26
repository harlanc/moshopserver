package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/services"
	"github.com/moshopserver/utils"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Auth_LoginByWeixin() {

	code := this.GetString("code")
	fulluserinfo := this.GetString("userInfo")
	clientIP := this.Ctx.Input.IP()

	userInfo := services.Login(code, fulluserinfo)
	if userInfo == nil {

	}

	o := orm.NewOrm()

	var user models.NideshopUser
	usertable := new(models.NideshopCategory)
	err := o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	if err != nil {
		newuser := models.NideshopUser{Username: utils.GetUUID(), Password: "", RegisterTime: utils.GetTimestamp(),
			RegisterIp: clientIP, Mobile: "", WeixinOpenid: userInfo.OpenID, Avatar: userInfo.AvatarUrl, Gender: userInfo.Gender,
			Nickname: userInfo.NickName}
		o.Insert(&newuser)
		o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	}

	userinfo := make(map[string]interface{})
	userinfo["id"] = user.Id
	userinfo["username"] = user.Username
	userinfo["nickname"] = user.Nickname
	userinfo["gender"] = user.Gender
	userinfo["avatar"] = user.Avatar
	userinfo["birthday"] = user.Birthday

	user.LastLoginIp = clientIP
	user.LastLoginTime = utils.GetTimestamp()

	if _, err := o.Update(&user); err == nil {

	}

	sessionKey := services.Create(utils.Int2String(user.Id))

	rtnInfo := make(map[string]interface{})
	rtnInfo["sessionKey"] = sessionKey
	rtnInfo["userInfo"] = userinfo

	data, err := json.Marshal(rtnInfo)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}
