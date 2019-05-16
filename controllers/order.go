package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type OrderController struct {
	beego.Controller
}

//It may need to be refactored.
func GetPageData(rawData []models.NideshopOrder, page int, size int) utils.PageData {

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
	GoodsList       []models.NideshopOrderGoods
	GoodsCount      int
	OrderStatusText string
	HandOption      models.OrderHandleOption
}

func (this *OrderController) Order_List() {

	o := orm.NewOrm()
	ordertable := new(models.NideshopOrder)
	var orders []models.NideshopOrder
	o.QueryTable(ordertable).Filter("user_id", getLoginUserId()).All(&orders)

	firstpagedorders := GetPageData(orders, 1, 10)

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

	data, err := json.Marshal(firstpagedorders)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}

	this.ServeJSON()
}

type OrderInfo struct {
	models.NideshopOrder
	ProviceName  string
	CityName     string
	DistrictName string
	FullRegion   string
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
	orderinfo.ProviceName = models.GetRegionName(order.Province)
	orderinfo.CityName = models.GetRegionName(order.City)
	orderinfo.DistrictName = models.GetRegionName(order.District)
	orderinfo.FullRegion = orderinfo.ProviceName + orderinfo.CityName + orderinfo.DistrictName

}
