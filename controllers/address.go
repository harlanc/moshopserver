package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type AddressController struct {
	beego.Controller
}

type AddressIndexRtnJson struct {
	Address      models.NideshopAddress
	ProviceName  string
	CityName     string
	DistrictName string
	FullRegion   string
}

func (this *AddressController) Address_Index() {

	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)
	var addresses []models.NideshopAddress

	o.QueryTable(addresstable).Filter("user_id", getLoginUserId()).All(&addresses)

	rtnaddress := make([]AddressIndexRtnJson, 0)

	for _, val := range addresses {

		provicename := models.GetRegionName(val.ProvinceId)
		cityname := models.GetRegionName(val.CityId)
		distinctname := models.GetRegionName(val.DistrictId)
		rtnaddress = append(rtnaddress, AddressIndexRtnJson{
			Address:      val,
			ProviceName:  provicename,
			CityName:     cityname,
			DistrictName: distinctname,
			FullRegion:   provicename + cityname + distinctname,
		})

	}

	data, err := json.Marshal(rtnaddress)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}
func (this *AddressController) Address_Detail() {
	id := this.GetString("id")

	intid := utils.String2Int(id)

	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)
	var address models.NideshopAddress

	o.QueryTable(addresstable).Filter("id", intid).Filter("user_id", getLoginUserId()).One(&address)

	provicename := models.GetRegionName(address.ProvinceId)
	cityname := models.GetRegionName(address.CityId)
	distinctname := models.GetRegionName(address.DistrictId)

	data, err := json.Marshal(AddressIndexRtnJson{
		Address:      address,
		ProviceName:  provicename,
		CityName:     cityname,
		DistrictName: distinctname,
		FullRegion:   provicename + cityname + distinctname,
	})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

func (this *AddressController) Address_Save() {

	addressid := this.GetString("id")
	name := this.GetString("name")
	mobile := this.GetString("mobile")
	provinceid := this.GetString("province_id")
	cityid := this.GetString("city_id")
	distinctid := this.GetString("distinct_id")
	address := this.GetString("address")
	userid := getLoginUserId()
	var isdefault int
	if this.GetString("id_default") == "true" {
		isdefault = 1
	} else {
		isdefault = 0
	}

	intcityid := utils.String2Int(cityid)
	intprovinceid := utils.String2Int(provinceid)
	intdistinctid := utils.String2Int(distinctid)

	// type NideshopAddress struct {
	// 	Address    string `orm:"not null default '' VARCHAR(120)"`
	// 	CityId     int    `orm:"not null default 0 SMALLINT(5)"`
	// 	CountryId  int    `orm:"not null default 0 SMALLINT(5)"`
	// 	DistrictId int    `orm:"not null default 0 SMALLINT(5)"`
	// 	Id         int    `orm:"not null pk autoincr MEDIUMINT(8)"`
	// 	IsDefault  int    `orm:"not null default 0 TINYINT(1)"`
	// 	Mobile     string `orm:"not null default '' VARCHAR(60)"`
	// 	Name       string `orm:"not null default '' VARCHAR(50)"`
	// 	ProvinceId int    `orm:"not null default 0 SMALLINT(5)"`
	// 	UserId     int    `orm:"not null default 0 index MEDIUMINT(8)"`
	// }

	addressdata := models.NideshopAddress{
		Address:    address,
		CityId:     intcityid,
		DistrictId: intdistinctid,
		ProvinceId: intprovinceid,
		Name:       name,
		Mobile:     mobile,
		UserId:     userid,
		IsDefault:  isdefault,
	}
	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)

	var intid int64
	if addressid == "" {
		id, _ := o.Insert(addressdata)
		intid = id
	} else {
		o.QueryTable(addresstable).Filter("id", intid).Filter("user_id", userid).Update(orm.Params{
			"is_default": 0,
		})
	}

	if isdefault == 1 {
		_, err := o.Raw("UPDATE nideshop_address SET is_default = 0 where id <> ? and user_id =", intid, userid).Exec()
		if err == nil {
			//res.RowsAffected()
			//fmt.Println("mysql row affected nums: ", num)
		}
	}
	var addressinfo models.NideshopAddress
	o.QueryTable(addresstable).Filter("id", intid).One(&addressinfo)

	data, err := json.Marshal(addressinfo)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

func (this *AddressController) Address_Delete() {

	addressid := this.GetString("id")
	intaddressid := utils.String2Int(addressid)

	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)
	o.QueryTable(addresstable).Filter("id", intaddressid).Filter("user_id", getLoginUserId()).Delete()

	return

}
