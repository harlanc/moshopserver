package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type CartController struct {
	beego.Controller
}

type CartTotal struct {
	GoodsCount         int
	GoodsAmount        int
	CheckedGoodsCount  int
	CheckedGoodsAmount int
}

type IndexCartData struct {
	CartList  []models.NideshopCart
	CartTotal CartTotal
}

type AddCartData struct {
	GoodsId                   string
	ProductId                 string
	GoodsSN                   string
	GoodsName                 string
	ListPicUrl                string
	Number                    string
	SessionId                 int
	UserId                    int
	RetailPrice               int
	MarketPrice               int
	GoodsSpecifitionNameValue string
	GoodsSpecifitionIds       string
	Checked                   int
}

func getCart() IndexCartData {

	o := orm.NewOrm()
	carttable := new(models.NideshopCart)
	var carts []models.NideshopCart
	o.QueryTable(carttable).Filter("user_id", getLoginUserId()).Filter("session_id", 1).All(carts)

	var goodsCount int
	var goodsAmount int
	var checkedGoodsCount int
	var checkedGoodsAmount int

	for _, val := range carts {
		goodsCount += val.Number
		goodsAmount += val.Number * val.RetailPrice
		if val.Checked == 0 {
			checkedGoodsCount += val.Number
			checkedGoodsAmount += val.Number * val.RetailPrice
		}

		goodstable := new(models.NideshopGoods)
		var goods models.NideshopGoods
		o.QueryTable(goodstable).Filter("id", val.GoodsId).One(&goods)
		val.ListPicUrl = goods.ListPicUrl
	}

	return IndexCartData{carts, CartTotal{goodsCount, goodsAmount, checkedGoodsCount, checkedGoodsAmount}}
}

func (this *CartController) Cart_Index() {

	data, err := json.Marshal(getCart())
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *CartController) Cart_Add() {

	goodsId := this.GetString("goodsId")
	productId := this.GetString("productId")
	number := this.GetString("number")

	intgoodsId := utils.String2Int(goodsId)
	intproductId := utils.String2Int(productId)
	intnumber := utils.String2Int(number)

	o := orm.NewOrm()
	goodstable := new(models.NideshopGoods)
	var goods models.NideshopGoods
	err := o.QueryTable(goodstable).Filter("id", goodsId).One(&goods)
	if err == orm.ErrNoRows || goods.IsDelete {
		this.CustomAbort(400, "商品已下架")
	}

	producttable := new(models.NideshopProduct)
	var product models.NideshopProduct
	err = o.QueryTable(producttable).Filter("goods_id", goodsId).Filter("id", productId).One(&product)
	if err == orm.ErrNoRows || product.GoodsNumber < intnumber {
		this.CustomAbort(400, "库存不足")
	}

	carttable := new(models.NideshopCart)
	var cart models.NideshopCart
	err = o.QueryTable(carttable).Filter("goods_id", goodsId).Filter("product_id", productId).One(&cart)

	if err == orm.ErrNoRows {
		var goodsSepcifitionValues []orm.Params

		if product.GoodsSpecificationIds != "" {
			goodsspecitable := new(models.NideshopGoodsSpecification)
			goodsspecificationids := strings.Split(product.GoodsSpecificationIds, "_")
			var intgoodsspecificationids []int
			for _, val := range goodsspecificationids {
				intgoodsspecificationids = append(intgoodsspecificationids, utils.String2Int(val))
			}
			o.QueryTable(goodsspecitable).Filter("goods_id", goodsId).Filter("id__in", intgoodsspecificationids).
				Values(&goodsSepcifitionValues, "value")
		}

		vals := utils.ExactMapValues2StringArray(goodsSepcifitionValues, "Value")
		cartData := models.NideshopCart{GoodsId: intgoodsId, ProductId: intproductId,
			GoodsSn: product.GoodsSn, GoodsName: goods.Name, ListPicUrl: goods.ListPicUrl,
			Number: intnumber, SessionId: "1", UserId: getLoginUserId(), RetailPrice: product.RetailPrice,
			MarketPrice: product.RetailPrice, GoodsSpecifitionNameValue: strings.Join(vals, ";"),
			GoodsSpecifitionIds: product.GoodsSpecificationIds, Checked: 1}
		o.Insert(&cartData)
	} else {
		if product.GoodsNumber < (intnumber + cart.Number) {
			this.CustomAbort(400, "库存不足")
		}
		o.QueryTable(carttable).Update(orm.Params{"number": orm.ColValue(orm.ColAdd, intnumber)})
	}

	data, err := json.Marshal(getCart())
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}
