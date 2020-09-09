package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"moshopserver/models"
	"moshopserver/utils"
)

type CatalogController struct {
	beego.Controller
}

type CurCategory struct {
	models.NideshopCategory
	SubCategoryList []models.NideshopCategory `json:"subCategoryList"`
}

type CateLogIndexRtnJson struct {
	CategoryList    []models.NideshopCategory `json:"categoryList"`
	CurrentCategory CurCategory               `json:"currentCategory"`
}

func (this *CatalogController) Catalog_Index() {

	categoryId := this.GetString("id")

	o := orm.NewOrm()

	var categories []models.NideshopCategory
	categorytable := new(models.NideshopCategory)
	o.QueryTable(categorytable).Filter("parent_id", 0).Limit(10).All(&categories)

	var currentCategory *models.NideshopCategory = nil

	if categoryId != "" {
		o.QueryTable(categorytable).Filter("id", categoryId).One(currentCategory)
	}

	if currentCategory == nil {
		currentCategory = &categories[0]
	}

	curCategory := new(CurCategory)

	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []models.NideshopCategory
		o.QueryTable(categorytable).Filter("parent_id", currentCategory.Id).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.NideshopCategory = *currentCategory
	}

	utils.ReturnHTTPSuccess(&this.Controller, CateLogIndexRtnJson{categories, *curCategory})
	this.ServeJSON()

}

type CateLogCurRtnJson struct {
	CurrentCategory CurCategory `json:"currentCategory"`
}

func (this *CatalogController) Catalog_Current() {

	categoryId := this.GetString("id")

	o := orm.NewOrm()
	categorytable := new(models.NideshopCategory)
	currentCategory := new(models.NideshopCategory)
	if categoryId != "" {
		o.QueryTable(categorytable).Filter("id", categoryId).One(currentCategory)
	}

	curCategory := new(CurCategory)
	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []models.NideshopCategory
		o.QueryTable(categorytable).Filter("parent_id", currentCategory.Id).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.NideshopCategory = *currentCategory
	}

	utils.ReturnHTTPSuccess(&this.Controller, CateLogCurRtnJson{*curCategory})
	this.ServeJSON()

}
