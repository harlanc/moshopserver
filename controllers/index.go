package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	o := orm.NewOrm()
	banner := models.NideshopAd{AdPositionId: 1}

	err := o.Read(&banner)

	if err == orm.ErrNoRows {
		fmt.Println("Can not find.")
	} else if err == orm.ErrMissPK {
		fmt.Println("Can not find main key.")
	} else {
		fmt.Println(banner.Id, banner.Name)
	}

	this.Data["banner"] = banner

}
