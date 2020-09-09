package services

import (
	"encoding/json"

	"github.com/astaxie/beego/httplib"

	"github.com/astaxie/beego"
	"moshopserver/utils"
)

type Traces struct {
	AcceptTime    string `json:"accept_time"`
	AcceptStation string `json:"accept_station"`
	Remark        string `json:"remark"`
}

type ExpressRtnInfo struct {
	Success      bool     `json:"success"`
	ShipperCode  string   `json:"shipper_code"`
	ShipperName  string   `json:"shipper_name"`
	LogisticCode string   `json:"logistic_code"`
	IsFinish     int      `json:"is_finish"`
	Traces       []Traces `json:"traces"`
	RequestTime  int64    `json:"request_time"`
}

type ExpressResult struct {
	Success bool     `json:"success"`
	State   int      `json:"state"`
	Traces  []Traces `json:"traces"`
}

func QueryExpress(shippercode, logisticcode string, ordercode string) ExpressRtnInfo {
	var expressinfo ExpressRtnInfo = ExpressRtnInfo{
		Success:      false,
		ShipperCode:  shippercode,
		ShipperName:  "",
		LogisticCode: logisticcode,
		IsFinish:     0,
		Traces:       make([]Traces, 0),
	}
	fromdata := GenerateFromData(shippercode, logisticcode, ordercode)

	posturl := beego.AppConfig.String("express::request_url")

	req := httplib.Post(posturl)
	req.Header("content-type", "application/x-www-form-urlencoded")
	//need fix
	jsondata, _ := json.Marshal(fromdata)
	req.Param("form", string(jsondata))

	var res ExpressResult
	req.ToJSON(&res)
	expressinfo.Success = res.Success
	if res.State == 3 {
		expressinfo.IsFinish = 1
	}
	expressinfo.Traces = append(expressinfo.Traces, res.Traces...)

	return expressinfo

}

type ExpressFromData struct {
	RequestData string
	EBusinessId string
	RequestType string
	DataSign    string
	DataType    string
}

func GenerateFromData(shippercode, logisticcode, ordercode string) ExpressFromData {
	requestdata := GenerateRequestData(shippercode, logisticcode, ordercode)
	encoderequestdata, _ := utils.UrlEncode(requestdata)
	return ExpressFromData{
		RequestData: encoderequestdata,
		EBusinessId: beego.AppConfig.String("express::appid"),
		RequestType: "1002",
		DataSign:    GenerateDataSign(requestdata),
		DataType:    "2"}

}

type ExpressRequestData struct {
	OrderCode    string
	ShipperCode  string
	LogisticCode string
}

func GenerateRequestData(shippercode, logisticcode, ordercode string) string {

	data, err := json.Marshal(ExpressRequestData{ordercode, shippercode, logisticcode})
	if err != nil {
		return ""
	} else {
		return string(data)
	}
}

func GenerateDataSign(requestdata string) string {

	md5str := utils.Md5(requestdata)
	appkey := beego.AppConfig.String("express::appkey")
	base64str := utils.Base64Encode(md5str + appkey)
	rv, err := utils.UrlEncode(base64str)
	if err == nil {
		return ""
	} else {
		return rv
	}
}
