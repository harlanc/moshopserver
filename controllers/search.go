package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type SearchController struct {
	beego.Controller
}

type SearchIndexRtnJson struct {
	DefaultKeyword     models.NideshopKeywords `json:"defaultKeyword"`
	HistoryKeyworkList []string                `json:"historyKeywordList"`
	HotKeywordList     []orm.Params            `json:"hotKeywordList"`
}

func updateJsonKeysSearch(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Keyword":
				delete(val, k)
				val["keyword"] = v
			case "IsHot":
				delete(val, k)
				val["is_hot"] = v
			}
		}
	}
}

func (this *SearchController) Search_Index() {
	o := orm.NewOrm()
	keywordstable := new(models.NideshopKeywords)
	var keyword models.NideshopKeywords

	o.QueryTable(keywordstable).Filter("is_default", 1).Limit(1).One(&keyword)

	var hotkeywords []orm.Params
	o.QueryTable(keywordstable).Distinct().Limit(10).Values(&hotkeywords, "keyword", "is_hot")
	updateJsonKeysSearch(hotkeywords)

	searchhistorytable := new(models.NideshopSearchHistory)
	var historykeywords []orm.Params
	o.QueryTable(searchhistorytable).Filter("user_id", getLoginUserId()).Distinct().Limit(10).Values(&historykeywords, "keyword")
	arraykeywords := utils.ExactMapValues2StringArray(historykeywords, "Keyword")

	utils.ReturnHTTPSuccess(&this.Controller, SearchIndexRtnJson{
		DefaultKeyword:     keyword,
		HistoryKeyworkList: arraykeywords,
		HotKeywordList:     hotkeywords,
	})

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
