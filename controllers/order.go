package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) Order_List() {

	o := orm.NewOrm()
	ordertable := new(models.NideshopOrder)
	var orders []models.NideshopOrder
	o.QueryTable(ordertable).Filter("user_id", getLoginUserId()).All(&orders)

	firstpagedorders := utils.GetPageDataInterface(orders, 1, 10)
	var newpageorders []models.NideshopOrder

	for _, val := range firstpagedorders.Data.([]models.NideshopOrder) {

	}
}
