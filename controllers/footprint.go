package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/bradfitz/slice"
	"moshopserver/models"
	"moshopserver/utils"
)

type FootprintController struct {
	beego.Controller
}

func (this *FootprintController) Footprint_Delete() {

	footprintId := this.GetString("footprintId")
	intfootprintid := utils.String2Int(footprintId)
	loginuserid := getLoginUserId()

	o := orm.NewOrm()
	footprinttable := new(models.NideshopFootprint)
	var footprint models.NideshopFootprint
	o.QueryTable(footprinttable).Filter("id", intfootprintid).Filter("user_id", loginuserid).One(&footprint)
	o.QueryTable(footprinttable).Filter("user_id", loginuserid).Filter("goods_id", footprint.GoodsId).Delete()

	//return this.
}

type FootprintListRtnJson struct {
	models.NideshopFootprint
	Name          string
	ListPicUrl    string
	GoodsBrief    string
	RetailPrice   float64
	AddTimeString string
}

func (this *FootprintController) Footprint_List() {
	qb, _ := orm.NewQueryBuilder("mysql")
	var listdata []FootprintListRtnJson

	o := orm.NewOrm()

	qb.Select("nf.*", "ng.name", "ng.list_pic_url", "ng.goods_brief", "ng.retail_price").
		From("nideshop_footprint nf").
		InnerJoin("nideshop_goods ng").
		On("nf.goods_id = ng.id").
		Where("nf.user_id =" + utils.Int2String(getLoginUserId())).
		OrderBy("id").Desc()
	sql := qb.String()
	o.Raw(sql).QueryRows(&listdata)

	var rvdata []FootprintListRtnJson
	var goodsids []int

	for _, item := range listdata {

		if !utils.ContainsInt(goodsids, item.GoodsId) {

			footprinttime := time.Unix(item.AddTime, 0)
			nowtime := time.Now()
			yesterdaytime := nowtime.Add(-24 * time.Hour)
			yesyesterdaytime := yesterdaytime.Add(-24 * time.Hour)

			if utils.DateEqual(footprinttime, nowtime) {
				item.AddTimeString = "今天"
			} else if utils.DateEqual(footprinttime, yesterdaytime) {
				item.AddTimeString = "昨天"
			} else if utils.DateEqual(footprinttime, yesyesterdaytime) {
				item.AddTimeString = "前天"
			}

			goodsids = append(goodsids, item.GoodsId)

			rvdata = append(rvdata, item)
		}
	}

	slice.Sort(rvdata, func(i, j int) bool {
		return rvdata[i].AddTime < rvdata[j].AddTime
	})

	utils.ReturnHTTPSuccess(&this.Controller, rvdata)
	this.ServeJSON()

}
