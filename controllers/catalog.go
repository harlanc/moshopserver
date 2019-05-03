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
	SubCategoryList []models.NideshopCategory
}

type CateLogRtnJson struct {
	CategoryList    []models.NideshopCategory
	CurrentCategory CurCategory
}

func (this *CatalogController) Catalog_Index() {
	categoryId := this.GetString("id")

	o := orm.NewOrm()

	var categories []models.NideshopCategory
	category := new(models.NideshopCategory)
	o.QueryTable(category).Filter("parent_id", 0).Limit(10).All(&categories)

	var currentCategory models.NideshopCategory

	if &categoryId != nil {
		o.QueryTable(category).Filter("id", categoryId).One(&currentCategory)
		if &currentCategory == nil {
			currentCategory = categories[0]
		}
	}

	var subCategories []models.NideshopCategory
	o.QueryTable(category).Filter("parent_id", currentCategory.Id).All(&subCategories)

	data, err := json.Marshal(CateLogRtnJson{categories, CurCategory{currentCategory, subCategories}})
	if err != nil {
		this.Data["json"] = err

	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

func (this *CatalogController) Catalog_Current() {
}
