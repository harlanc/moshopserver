package models

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"

	"github.com/harlanc/moshopserver/services"
	"github.com/harlanc/moshopserver/utils"
)

/**
 * 获取最新的订单物流信息
 * @param orderId
 * @returns {Promise.<*>}
 */
func GetLatestOrderExpress(orderid int) services.ExpressRtnInfo {
	var expressinfo services.ExpressRtnInfo = services.ExpressRtnInfo{
		ShipperCode:  "",
		ShipperName:  "",
		LogisticCode: "",
		IsFinish:     0,
		RequestTime:  0,
		Traces:       make([]services.Traces, 0),
	}

	o := orm.NewOrm()
	orderexpresstable := new(NideshopOrderExpress)
	var orderexpress NideshopOrderExpress
	err := o.QueryTable(orderexpresstable).Filter("order_id", orderid).One(&orderexpress)
	if err == orm.ErrNoRows {
		return expressinfo
	}

	if orderexpress.ShipperCode == "" || orderexpress.LogisticCode == "" {
		return expressinfo
	}
	expressinfo.ShipperCode = orderexpress.ShipperCode
	expressinfo.ShipperName = orderexpress.ShipperName
	expressinfo.LogisticCode = orderexpress.LogisticCode
	expressinfo.IsFinish = orderexpress.IsFinish
	expressinfo.RequestTime = utils.GetTimestamp()
	var res []services.Traces
	json.Unmarshal([]byte(orderexpress.Traces), &res)
	expressinfo.Traces = res

	if orderexpress.IsFinish == 1 {
		return expressinfo
	}

	expressserviceres := services.QueryExpress(expressinfo.ShipperCode, expressinfo.LogisticCode, "")
	nowtime := utils.GetTimestamp()

	if expressserviceres.Success {
		expressinfo.Traces = expressserviceres.Traces
		expressinfo.IsFinish = expressserviceres.IsFinish
		expressinfo.RequestTime = nowtime
	}

	traces, _ := json.Marshal(expressinfo.Traces)

	o.QueryTable(orderexpresstable).Filter("id", orderexpress.Id).Update(orm.Params{
		"request_time":  nowtime,
		"update_time":   nowtime,
		"traces":        traces,
		"is_finish":     expressinfo.IsFinish,
		"request_count": orm.ColValue(orm.ColAdd, 1)})

	return expressinfo
}
