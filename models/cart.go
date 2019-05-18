package models

import "github.com/astaxie/beego/orm"

func ClearBuyGoods(userid int) {

	o := orm.NewOrm()

	carttable := new(NideshopCart)

	o.QueryTable(carttable).Filter("user_id", userid).Filter("session_id", 1).Filter("checked", 1).Delete()

}
