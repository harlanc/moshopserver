package utils

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	uuid "github.com/satori/go.uuid"
)

func String2Int(val string) int {

	goodsId_int, err := strconv.Atoi(val)
	if err != nil {
		return -1
	} else {
		return goodsId_int
	}
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}

func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	} else {
		return uuid.String()
	}
}

//the result likes 1423361979
func GetTimestamp() int64 {
	return time.Now().Unix()
}

//the result likes 2015-02-08 10:19:39 AM
func FormatTimestamp(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05 PM")
}

func ExactMapValues2Int64Array(maparray []orm.Params, key string) []int64 {

	var vals []int64
	for _, value := range maparray {
		vals = append(vals, value[key].(int64))
	}
	return vals
}

type PageData struct {
	NumsPerPage int
	CurrentPage int
	Count       int
	TotalPages  int
	Data        interface{}
}

//page begins from 1
func GetPageData(rawData []orm.Params, page int, size int) PageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []orm.Params

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalpages, Data: pagedata}
}
