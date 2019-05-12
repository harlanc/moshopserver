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

type FilterCategory struct {
	Id      int
	Name    string
	Checked bool
}

type ListRtnJson struct {
	utils.PageData
	FilterCategories []FilterCategory
	GoodsList        []orm.Params
}

type Banner struct {
	Url     string
	Name    string
	Img_url string
}

type NewRtnJson struct {
	BannerInfo Banner
}

type HotRtnJson struct {
	BannerInfo Banner
}

type CountRtnJson struct {
	GoodsCount int64
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

	userhascollect := models.IsUserHasCollect(loginuserid, 0, intGoodsId)

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

	o := orm.NewOrm()
	goodstable := new(models.NideshopGoods)
	rs := o.QueryTable(goodstable)

	if isNew != "" {
		rs = rs.Filter("is_new", isNew)
	}
	if isHot != "" {
		rs = rs.Filter("is_hot", isHot)
	}

	if keyword != "" {
		rs = rs.Filter("icontains", keyword)
		searchhistory := models.NideshopSearchHistory{Keyword: keyword, UserId: getLoginUserId(),
			AddTime: utils.GetTimestamp()}
		o.Insert(&searchhistory)
	}
	if brandId != "" {
		rs = rs.Filter("brand_id", brandId)
	}

	var categoryids []orm.Params
	rs.Limit(10000).Values(&categoryids, "category_id")
	categoryintids := utils.ExactMapValues2Int64Array(categoryids, "Id")

	var filterCategories = []FilterCategory{FilterCategory{Id: 0, Name: "全部", Checked: false}}

	if len(categoryintids) > 0 {
		var parentids []orm.Params
		categorytable := new(models.NideshopCategory)
		o.QueryTable(categorytable).Filter("id__in", categoryintids).Limit(10000).Values(&parentids, "parent_id")
		parentintids := utils.ExactMapValues2Int64Array(parentids, "ParentId")

		var parentcategories []orm.Params
		o.QueryTable(categorytable).Filter("id__in", parentintids).OrderBy("sort_order").Values(&parentcategories, "id", "name")

		for _, value := range parentcategories {
			id := value["id"].(int)
			checked := (categoryId == "" && id == 0)
			filterId := utils.String2Int(categoryId)

			filterCategories = append(filterCategories, FilterCategory{Id: filterId, Name: value["name"].(string), Checked: checked})
		}
	}

	if categoryId != "" {
		intcategoryId := utils.String2Int(categoryId)
		if intcategoryId > 0 {
			rs = rs.Filter("category_id__in", models.GetCategoryWhereIn(intcategoryId))
		}
	}

	if sort == "price" {
		orderstr := "retail_price"
		if order == "desc" {
			orderstr = "-" + orderstr
		}
		rs = rs.OrderBy(orderstr)
	} else {
		rs = rs.OrderBy("-id")
	}

	var rawData []orm.Params
	rs.Values(&rawData, "id", "name", "list_pic_url", "retail_price")

	intPage := utils.String2Int(page)
	intSize := utils.String2Int(size)

	pageData := utils.GetPageData(rawData, intPage, intSize)

	data, err := json.Marshal(ListRtnJson{PageData: pageData, FilterCategories: filterCategories, GoodsList: pageData.Data.([]orm.Params)})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Filter() {

	categoryId := this.GetString("categoryId")
	keyword := this.GetString("keyword")
	isNew := this.GetString("isNew")
	isHot := this.GetString("isHot")

	o := orm.NewOrm()
	goodstable := new(models.NideshopGoods)
	rs := o.QueryTable(goodstable)

	if categoryId != "" {
		intcategoryId := utils.String2Int(categoryId)
		rs = rs.Filter("category_id__in", models.GetChildCategoryId(intcategoryId))
	}

	if isNew != "" {
		rs = rs.Filter("is_new", isNew)
	}
	if isHot != "" {
		rs = rs.Filter("is_hot", isHot)
	}

	if keyword != "" {
		rs = rs.Filter("icontains", keyword)
	}

	var filterCategories = []FilterCategory{FilterCategory{Id: 0, Name: "全部"}}

	var categoryids []orm.Params
	rs.Limit(10000).Values(&categoryids, "category_id")
	categoryintids := utils.ExactMapValues2Int64Array(categoryids, "Id")

	if len(categoryintids) > 0 {

		var parentids []orm.Params
		categorytable := new(models.NideshopCategory)
		o.QueryTable(categorytable).Filter("id__in", categoryintids).Limit(10000).Values(&parentids, "parent_id")
		parentintids := utils.ExactMapValues2Int64Array(parentids, "ParentId")

		var parentcategories []orm.Params
		rs.OrderBy("sort_order").Filter("id__in", parentintids).Values(&parentcategories, "id", "name")

		for _, value := range parentcategories {
			id := value["id"].(int)
			filterCategories = append(filterCategories, FilterCategory{Id: id, Name: value["name"].(string)})
		}

	}

	data, err := json.Marshal(filterCategories)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_New() {

	banner := Banner{Url: "", Name: "坚持初心，为你寻觅世间好物", Img_url: "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png"}
	data, err := json.Marshal(NewRtnJson{BannerInfo: banner})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

func (this *GoodsController) Goods_Hot() {

	banner := Banner{Url: "", Name: "大家都在买的严选好物", Img_url: "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png"}
	data, err := json.Marshal(HotRtnJson{BannerInfo: banner})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Related() {

	goodsId := this.GetString("id")

	o := orm.NewOrm()
	relatedgoodstable := new(models.NideshopRelatedGoods)
	var rgids []orm.Params
	o.QueryTable(relatedgoodstable).Filter("goods_id", utils.String2Int(goodsId)).Values(&rgids, "related_goods_id")
	rgintids := utils.ExactMapValues2Int64Array(rgids, "RelatedGoodsId")

	goodstable := new(models.NideshopGoods)
	var relatedgoods []orm.Params

	if len(rgids) == 0 {
		var goodscategory models.NideshopGoods
		o.QueryTable(goodstable).Filter("id", goodsId).One(&goodscategory)
		o.QueryTable(goodstable).Filter("category_id", goodscategory.CategoryId).Limit(8).Values(&relatedgoods, "id", "name", "list_pic_url", "retail_price")

	} else {
		o.QueryTable(goodstable).Filter("id__in", rgintids).Values(&relatedgoods, "id", "name", "list_pic_url", "retail_price")
	}

	data, err := json.Marshal(relatedgoods)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()
}

func (this *GoodsController) Goods_Count() {

	o := orm.NewOrm()
	goodstable := new(models.NideshopGoods)

	count, _ := o.QueryTable(goodstable).Filter("is_delete", 0).Filter("is_on_sale", 1).Count()

	data, err := json.Marshal(CountRtnJson{GoodsCount: count})
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
	this.ServeJSON()

}

// func (this *GoodsController) getLoginUserId() {
// 	//	return this.Ctx.UserId

// }
