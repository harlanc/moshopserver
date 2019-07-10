package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/harlanc/moshopserver/models"
	"github.com/harlanc/moshopserver/services"
	"github.com/harlanc/moshopserver/utils"
)

type OrderController struct {
	beego.Controller
}

//It may need to be refactored.
func GetOrderPageData(rawData []models.NideshopOrder, page int, size int) utils.PageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []models.NideshopOrder

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return utils.PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalpages, Data: pagedata}
}

type OrderListRtnJson struct {
	models.NideshopOrder
	GoodsList       []models.NideshopOrderGoods `json:"goodList"`
	GoodsCount      int                         `json:"goodsCount"`
	OrderStatusText string                      `json:"order_status_text"`
	HandOption      models.OrderHandleOption    `json:"handleOption"`
}

func (this *OrderController) Order_List() {

	o := orm.NewOrm()
	ordertable := new(models.NideshopOrder)
	var orders []models.NideshopOrder
	o.QueryTable(ordertable).Filter("user_id", getLoginUserId()).All(&orders)

	firstpagedorders := GetOrderPageData(orders, 1, 10)

	var rtnorderlist []OrderListRtnJson
	ordergoodstable := new(models.NideshopOrderGoods)
	var ordergoods []models.NideshopOrderGoods
	qsordergoods := o.QueryTable(ordergoodstable)
	for _, val := range firstpagedorders.Data.([]models.NideshopOrder) {
		qsordergoods.Filter("order_id", val.Id).All(&ordergoods)
		var goodscount int
		for _, val := range ordergoods {
			goodscount += val.Number
		}
		orderstatustext := models.GetOrderStatusText(val.Id)
		orderhandoption := models.GetOrderHandleOption(val.Id)
		var orderlistrtn OrderListRtnJson = OrderListRtnJson{val, ordergoods, goodscount, orderstatustext, orderhandoption}

		rtnorderlist = append(rtnorderlist, orderlistrtn)

	}

	firstpagedorders.Data = rtnorderlist

	utils.ReturnHTTPSuccess(&this.Controller, firstpagedorders)
	this.ServeJSON()
}

type OrderInfo struct {
	models.NideshopOrder
	ProvinceName        string                  `json:"province_name"`
	CityName            string                  `json:"city_name"`
	DistrictName        string                  `json:"district_name"`
	FullRegion          string                  `json:"full_region"`
	Express             services.ExpressRtnInfo `json:"express"`
	OrderStatusText     string                  `json:"order_status_text"`
	FormatAddTime       string                  `json:"add_time"`
	FormatFinalPlayTime string                  `json:"final_pay_time"`
}

type OrderDetailRtnJson struct {
	OrderInfo    OrderInfo                   `json:"orderInfo"`
	OrderGoods   []models.NideshopOrderGoods `json:"orderGoods"`
	HandleOption models.OrderHandleOption    `json:"handleOption"`
}

func (this *OrderController) Order_Detail() {

	orderId := this.GetString("orderId")
	intorderId := utils.String2Int(orderId)

	o := orm.NewOrm()
	ordertable := new(models.NideshopOrder)
	var order models.NideshopOrder
	err := o.QueryTable(ordertable).Filter("id", intorderId).Filter("user_id", getLoginUserId()).One(&order)

	if err == orm.ErrNoRows {
		this.Abort("订单不存在")
	}

	var orderinfo OrderInfo = OrderInfo{NideshopOrder: order}
	orderinfo.ProvinceName = models.GetRegionName(order.Province)
	orderinfo.CityName = models.GetRegionName(order.City)
	orderinfo.DistrictName = models.GetRegionName(order.District)
	orderinfo.FullRegion = orderinfo.ProvinceName + orderinfo.CityName + orderinfo.DistrictName

	lastestexpressinfo := models.GetLatestOrderExpress(intorderId)
	orderinfo.Express = lastestexpressinfo

	ordergoodstable := new(models.NideshopOrderGoods)
	var ordergoods []models.NideshopOrderGoods

	o.QueryTable(ordergoodstable).Filter("order_id", intorderId).All(&ordergoods)

	orderinfo.OrderStatusText = models.GetOrderStatusText(intorderId)
	orderinfo.FormatAddTime = utils.FormatTimestamp(orderinfo.AddTime, "2006-01-02 03:04:05 PM")
	orderinfo.FormatFinalPlayTime = utils.FormatTimestamp(1234, "04:05")

	if orderinfo.OrderStatus == 0 {
		//todo 订单超时逻辑
	}

	handleoption := models.GetOrderHandleOption(intorderId)
	utils.ReturnHTTPSuccess(&this.Controller, OrderDetailRtnJson{
		OrderInfo:    orderinfo,
		OrderGoods:   ordergoods,
		HandleOption: handleoption,
	})
	this.ServeJSON()
}

func (this *OrderController) Order_Submit() {
	addressId := this.GetString("addressId")
	//couponId := this.GetString("couponId")
	postscript := this.GetString("postscript")
	intaddressId := utils.String2Int(addressId)
	//intcouponId := utils.String2Int(couponId)

	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)
	var address models.NideshopAddress

	err := o.QueryTable(addresstable).Filter("id", intaddressId).One(&address)
	if err == orm.ErrNoRows {
		this.Abort("请选择收获地址")
	}

	carttable := new(models.NideshopCart)
	var carts []models.NideshopCart
	_, err = o.QueryTable(carttable).Filter("user_id", getLoginUserId()).Filter("session_id", 1).Filter("checked", 1).All(&carts)
	if err == orm.ErrNoRows {
		this.Abort("请选择商品")
	}

	var freightPrice float64 = 0
	var goodstotalprice float64 = 0

	for _, val := range carts {
		goodstotalprice += float64(val.Number) * val.RetailPrice
	}

	var couponprice float64
	ordertotalprice := goodstotalprice + freightPrice - couponprice
	actualprice := ordertotalprice - 0
	currenttime := utils.GetTimestamp()

	orderinfo := models.NideshopOrder{
		OrderSn:      models.GenerateOrderNumber(),
		UserId:       getLoginUserId(),
		Consignee:    address.Name,
		Mobile:       address.Mobile,
		Province:     address.ProvinceId,
		City:         address.CityId,
		District:     address.DistrictId,
		Address:      address.Address,
		FreightPrice: 0,
		Postscript:   postscript,
		CouponId:     0,
		CouponPrice:  couponprice,
		AddTime:      currenttime,
		GoodsPrice:   goodstotalprice,
		OrderPrice:   ordertotalprice,
		ActualPrice:  actualprice,
	}

	orderid, err := o.Insert(&orderinfo)
	if err != nil {
		this.Abort("订单提交失败")
	}
	orderinfo.Id = int(orderid)

	for _, item := range carts {
		ordergood := models.NideshopOrderGoods{
			OrderId:                   int(orderid),
			GoodsId:                   item.GoodsId,
			GoodsSn:                   item.GoodsSn,
			ProductId:                 item.ProductId,
			GoodsName:                 item.GoodsName,
			ListPicUrl:                item.ListPicUrl,
			MarketPrice:               item.MarketPrice,
			RetailPrice:               item.RetailPrice,
			Number:                    item.Number,
			GoodsSpecifitionNameValue: item.GoodsSpecifitionNameValue,
			GoodsSpecifitionIds:       item.GoodsSpecifitionIds,
		}
		o.Insert(&ordergood)
	}
	models.ClearBuyGoods(getLoginUserId())

	utils.ReturnHTTPSuccess(&this.Controller, orderinfo)
	this.ServeJSON()

}

func (this *OrderController) Order_Express() {
	orderId := this.GetString("orderId")
	intorderId := utils.String2Int(orderId)

	if orderId == "" {
		this.Abort("订单不存在")
	}

	latestexpressinfo := models.GetLatestOrderExpress(intorderId)

	utils.ReturnHTTPSuccess(&this.Controller, latestexpressinfo)
	this.ServeJSON()
}
