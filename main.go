package main

import (
	"github.com/astaxie/beego"
	_ "github.com/moshopserver/models"
	_ "github.com/moshopserver/routers"
	_ "github.com/moshopserver/utils"
)

func main() {

	//beego.Router("api/index/index")
	// o := orm.NewOrm()

	// user := models.NideshopAttributeCategory{Name: "slene", Id: 1245, Enabled: 1}

	// // insert
	// id, err := o.Insert(&user)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(id)

	// }
	// config, _ := utils.GetConfig()

	// fmt.Println(config)

	//controller := controllers.IndexController
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.Run() // listen and serve on 0.0.0.0:8080

}
