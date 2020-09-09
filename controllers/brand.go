package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"moshopserver/models"
	"moshopserver/utils"
)

type BrandController struct {
	beego.Controller
}

func (this *BrandController) Brand_List() {

	page := this.GetString("page")
	size := this.GetString("size")

	var intsize int = 10
	if size != "" {
		intsize = utils.String2Int(size)
	}

	var intpage int = 1
	if page != "" {
		intpage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	brandtable := new(models.NideshopBrand)
	var brands []orm.Params
	o.QueryTable(brandtable).Values(&brands, "id", "name", "floor_price", "app_list_pic_url")

	pagedata := utils.GetPageData(brands, intpage, intsize)

	utils.ReturnHTTPSuccess(&this.Controller, pagedata)
	this.ServeJSON()

}

type BrandDetailRtnJson struct {
	Data models.NideshopBrand
}

func (this *BrandController) Brand_Detail() {
	id := this.GetString("id")
	intid := utils.String2Int(id)

	o := orm.NewOrm()
	brandtable := new(models.NideshopBrand)
	var brand models.NideshopBrand

	o.QueryTable(brandtable).Filter("id", intid).One(&brand)

	utils.ReturnHTTPSuccess(&this.Controller, BrandDetailRtnJson{brand})
	this.ServeJSON()
}
