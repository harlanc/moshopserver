package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/utils"
)

func GetChildCategoryId(categoryid int) []int64 {

	o := orm.NewOrm()
	categorytable := new(NideshopCategory)
	var childids []orm.Params
	o.QueryTable(categorytable).Filter("parent_id", categoryid).Limit(10000).Values(&childids, "id")
	childintids := utils.ExactMapValues2Int64Array(childids, "Id")
	return childintids
}

func GetCategoryWhereIn(categoryid int) []int64 {

	childintids := GetChildCategoryId(categoryid)
	childintids = append(childintids, int64(categoryid))
	return childintids
}
