package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
)

type CatalogController struct {
	beego.Controller
}

type CurCategory struct {
	models.NideshopCategory
	SubCategoryList []models.NideshopCategory `json:"subCategoryList"`
}

type CateLogRtnJson struct {
	CategoryList    []models.NideshopCategory `json:"categoryList"`
	CurrentCategory CurCategory               `json:"currentCategory"`
}

func getCurCategory(categoryId string) CurCategory {

	o := orm.NewOrm()

	var currentCategory models.NideshopCategory
	category := new(models.NideshopCategory)
	if &categoryId != nil {
		o.QueryTable(category).Filter("id", categoryId).One(&currentCategory)
	}

	var subCategories []models.NideshopCategory
	o.QueryTable(category).Filter("parent_id", currentCategory.Id).All(&subCategories)
	return CurCategory{currentCategory, subCategories}
}

func (this *CatalogController) Catalog_Index() {
	categoryId := this.GetString("id")

	o := orm.NewOrm()

	var categories []models.NideshopCategory
	category := new(models.NideshopCategory)
	o.QueryTable(category).Filter("parent_id", 0).Limit(10).All(&categories)

	data, err := json.Marshal(CateLogRtnJson{categories, getCurCategory(categoryId)})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

func (this *CatalogController) Catalog_Current() {

	categoryId := this.GetString("id")

	data, err := json.Marshal(CateLogRtnJson{CurrentCategory: getCurCategory(categoryId)})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}
