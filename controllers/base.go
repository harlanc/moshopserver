package controllers

import (
	"github.com/moshopserver/services"
	"github.com/moshopserver/utils"
)

// type BaseController struct {
// 	beego.Controller
// }

// var userId string

// //https://blog.csdn.net/hzwy23/article/details/53314306过滤器

// func (this *BaseController) init() {
// 	token := this.Ctx.Input.Header("x-nideshop-token")

// 	userId = services.GetUserID(token)

// 	controller, action := this.GetControllerAndAction()

// 	publiccontrollerlist := beego.AppConfig.String("controller::publicController")
// 	publicactionlist := beego.AppConfig.String("action::publicAction")

// 	if !strings.Contains(publiccontrollerlist, controller) && !strings.Contains(publicactionlist, action) {
// 		if userId == "" {
// 			this.Abort("401")
// 		}
// 	}
// }

func getLoginUserId() int {
	intuserId := utils.String2Int(services.LoginUserId)
	return intuserId
}
