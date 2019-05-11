package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/moshopserver/services"
)

type BaseController struct {
	beego.Controller
}

var userId string

func (this *BaseController) init() {
	token := this.Ctx.Input.Header("x-nideshop-token")

	userId = services.GetUserID(token)

	controller, action := this.GetControllerAndAction()

	publiccontrollerlist := beego.AppConfig.String("controller::publicController")
	publicactionlist := beego.AppConfig.String("action::publicAction")

	if !strings.Contains(publiccontrollerlist, controller) && !strings.Contains(publicactionlist, action) {
		if userId == "" {
			this.Abort("401")
		}
	}
}

func getLoginUserId() string {
	return userId
}
