package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
)

type SearchController struct {
	beego.Controller
}

type SearchIndexRtnJson struct {
	DefaultKeyword     models.NideshopKeywords
	HistoryKeyworkList []orm.Params
	HotKeywordList     []orm.Params
}

func (this *SearchController) Search_Index() {
	o := orm.NewOrm()
	keywordstable := new(models.NideshopKeywords)
	var keyword models.NideshopKeywords

	o.QueryTable(keywordstable).Filter("is_default", 1).Limit(1).One(&keyword)

	var hotkeywords []orm.Params
	o.QueryTable(keywordstable).Distinct().Limit(10).Values(&hotkeywords, "keyword", "is_hot")

	searchhistorytable := new(models.NideshopSearchHistory)
	var historykeywords []orm.Params
	o.QueryTable(searchhistorytable).Filter("user_id", getLoginUserId()).Distinct().Limit(10).Values(&historykeywords, "keyword")

	data, err := json.Marshal(SearchIndexRtnJson{
		DefaultKeyword:     keyword,
		HistoryKeyworkList: historykeywords,
		HotKeywordList:     hotkeywords,
	})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}

	this.ServeJSON()

}

func (this *SearchController) Search_Helper() {

	keyword := this.GetString("keyword")

	o := orm.NewOrm()
	keywordstable := new(models.NideshopKeywords)

	var reskeywords []orm.Params
	o.QueryTable(keywordstable).Filter("keyword__icontains", keyword).Distinct().Limit(10).Values(&reskeywords, "keyword")

	data, err := json.Marshal(reskeywords)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}

	this.ServeJSON()

}

func (this *SearchController) Search_Clearhistory() {

	o := orm.NewOrm()
	keywordstable := new(models.NideshopKeywords)

	o.QueryTable(keywordstable).Filter("user_id", getLoginUserId()).Delete()

	return

}
