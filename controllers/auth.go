package controllers

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"

	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/services"
)

func (this *CatalogController) Auth_loginByWeixin() {
	code := this.GetString("code")
	fulluserinfo := this.GetString("userInfo")
	clientIP := this.Ctx.Input.IP

	userInfo := services.Login(code, fulluserinfo)
	if userInfo == nil {

	}

	o := orm.NewOrm()

	var users []orm.Params
	usertable := new(models.NideshopCategory)
	nums, err := o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).Values(&users, "id")
	if nums == 0 {
		uuid, err := uuid.NewV4()
		newuser := models.NideshopUser{Username: uuid.String(),Password:"",RegisterTime:}

	}

	var categories []models.NideshopCategory
	category := new(models.NideshopCategory)
	o.QueryTable(category).Filter("parent_id", 0).Limit(10).All(&categories)

	data, err := json.Marshal(CateLogRtnJson{categories, getCurCategory(categoryId)})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}
