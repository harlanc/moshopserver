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
	GoodsCount         int     `json:"goodsCount"`
	GoodsAmount        float64 `json:"goodsAmount"`
	CheckedGoodsCount  int     `json:"checkedGoodsCount"`
	CheckedGoodsAmount float64 `json:"checkedGoodsAmount"`
}
type GoodsCount struct {
	CartTotal CartTotal `json:"cartTotal"`
}
type IndexCartData struct {
	CartList  []models.NideshopCart `json:"cartList"`
	CartTotal CartTotal             `json:"cartTotal"`
}

type GoodsSpecifition struct {
	models.NideshopGoodsSpecification
	Name string
}

func getCart() IndexCartData {

	o := orm.NewOrm()
	carttable := new(models.NideshopCart)
	var carts []models.NideshopCart
	o.QueryTable(carttable).Filter("user_id", getLoginUserId()).Filter("session_id", 1).All(&carts)

	var goodsCount int
	var goodsAmount float64
	var checkedGoodsCount int
	var checkedGoodsAmount float64

	for _, val := range carts {
		goodsCount += val.Number
		goodsAmount += float64(val.Number) * val.RetailPrice
		if val.Checked == 0 {
			checkedGoodsCount += val.Number
			checkedGoodsAmount += float64(val.Number) * val.RetailPrice
		}

		goodstable := new(models.NideshopGoods)
		var goods models.NideshopGoods
		o.QueryTable(goodstable).Filter("id", val.GoodsId).One(&goods)
		val.ListPicUrl = goods.ListPicUrl
	}

	return IndexCartData{carts, CartTotal{goodsCount, goodsAmount, checkedGoodsCount, checkedGoodsAmount}}
}

func (this *CartController) Cart_Index() {

	utils.ReturnHTTPSuccess(&this.Controller, getCart())
	this.ServeJSON()
}

type CartAddBody struct {
	GoodsId   int `json:"goodsId"`
	ProductId int `json:"productId"`
	Number    int `json:"number"`
}

func (this *CartController) Cart_Add() {

	var ab CartAddBody
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body, &ab)

	intgoodsId := ab.GoodsId
	intproductId := ab.ProductId
	intnumber := ab.Number

	o := orm.NewOrm()
	goodstable := new(models.NideshopGoods)
	var goods models.NideshopGoods
	err := o.QueryTable(goodstable).Filter("id", intgoodsId).One(&goods)
	if err == orm.ErrNoRows || goods.IsDelete {
		this.CustomAbort(400, "商品已下架")
	}

	producttable := new(models.NideshopProduct)
	var product models.NideshopProduct
	err = o.QueryTable(producttable).Filter("goods_id", intgoodsId).Filter("id", intproductId).One(&product)
	if err == orm.ErrNoRows || product.GoodsNumber < intnumber {
		this.CustomAbort(400, "库存不足")
	}

	carttable := new(models.NideshopCart)
	var cart models.NideshopCart
	err = o.QueryTable(carttable).Filter("goods_id", intgoodsId).Filter("product_id", intproductId).
		Filter("user_id", getLoginUserId()).One(&cart)

	if err == orm.ErrNoRows {
		var goodsSepcifitionValues []orm.Params

		if product.GoodsSpecificationIds != "" {
			goodsspecitable := new(models.NideshopGoodsSpecification)
			goodsspecificationids := strings.Split(product.GoodsSpecificationIds, "_")
			var intgoodsspecificationids []int
			for _, val := range goodsspecificationids {
				intgoodsspecificationids = append(intgoodsspecificationids, utils.String2Int(val))
			}
			o.QueryTable(goodsspecitable).Filter("goods_id", intgoodsId).Filter("id__in", intgoodsspecificationids).
				Values(&goodsSepcifitionValues, "value")
		}

		vals := utils.ExactMapValues2StringArray(goodsSepcifitionValues, "Value")
		cartData := models.NideshopCart{
			GoodsId:                   intgoodsId,
			ProductId:                 intproductId,
			GoodsSn:                   product.GoodsSn,
			GoodsName:                 goods.Name,
			ListPicUrl:                goods.ListPicUrl,
			Number:                    intnumber,
			SessionId:                 "1",
			UserId:                    getLoginUserId(),
			RetailPrice:               product.RetailPrice,
			MarketPrice:               product.RetailPrice,
			GoodsSpecifitionNameValue: strings.Join(vals, ";"),
			GoodsSpecifitionIds:       product.GoodsSpecificationIds,
			Checked:                   1}
		o.Insert(&cartData)
	} else {
		if product.GoodsNumber < (intnumber + cart.Number) {
			this.CustomAbort(400, "库存不足")
		}
		o.QueryTable(carttable).Update(orm.Params{"number": orm.ColValue(orm.ColAdd, intnumber)})
	}

	utils.ReturnHTTPSuccess(&this.Controller, getCart())
	this.ServeJSON()
}

func (this *CartController) Cart_Update() {

	goodsId := this.GetString("goodsId")
	productId := this.GetString("productId")
	number := this.GetString("number")
	id := this.GetString("id")

	intgoodsId := utils.String2Int(goodsId)
	intproductId := utils.String2Int(productId)
	intnumber := utils.String2Int(number)
	intid := utils.String2Int(id)

	o := orm.NewOrm()

	producttable := new(models.NideshopProduct)
	var product models.NideshopProduct
	errproduct := o.QueryTable(producttable).Filter("goods_id", goodsId).Filter("id", productId).One(&product)
	if errproduct == orm.ErrNoRows || product.GoodsNumber < intnumber {
		this.CustomAbort(400, "库存不足")
	}

	carttable := new(models.NideshopCart)
	var cart models.NideshopCart
	if cart.ProductId == intproductId {
		o.QueryTable(carttable).Filter("id", intid).Update(orm.Params{"number": intnumber})

		data, errjson := json.Marshal(getCart())
		if errjson != nil {
			this.Data["json"] = errjson
		} else {
			this.Data["json"] = json.RawMessage(string(data))
		}
		this.ServeJSON()
	}

	var newcart models.NideshopCart
	errcart := o.QueryTable(carttable).Filter("goods_id", intgoodsId).Filter("product_id", intproductId).One(&newcart)
	if errcart == orm.ErrNoRows {
		var goodsspecifitons []GoodsSpecifition
		if product.GoodsSpecificationIds != "" {
			goodsspecificationids := strings.Split(product.GoodsSpecificationIds, "_")
			qb, _ := orm.NewQueryBuilder("mysql")

			qb.Select("ngs.*", "ns.name").
				From("nideoshop_goods_specification ngs").
				InnerJoin("nideshop_specification ns").On("ns.id = ngs.specification_id").
				Where("ngs.goods_id =" + goodsId).And("ngs.id").In(strings.Join(goodsspecificationids, ","))
			sql := qb.String()
			o.Raw(sql).QueryRows(&goodsspecifitons)
		}

		goodsspecifitonsjson, _ := json.Marshal(goodsspecifitons)
		o.QueryTable(carttable).Filter("id", intid).Update(orm.Params{
			"number":                       intnumber,
			"goods_specifition_name_value": goodsspecifitonsjson,
			"retail_price":                 product.RetailPrice,
			"market_price":                 product.RetailPrice,
			"product_id":                   intproductId,
			"goods_sn":                     product.GoodsSn})

	} else {

		newNumber := intnumber + newcart.Number
		if errproduct == orm.ErrNoRows || product.GoodsNumber < newNumber {
			this.CustomAbort(400, "库存不足")
		}
		o.QueryTable(carttable).Filter("id", newcart.Id).Delete()
		o.QueryTable(carttable).Filter("id", intid).Update(orm.Params{
			"number":                       newNumber,
			"goods_specifition_name_value": newcart.GoodsSpecifitionNameValue,
			"goods_specifition_ids":        newcart.GoodsSpecifitionIds,
			"retail_price":                 newcart.RetailPrice,
			"market_price":                 newcart.RetailPrice,
			"product_id":                   intproductId,
			"goods_sn":                     product.GoodsSn})
	}

	utils.ReturnHTTPSuccess(&this.Controller, getCart())
	this.ServeJSON()
}

func (this *CartController) Cart_Checked() {

	isChecked := this.GetString("isChecked")
	productIds := this.GetString("productIds")

	intisChecked := utils.String2Int(isChecked)

	if productIds == "" {
		this.Abort("删除出错")
	}
	productIdsarray := strings.Split(productIds, ",")

	o := orm.NewOrm()
	carttable := new(models.NideshopCart)
	o.QueryTable(carttable).Filter("product_id__in", productIdsarray).Update(orm.Params{
		"checked": intisChecked,
	})

	utils.ReturnHTTPSuccess(&this.Controller, getCart())
	this.ServeJSON()
}

func (this *CartController) Cart_Delete() {

	productIds := this.GetString("productIds")
	if productIds == "" {
		this.Abort("删除出错")
	}
	productidsarray := strings.Split(productIds, ",")

	o := orm.NewOrm()
	carttable := new(models.NideshopCart)
	o.QueryTable(carttable).Filter("product_id__in", productidsarray).Delete()

	utils.ReturnHTTPSuccess(&this.Controller, getCart())
	this.ServeJSON()

}

func (this *CartController) Cart_GoodsCount() {

	cartData := getCart()
	goodscount := GoodsCount{CartTotal: CartTotal{GoodsCount: cartData.CartTotal.GoodsCount}}
	utils.ReturnHTTPSuccess(&this.Controller, goodscount)
	this.ServeJSON()
}

type CartAddress struct {
	models.NideshopAddress
	ProviceName  string
	CityName     string
	DistrictName string
	FullRegion   string
}

type CheckoutRtnJson struct {
	Address          CartAddress
	FreightPrice     float64
	CheckedCoupon    []models.NideshopUserCoupon
	CouponList       []models.NideshopUserCoupon
	CouponPrice      float64
	CheckedGoodsList []models.NideshopCart
	GoodsTotalPrice  float64
	OrderTotalPrice  float64
	ActualPrice      float64
}

func (this *CartController) Cart_Checkout() {

	addressId := this.GetString("addressId")
	intaddressid := utils.String2Int(addressId)

	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)
	var address CartAddress
	var err error
	if addressId != "" {
		err = o.QueryTable(addresstable).Filter("is_default", 1).Filter("user_id", getLoginUserId()).One(&address)
	} else {
		err = o.QueryTable(addresstable).Filter("id", intaddressid).Filter("user_id", getLoginUserId()).One(&address)
	}

	if err != orm.ErrNoRows {
		address.ProviceName = models.GetRegionName(address.ProvinceId)
		address.CityName = models.GetRegionName(address.CityId)
		address.DistrictName = models.GetRegionName(address.DistrictId)
		address.FullRegion = address.ProviceName + address.CityName + address.DistrictName
	}

	var freightPrice float64 = 0.0
	cartData := getCart()
	var checkedgoodslist []models.NideshopCart
	for _, val := range cartData.CartList {
		if val.Checked == 1 {
			checkedgoodslist = append(checkedgoodslist, val)
		}
	}

	var couponPrice float64 = 0.0

	goodstotalprice := cartData.CartTotal.CheckedGoodsAmount
	ordertotalprice := cartData.CartTotal.CheckedGoodsAmount + freightPrice - couponPrice
	actualPrice := ordertotalprice - 0

	utils.ReturnHTTPSuccess(&this.Controller, CheckoutRtnJson{
		Address:      address,
		FreightPrice: freightPrice,
		// checkedCoupon: {},
		// couponList: couponList,
		CouponPrice:      couponPrice,
		CheckedGoodsList: checkedgoodslist,
		GoodsTotalPrice:  goodstotalprice,
		OrderTotalPrice:  ordertotalprice,
		ActualPrice:      actualPrice,
	})
	this.ServeJSON()
}
