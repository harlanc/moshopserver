package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/harlanc/moshopserver/models"
	"github.com/harlanc/moshopserver/utils"
)

type CollectController struct {
	beego.Controller
}

type CollectListRtnJson struct {
	models.NideshopCollect
	ListPicUrl  string  `json:"list_pic_url"`
	GoodsBrief  string  `json:"goods_brief"`
	RetailPrice float64 `json:"retail_price"`
}

func (this *CollectController) Collect_List() {

	typeId := this.GetString("typeId")

	qb, _ := orm.NewQueryBuilder("mysql")
	var list []CollectListRtnJson

	qb.Select("nc.*", "ng.name", "ng.list_pic_url", "ng.goods_brief", "ng.retail_price").
		From("nideshop_collect nc").
		LeftJoin("nideshop_goods ng").
		On("nc.value_id = ng.id").
		Where("gc.user_id =" + utils.Int2String(getLoginUserId()) + "and gc.type_id = " + typeId)

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&list)

	utils.ReturnHTTPSuccess(&this.Controller, list)
	this.ServeJSON()

}

type AddorDeleteRtnJson struct {
	HandleType string
}

func (this *CollectController) Collect_AddorDelete() {
	typeId := this.GetString("typeId")
	valueId := this.GetString("valueId")

	inttypeId := utils.String2Int(typeId)
	intvalueId := utils.String2Int(valueId)

	o := orm.NewOrm()
	collecttable := new(models.NideshopCollect)
	qs := o.QueryTable(collecttable)

	var collect models.NideshopCollect
	var rvjson AddorDeleteRtnJson

	err := qs.Filter("type_id", inttypeId).Filter("value_id", intvalueId).Filter("user_id", getLoginUserId()).One(&collect)

	if err == orm.ErrNoRows {
		_, err = o.Insert(models.NideshopCollect{
			TypeId:  inttypeId,
			ValueId: intvalueId,
			UserId:  getLoginUserId(),
			AddTime: utils.GetTimestamp(),
		})
		rvjson = AddorDeleteRtnJson{HandleType: "add"}

	} else {
		_, err = qs.Filter("id", collect.Id).Delete()
		rvjson = AddorDeleteRtnJson{HandleType: "delete"}
	}

	if err != nil {
		this.Abort(err.Error())
	}

	utils.ReturnHTTPSuccess(&this.Controller, rvjson)

	this.ServeJSON()

}
