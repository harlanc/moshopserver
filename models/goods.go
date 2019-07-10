package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/harlanc/moshopserver/utils"
)

type SpecificationData struct {
	NideshopGoodsSpecification
	Name string
}

type SpecificationItem struct {
	Specification_id int
	Name             string
	List             []SpecificationData
}

func GetProductList(goodsId int) []NideshopProduct {
	o := orm.NewOrm()

	var products []NideshopProduct
	product := new(NideshopProduct)

	o.QueryTable(product).Filter("goods_id", goodsId).All(&products)

	return products

}

func GetSpecificationList(goodsId int) []SpecificationItem {

	qb, _ := orm.NewQueryBuilder("mysql")

	var specifications []SpecificationData

	qb.Select("gs.*", "s.name").
		From("nideshop_goods_specification gs").
		InnerJoin("nideshop_specification s").On("gs.specification_id = s.id").
		Where("gs.specification_id =" + utils.Int2String(goodsId))

	sql := qb.String()

	o := orm.NewOrm()
	o.Raw(sql, 20).QueryRows(&specifications)

	var label map[int]int
	specificationList := make([]SpecificationItem, 0)
	var idx int = 0

	for _, item := range specifications {

		if v, ok := label[item.Id]; ok {
			specificationList[v].List = append(specificationList[v].List, item)
		} else {

			specificationList = append(specificationList, SpecificationItem{item.Id, item.Name, []SpecificationData{item}})
			label[item.Id] = idx
			idx += 1
		}
	}

	return specificationList

}
