package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type IndexController struct {
	beego.Controller
}

type newCategoryList struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	GoodsList []orm.Params `json:"goodsList"`
}

type IndexRtnJson struct {
	Banners      []models.NideshopAd      `json:"banner"`
	Channels     []models.NideshopChannel `json:"channel"`
	Newgoods     []orm.Params             `json:"newGoodsList"`
	Hotgoods     []orm.Params             `json:"hotGoodsList"`
	BrandList    []models.NideshopBrand   `json:"brandList"`
	TopicList    []models.NideshopTopic   `json:"topicList"`
	CategoryList []newCategoryList        `json:"categoryList"`
}

func updateJsonKeysIndex(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "ListPicUrl":
				delete(val, k)
				val["list_pic_url"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			}
		}
	}
}

func (this *IndexController) Index_Index() {

	o := orm.NewOrm()

	var banners []models.NideshopAd
	ad := new(models.NideshopAd)
	o.QueryTable(ad).Filter("ad_position_id", 1).All(&banners)

	var channels []models.NideshopChannel
	channel := new(models.NideshopChannel)
	o.QueryTable(channel).OrderBy("sort_order").All(&channels)

	var newgoods []orm.Params
	goods := new(models.NideshopGoods)
	o.QueryTable(goods).Filter("is_new", 1).Limit(4).Values(&newgoods, "id", "name", "list_pic_url", "retail_price")
	updateJsonKeysIndex(newgoods)

	var hotgoods []orm.Params
	o.QueryTable(goods).Filter("is_hot", 1).Limit(3).Values(&hotgoods, "id", "name", "list_pic_url", "retail_price", "goods_brief")
	updateJsonKeysIndex(hotgoods)

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
		var mapids []orm.Params
		o.QueryTable(category).Filter("parent_id", categoryItem.Id).Values(&mapids, "id")

		// var valIds []int64
		// for _, value := range mapids {
		// 	valIds = append(valIds, value["Id"].(int64))
		// }

		valIds := utils.ExactMapValues2Int64Array(mapids, "Id")

		var categorygoods []orm.Params
		o.QueryTable(goods).Filter("category_id__in", valIds).Limit(7).Values(&categorygoods, "id", "name", "list_pic_url", "retail_price")
		updateJsonKeysIndex(categorygoods)
		newList = append(newList, newCategoryList{categoryItem.Id, categoryItem.Name, categorygoods})
	}

	utils.ReturnHTTPSuccess(&this.Controller, IndexRtnJson{banners, channels, newgoods, hotgoods, brandList, topicList, newList})

	this.ServeJSON()

}
