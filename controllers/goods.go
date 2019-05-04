package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type GoodsController struct {
	beego.Controller
}

type SkuRtnJson struct {
	ProductList       []models.NideshopProduct
	SpecificationList []models.SpecificationItem
}

func (this *GoodsController) Goods_Index() {
	o := orm.NewOrm()

	var goods []models.NideshopGoods
	good := new(models.NideshopGoods)
	o.QueryTable(good).All(&goods)

	data, err := json.Marshal(goods)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Sku() {

	goodsId := this.GetString("id")
	goodsId_int := utils.String2Int(goodsId)

	plist := models.GetProductList(goodsId_int)
	slist := models.GetSpecificationList(goodsId_int)

	data, err := json.Marshal(SkuRtnJson{ProductList: plist, SpecificationList: slist})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Detail() {
	goodsId := this.GetString("id")
	intGoodsId := utils.String2Int(goodsId)

	o := orm.NewOrm()

	var goods []models.NideshopGoods
	good := new(models.NideshopGoods)
	o.QueryTable(good).Filter("id", intGoodsId).All(&goods)

	var galleries []models.NideshopGoodsGallery
	gallery := new(models.NideshopGoodsGallery)
	o.QueryTable(gallery).Filter("goods_id", intGoodsId).Limit(4).All(&galleries)

	qb, _ := orm.NewQueryBuilder("mysql")

	var attributes []models.NideshopAttribute

	qb.Select("gs.*", "s.name").
		From("nideshop_goods_specification gs").
		InnerJoin("nideshop_specification s").On("gs.specification_id = s.id").
		Where("gs.specification_id =" + utils.Int2String(goodsId))

	sql := qb.String()

	o := orm.NewOrm()
	o.Raw(sql, 20).QueryRows(&specifications)

}
