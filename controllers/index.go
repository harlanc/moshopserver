package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {

	o := orm.NewOrm()

	var banners []*models.NideshopAd
	ad := new(models.NideshopAd)
	o.QueryTable(ad).Filter("id", 1).All(&banners)

	var channels []*models.NideshopChannel
	channel := new(models.NideshopChannel)
	o.QueryTable(channel).OrderBy("sort_order").All(&channels)

	var newgoogs []orm.Params
	goods := new(models.NideshopGoods)
	o.QueryTable(goods).Filter("is_new", 1).Limit(4).Values(&newgoogs, "id", "name", "list_pic_url", "retail_price")

	var hotgoogs []orm.Params
	o.QueryTable(goods).Filter("is_hot", 1).Limit(3).Values(&hotgoogs, "id", "name", "list_pic_url", "retail_price", "goods_brief")

	var brandList []models.NideshopBrand
	brand := new(models.NideshopBrand)
	o.QueryTable(brand).Filter("is_new", 1).OrderBy("new_sort_order").Limit(4).All(brandList)

	var topicList []models.NideshopTopic
	topic := new(models.NideshopTopic)
	o.QueryTable(topic).Limit(3).All(topicList)

	var categoryList []models.NideshopCategory
	category := new(models.NideshopCategory)
	o.QueryTable(category).Filter("parent_id", 0).Exclude("name", "推荐").All(categoryList)

	type newCategoryList struct {
		id        int
		name      string
		goodsList []orm.Params
	}

	var newList []newCategoryList

	for _, categoryItem := range categoryList {
		var ids []orm.Params
		o.QueryTable(category).Filter("parent_id", categoryItem.Id).Values(&ids, "id")

		var categorygoods []orm.Params
		o.QueryTable(goods).Filter("category_id__in", ids).Limit(7).Values(&categorygoods, "id", "name", "list_pic_url", "retail_price")

		newList = append(newList, newCategoryList{categoryItem.Id, categoryItem.Name, categorygoods})

	}

	// qb, _ := orm.NewQueryBuilder("mysql")
	// o := orm.NewOrm()
	// var sql string

	// var banners []models.NideshopAd
	// qb.Select("*").
	// 	From("nideshop_ad").
	// 	Where("id == 1")
	// sql = qb.String()

	// o.Raw(sql, 20).QueryRows(&banners)
	// this.Data["banner"] = banners

	// var channels []models.NideshopChannel
	// qb.Select("*").From("nideshop_channel").OrderBy("asc")
	// sql = qb.String()
	// o.Raw(sql, 20).QueryRows(&channels)
	// this.Data["channel"] = channels

	// var newgoogs []models.NideshopGoods
	// qb.Select("id,name,list_pic_url,retail_price").From("nideshop_goods").Where("is_new == 1").Limit(4)
	// sql = qb.String()
	// o.Raw(sql, 20).QueryRows(&newgoogs)
	// this.Data["newGoodsList"] = newgoogs

	// var hotgoogs []models.NideshopGoods
	// qb.Select("id,name,list_pic_url,retail_price,goods_brief").From("nideshop_goods").Where("is_hot == 1").Limit(3)
	// sql = qb.String()
	// o.Raw(sql, 20).QueryRows(&hotgoogs)
	// this.Data["hotGoodsList"] = hotgoogs

	// var brandList []models.NideshopBrand
	// qb.Select("*").From("nideshop_brand").Where("is_new == 1").OrderBy("asc").Limit(3)
	// sql = qb.String()
	// o.Raw(sql, 20).QueryRows(&brandList)
	// this.Data["brandList"] = brandList

	// var topicList []models.NideshopTopic
	// qb.Select("*").From("nideshop_topic").Limit(3)
	// sql = qb.String()
	// o.Raw(sql, 20).QueryRows(&topicList)
	// this.Data["topicList"] = topicList

	// type newCategoryList struct {
	// 	id        int
	// 	name      string
	// 	goodsList models.NideshopGoods
	// }

	// var caregoryList []models.NideshopCategory
	// var newCategoryList []newCategoryList
	// qb.Select("*").From("nideshop_category").Where("parent_id == 0 && name <> '推荐'")
	// sql = qb.String()
	// o.Raw(sql, 20).QueryRows(&caregoryList)
	// for _, categoryItem := range caregoryList {
	// 	var ids []int
	// 	qb.Select("id").From("nideshop_category").Where("parent_id == " + strconv.Itoa(categoryItem.Id)).Limit(100)
	// 	sql = qb.String()
	// 	o.Raw(sql, 20).QueryRows(&ids)

	// }

	// o := orm.NewOrm()
	// / banner := models.NideshopAd{AdPositionId: 1}

	// err := o.Read(&banner)

	// if err == orm.ErrNoRows {
	// 	fmt.Println("Can not find.")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("Can not find main key.")
	// } else {
	// 	fmt.Println(banner.Id, banner.Name)
	// }

}
