package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type TopicController struct {
	beego.Controller
}

func updateJsonKeysTopic(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "PriceInfo":
				delete(val, k)
				val["price_info"] = v
			case "ScenePicUrl":
				delete(val, k)
				val["scene_pic_url"] = v
			case "Subtitle":
				delete(val, k)
				val["subtitle"] = v
			case "Title":
				delete(val, k)
				val["title"] = v
			}
		}
	}
}

func (this *TopicController) Topic_List() {

	page := this.GetString("page")
	size := this.GetString("size")

	var intpage, intsize int
	if page == "" {
		intpage = 1
	} else {
		intpage = utils.String2Int(page)
	}

	if size == "" {
		intsize = 10
	} else {
		intsize = utils.String2Int(size)
	}

	o := orm.NewOrm()
	topictable := new(models.NideshopTopic)
	var topics []orm.Params
	o.QueryTable(topictable).Values(&topics, "id", "title", "price_info", "scene_pic_url", "subtitle")
	updateJsonKeysTopic(topics)

	rtndata := utils.GetPageData(topics, intpage, intsize)

	utils.ReturnHTTPSuccess(&this.Controller, rtndata)
	this.ServeJSON()

}

func (this *TopicController) Topic_Detail() {
	id := this.GetString("id")
	intid := utils.String2Int(id)

	o := orm.NewOrm()
	topictable := new(models.NideshopTopic)
	var topic models.NideshopTopic

	o.QueryTable(topictable).Filter("id", intid).One(&topic)

	utils.ReturnHTTPSuccess(&this.Controller, topic)
	this.ServeJSON()

}

func (this *TopicController) Topic_Related() {

	o := orm.NewOrm()
	topictable := new(models.NideshopTopic)
	var topics []orm.Params
	o.QueryTable(topictable).Limit(4).Values(&topics, "id", "title", "price_info", "scene_pic_url", "subtitle")

	utils.ReturnHTTPSuccess(&this.Controller, topics)
	this.ServeJSON()
}
