package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
)

type IndexController struct {
	beego.Controller
}

type newCategoryList struct {
	id        int
	name      string
	goodsList []orm.Params
}

type RtnJson struct {
	Banners      []models.NideshopAd
	Channels     []models.NideshopChannel
	Newgoods     []orm.Params
	Hotgoods     []orm.Params
	BrandList    []models.NideshopBrand
	TopicList    []models.NideshopTopic
	CategoryList []newCategoryList
}

func (this *IndexController) Get() {

	o := orm.NewOrm()

	var banners []models.NideshopAd
	ad := new(models.NideshopAd)
	o.QueryTable(ad).Filter("id", 1).All(&banners)

	var channels []models.NideshopChannel
	channel := new(models.NideshopChannel)
	o.QueryTable(channel).OrderBy("sort_order").All(&channels)

	var newgoods []orm.Params
	goods := new(models.NideshopGoods)
	o.QueryTable(goods).Filter("is_new", 1).Limit(4).Values(&newgoods, "id", "name", "list_pic_url", "retail_price")

	var hotgoods []orm.Params
	o.QueryTable(goods).Filter("is_hot", 1).Limit(3).Values(&hotgoods, "id", "name", "list_pic_url", "retail_price", "goods_brief")

	var brandList []models.NideshopBrand
	brand := new(models.NideshopBrand)
	o.QueryTable(brand).Filter("is_new", 1).OrderBy("new_sort_order").Limit(4).All(&brandList)

	var topicList []models.NideshopTopic
	topic := new(models.NideshopTopic)
	o.QueryTable(topic).Limit(3).All(&topicList)

	var categoryList []models.NideshopCategory
	category := new(models.NideshopCategory)
	o.QueryTable(category).Filter("parent_id", 0).Exclude("name", "推荐").All(&categoryList)

	var newList []newCategoryList

	for _, categoryItem := range categoryList {
		var ids []orm.Params
		o.QueryTable(category).Filter("parent_id", categoryItem.Id).Values(&ids, "id")

		var categorygoods []orm.Params
		o.QueryTable(goods).Filter("category_id__in", ids).Limit(7).Values(&categorygoods, "id", "name", "list_pic_url", "retail_price")

		newList = append(newList, newCategoryList{categoryItem.Id, categoryItem.Name, categorygoods})
	}

	data, err := json.Marshal(RtnJson{banners, channels, newgoods, hotgoods, brandList, topicList, newList})
	if err != nil {
		this.Data["json"] = err

	} else {
		this.Data["json"] = string(data)
	}

	this.ServeJSON()

}
