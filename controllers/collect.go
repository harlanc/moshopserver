package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
)

func isUserHasCollect(userId, typeId, valueId int) bool {

	o := orm.NewOrm()

	var collect models.NideshopCollect
	collecttable := new(models.NideshopCollect)

	err := o.QueryTable(collecttable).Filter("type_id", typeId).Filter("value_id", valueId).Filter("user_id", userId).One(&collect)

	return err == nil

}
