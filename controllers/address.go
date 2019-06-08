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

type AddressListRtnJson struct {
	models.NideshopAddress
	ProviceName  string `json:"provice_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

func (this *AddressController) Address_List() {

	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)
	var addresses []models.NideshopAddress

	o.QueryTable(addresstable).Filter("user_id", getLoginUserId()).All(&addresses)

	rtnaddress := make([]AddressListRtnJson, 0)

	for _, val := range addresses {

		provicename := models.GetRegionName(val.ProvinceId)
		cityname := models.GetRegionName(val.CityId)
		distinctname := models.GetRegionName(val.DistrictId)
		rtnaddress = append(rtnaddress, AddressListRtnJson{
			NideshopAddress: val,
			ProviceName:     provicename,
			CityName:        cityname,
			DistrictName:    distinctname,
			FullRegion:      provicename + cityname + distinctname,
		})

	}

	utils.ReturnHTTPSuccess(&this.Controller, rtnaddress)
	this.ServeJSON()

}
func (this *AddressController) Address_Detail() {
	id := this.GetString("id")

	intid := utils.String2Int(id)

	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)
	var address models.NideshopAddress

	err := o.QueryTable(addresstable).Filter("id", intid).Filter("user_id", getLoginUserId()).One(&address)

	var val AddressListRtnJson

	if err != orm.ErrNoRows {

		provicename := models.GetRegionName(address.ProvinceId)
		cityname := models.GetRegionName(address.CityId)
		distinctname := models.GetRegionName(address.DistrictId)
		val = AddressListRtnJson{
			NideshopAddress: address,
			ProviceName:     provicename,
			CityName:        cityname,
			DistrictName:    distinctname,
			FullRegion:      provicename + cityname + distinctname,
		}
	}
	utils.ReturnHTTPSuccess(&this.Controller, val)
	this.ServeJSON()
}

type AddressSaveBody struct {
	Address    string `json:"address"`
	CityId     int    `json:"city_id"`
	DistrictId int    `json:"district_id"`
	IsDefault  bool   `json:"is_default"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	ProvinceId int    `json:"province_id"`
	AddressId  int    `json:"address_id"`
}

func (this *AddressController) Address_Save() {

	var asb AddressSaveBody
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body, &asb)

	address := asb.Address
	name := asb.Name
	mobile := asb.Mobile
	provinceid := asb.ProvinceId
	cityid := asb.CityId
	distinctid := asb.DistrictId
	isdefault := asb.IsDefault
	addressid := asb.AddressId
	userid := getLoginUserId()
	var intisdefault int
	if isdefault {
		intisdefault = 1
	} else {
		intisdefault = 0
	}

	intcityid := cityid
	intprovinceid := provinceid
	intdistinctid := distinctid

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
		IsDefault:  intisdefault,
	}
	o := orm.NewOrm()
	addresstable := new(models.NideshopAddress)

	var intid int64
	if addressid == 0 {
		id, err := o.Insert(&addressdata)
		if err == nil {
			intid = id
		}
	} else {
		o.QueryTable(addresstable).Filter("id", intid).Filter("user_id", userid).Update(orm.Params{
			"is_default": 0,
		})
	}

	if isdefault {
		_, err := o.Raw("UPDATE nideshop_address SET is_default = 0 where id <> ? and user_id = ?", intid, userid).Exec()
		if err == nil {
			//res.RowsAffected()
			//fmt.Println("mysql row affected nums: ", num)
		}
	}
	var addressinfo models.NideshopAddress
	o.QueryTable(addresstable).Filter("id", intid).One(&addressinfo)

	utils.ReturnHTTPSuccess(&this.Controller, addressinfo)
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
