package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	_ "github.com/moshopserver/models"
	_ "github.com/moshopserver/routers"
)

func main() {

	//beego.Router("api/index/index")
	o := orm.NewOrm()

	user := models.NideshopAttributeCategory{Name: "slene", Id: 1245, Enabled: 1}

	// insert
	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
		fmt.Println(id)

	}

	beego.Run() // listen and serve on 0.0.0.0:8080

}
