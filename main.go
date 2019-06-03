package main

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/moshopserver/models"
	_ "github.com/moshopserver/routers"
	"github.com/moshopserver/services"
	_ "github.com/moshopserver/utils"
)

func getControllerAndAction(rawvalue string) (controller, action string) {
	vals := strings.Split(rawvalue, "/")
	return vals[2], vals[2] + "/" + vals[3]
}

var UserId string

func FilterFunc(ctx *context.Context) {

	token := ctx.Input.Header("x-nideshop-token")
	UserId = services.GetUserID(token)
	controller, action := getControllerAndAction(ctx.Request.RequestURI)

	publiccontrollerlist := beego.AppConfig.String("controller::publicController")
	publicactionlist := beego.AppConfig.String("action::publicAction")

	if !strings.Contains(publiccontrollerlist, controller) && !strings.Contains(publicactionlist, action) {
		if UserId == "" {
			ctx.Redirect(401, "auth/loginByWeixin")
			//http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
		}
	}
}

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = "192.168.0.102"
	beego.BConfig.Listen.HTTPPort = 8080

	beego.InsertFilter("/api/*", beego.BeforeExec, FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080

}
