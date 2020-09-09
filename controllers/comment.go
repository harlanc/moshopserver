package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"moshopserver/models"
	"moshopserver/utils"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Comment_Post() {
	typeId := this.GetString("typeId")
	valueId := this.GetString("valueId")
	content := this.GetString("content")

	inttypeId := utils.String2Int(typeId)
	intvalueId := utils.String2Int(valueId)

	var comment models.NideshopComment = models.NideshopComment{
		AddTime: utils.GetTimestamp(),
		Content: utils.Base64Encode(content),
		// Id
		// NewContent
		// Status
		TypeId:  inttypeId,
		UserId:  getLoginUserId(),
		ValueId: intvalueId,
	}

	o := orm.NewOrm()
	_, err := o.Insert(&comment)
	if err != nil {
		this.Abort("添加评论成功")
	} else {
		this.Abort("评论保存失败")
	}

}

type CommentCountRtnJson struct {
	AllCount    int64
	HasPicCount int
}

func (this *CommentController) Comment_Count() {

	typeId := this.GetString("typeId")
	valueId := this.GetString("valueId")
	inttypeId := utils.String2Int(typeId)
	intvalueId := utils.String2Int(valueId)

	o := orm.NewOrm()
	commenttable := new(models.NideshopComment)
	allcount, _ := o.QueryTable(commenttable).Filter("type_id", inttypeId).Filter("value_id", intvalueId).Count()

	qb, _ := orm.NewQueryBuilder("mysql")
	var list []models.NideshopComment

	qb.Select("nc.*").
		From("nideshop_comment nc").
		RightJoin("nideshop_comment_picture ncp").
		On("nc.id = ncp.comment_id").
		Where("nc.type_id =" + typeId + "and nc.value_id = " + valueId)

	sql := qb.String()
	o.Raw(sql).QueryRows(&list)
	haspiccount := len(list)

	utils.ReturnHTTPSuccess(&this.Controller, CommentCountRtnJson{allcount, haspiccount})
	this.ServeJSON()
}

//It may need to be refactored.
func GetCommentPageData(rawData []models.NideshopComment, page int, size int) utils.PageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []models.NideshopComment

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return utils.PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalpages, Data: pagedata}
}

type CommenListtRtnJson struct {
	Comment  string
	TypeId   int
	ValueId  int
	Id       int
	AddTime  string
	UserInfo orm.Params
	PicList  []models.NideshopCommentPicture
}

func (this *CommentController) Comment_List() {

	typeId := this.GetString("typeId")
	valueId := this.GetString("valueId")
	page := this.GetString("page")
	size := this.GetString("size")
	showType := this.GetString("showType")
	inttypeId := utils.String2Int(typeId)
	intvalueId := utils.String2Int(valueId)

	intshowtype := utils.String2Int(showType)

	var intsize int = 10
	if size != "" {
		intsize = utils.String2Int(size)
	}

	var intpage int = 1
	if page != "" {
		intpage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	commenttable := new(models.NideshopComment)
	var pagedata utils.PageData
	var comments []models.NideshopComment
	if intshowtype != 1 {
		o.QueryTable(commenttable).Filter("type_id", inttypeId).Filter("value_id", intvalueId).All(&comments)

	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("nc.*").
			From("nideshop_comment nc").
			InnerJoin("nideshop_comment_picture ncp").
			On("c.id = ncp.comment_id").
			Where("c.type_id =" + typeId + "and c.value_id = " + valueId)

		sql := qb.String()
		o := orm.NewOrm()
		o.Raw(sql).QueryRows(&comments)
	}

	pagedata = GetCommentPageData(comments, intpage, intsize)

	var rtncomments []CommenListtRtnJson
	usertable := new(models.NideshopUser)
	commentpictable := new(models.NideshopCommentPicture)

	for _, val := range pagedata.Data.([]models.NideshopComment) {

		var users []orm.Params
		var commentpictures []models.NideshopCommentPicture
		o.QueryTable(usertable).Filter("id", val.UserId).Values(&users, "username", "avatar", "nickname")
		o.QueryTable(commentpictable).Filter("comment_id", val.Id).All(&commentpictures)
		rtncomments = append(rtncomments, CommenListtRtnJson{
			Comment:  val.Content,
			TypeId:   val.TypeId,
			ValueId:  val.ValueId,
			Id:       val.Id,
			AddTime:  utils.FormatTimestamp(val.AddTime, "2006-01-02 03:04:05 PM"),
			UserInfo: users[0],
			PicList:  commentpictures,
		})

	}
	pagedata.Data = rtncomments

	utils.ReturnHTTPSuccess(&this.Controller, pagedata)
	this.ServeJSON()

}
