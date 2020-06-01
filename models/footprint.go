package models

import (
	"github.com/astaxie/beego/orm"
	"moshopserver/utils"
)

func AddFootprint(userId, goodsId int) {

	o := orm.NewOrm()
	if userId > 0 && goodsId > 0 {
		footprintval := NideshopFootprint{GoodsId: goodsId, UserId: userId, AddTime: utils.GetTimestamp()}
		o.Insert(&footprintval)
	}
}
