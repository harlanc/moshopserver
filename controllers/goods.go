package controllers

import (
	"encoding/base64"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/moshopserver/models"
	"github.com/moshopserver/utils"
)

type GoodsController struct {
	beego.Controller
}

type SkuRtnJson struct {
	ProductList       []models.NideshopProduct
	SpecificationList []models.SpecificationItem
}

type DetailRtnJson struct {
	SkuRtnJson
	Goods          models.NideshopGoods
	Galleries      []models.NideshopGoodsGallery
	Attribute      []Attribute
	Issues         []models.NideshopGoodsIssue
	UserHasCollect bool
	Comment        Comment
	Brand          models.NideshopBrand
}

type CategoryRtnJson struct {
	CurCategory     models.NideshopCategory
	ParentCategory  models.NideshopCategory
	BrotherCategory []models.NideshopCategory
}

type Attribute struct {
	Value string
	Name  string
}

type CommentUser struct {
	NickName string
	UserName string
	Avatar   string
}

type CommentInfo struct {
	Content  string
	AddTime  int64
	NickName string
	Avatar   string
	PicList  []models.NideshopCommentPicture
}

type Comment struct {
	Count int64
	Data  CommentInfo
}

func (this *GoodsController) Goods_Index() {
	o := orm.NewOrm()

	var goods []models.NideshopGoods
	good := new(models.NideshopGoods)
	o.QueryTable(good).All(&goods)

	data, err := json.Marshal(goods)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Sku() {

	goodsId := this.GetString("id")
	goodsId_int := utils.String2Int(goodsId)

	plist := models.GetProductList(goodsId_int)
	slist := models.GetSpecificationList(goodsId_int)

	data, err := json.Marshal(SkuRtnJson{ProductList: plist, SpecificationList: slist})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Detail() {

	goodsId := this.GetString("id")
	intGoodsId := utils.String2Int(goodsId)

	o := orm.NewOrm()

	var goodone models.NideshopGoods
	good := new(models.NideshopGoods)
	o.QueryTable(good).Filter("id", intGoodsId).One(&goodone)

	var galleries []models.NideshopGoodsGallery
	gallery := new(models.NideshopGoodsGallery)
	o.QueryTable(gallery).Filter("goods_id", intGoodsId).Limit(4).All(&galleries)

	qb, _ := orm.NewQueryBuilder("mysql")
	var attributes []Attribute
	qb.Select("gs.value", "a.name").
		From("nideshop_goods_attribute ga").
		InnerJoin("nideshop_attribute a").On("ga.attribute_id = a.id").
		Where("ga.goods_id =" + goodsId).OrderBy("ga.id").Asc()
	sql := qb.String()
	o.Raw(sql, 20).QueryRows(&attributes)

	var issues []models.NideshopGoodsIssue
	issue := new(models.NideshopGoodsIssue)
	o.QueryTable(issue).All(&issues)

	var brandone models.NideshopBrand
	brand := new(models.NideshopBrand)
	o.QueryTable(brand).Filter("id", goodone.BrandId).One(&brandone)

	comment := new(models.NideshopComment)
	commentCount, _ := o.QueryTable(comment).Filter("value_id", intGoodsId).Filter("type_id", 0).Count()
	var hotcommentone models.NideshopComment
	o.QueryTable(comment).Filter("value_id", intGoodsId).Filter("type_id", 0).One(&hotcommentone)

	var commentInfo CommentInfo

	if &hotcommentone != nil {

		user := new(models.NideshopUser)
		var commentUsers []orm.Params
		o.QueryTable(user).Filter("id", hotcommentone.UserId).Values(&commentUsers, "nickname", "username", "avatar")
		content, _ := base64.StdEncoding.DecodeString(hotcommentone.Content)

		var commentpictures []models.NideshopCommentPicture
		commentpicture := new(models.NideshopCommentPicture)
		o.QueryTable(commentpicture).Filter("comment_id", hotcommentone.Id).All(&commentpictures)

		commentInfo = CommentInfo{Content: string(content), AddTime: hotcommentone.AddTime, NickName: user.Nickname, Avatar: user.Avatar, PicList: commentpictures}

	}

	commentval := Comment{Count: commentCount, Data: commentInfo}
	loginuserid := utils.String2Int(getLoginUserId())

	userhascollect := isUserHasCollect(loginuserid, 0, intGoodsId)

	models.AddFootprint(loginuserid, intGoodsId)

	plist := models.GetProductList(intGoodsId)
	slist := models.GetSpecificationList(intGoodsId)

	data, err := json.Marshal(DetailRtnJson{Goods: goodone, Galleries: galleries, Attribute: attributes,
		UserHasCollect: userhascollect, Issues: issues, Comment: commentval, Brand: *brand,
		SkuRtnJson: SkuRtnJson{ProductList: plist, SpecificationList: slist}})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Category() {

	goodsId := this.GetString("id")
	intgoogsid := utils.String2Int(goodsId)

	o := orm.NewOrm()
	var curcategory models.NideshopCategory
	var parentcategory models.NideshopCategory
	var brothercategory []models.NideshopCategory

	category := new(models.NideshopCategory)

	o.QueryTable(category).Filter("id", intgoogsid).One(&curcategory)
	o.QueryTable(category).Filter("id", curcategory.ParentId).One(&parentcategory)
	o.QueryTable(category).Filter("parent_id", curcategory.ParentId).One(&brothercategory)

	data, err := json.Marshal(CategoryRtnJson{CurCategory: curcategory,
		ParentCategory: parentcategory, BrotherCategory: brothercategory})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}
func (this *GoodsController) Goods_List() {
	categoryId := this.GetString("categoryId")
	brandId := this.GetString("brandId")
	keyword := this.GetString("keyword")
	isNew := this.GetString("isNew")
	isHot := this.GetString("isHot")
	page := this.GetString("page")
	size := this.GetString("size")
	sort := this.GetString("sort")
	order := this.GetString("order")

}

func (this *GoodsController) getLoginUserId() {
	//	return this.Ctx.UserId

}
