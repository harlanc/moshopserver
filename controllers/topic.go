package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type TopicController struct {
	beego.Controller
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

	rtndata := utils.GetPageData(topics, intpage, intsize)

	data, err := json.Marshal(rtndata)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

func (this *TopicController) Topic_Detail() {
	id := this.GetString("id")
	intid := utils.String2Int(id)

	o := orm.NewOrm()
	topictable := new(models.NideshopTopic)
	var topic models.NideshopTopic

	o.QueryTable(topictable).Filter("id", intid).One(&topic)

	data, err := json.Marshal(topic)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

func (this *TopicController) Topic_Related() {

	o := orm.NewOrm()
	topictable := new(models.NideshopTopic)
	var topics []orm.Params
	o.QueryTable(topictable).Limit(4).Values(&topics, "id", "title", "price_info", "scene_pic_url", "subtitle")

	data, err := json.Marshal(topics)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}
