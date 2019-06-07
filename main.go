package main

import (
	"github.com/astaxie/beego"
	_ "github.com/moshopserver/models"
	_ "github.com/moshopserver/routers"
	_ "github.com/moshopserver/utils"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = "192.168.0.104"
	beego.BConfig.Listen.HTTPPort = 8080

	//beego.InsertFilter("/api/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080

}
