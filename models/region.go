package models

import (
	"github.com/astaxie/beego/orm"
)

func GetRegionName(regionid int) string {

	o := orm.NewOrm()
	regiontable := new(NideshopRegion)
	var region NideshopRegion
	o.QueryTable(regiontable).Filter("id", regionid).One(&region)

	return region.Name

}
