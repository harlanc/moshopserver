package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type RegionController struct {
	beego.Controller
}

func (this *RegionController) Region_Info() {

	regionId := this.GetString("regionId")
	intregionid := utils.String2Int(regionId)

	o := orm.NewOrm()
	regiontable := new(models.NideshopRegion)
	var region models.NideshopRegion
	o.QueryTable(regiontable).Filter("id", intregionid).One(&region)

	data, err := json.Marshal(region)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}

	this.ServeJSON()

}

func (this *RegionController) Region_List() {

	parentId := this.GetString("parentId")
	intparentid := utils.String2Int(parentId)

	o := orm.NewOrm()
	regiontable := new(models.NideshopRegion)
	var regions []models.NideshopRegion
	o.QueryTable(regiontable).Filter("parent_id", intparentid).All(&regions)

	data, err := json.Marshal(regions)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}

	this.ServeJSON()
}
