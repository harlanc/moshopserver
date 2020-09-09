package controllers

import (
	"encoding/base64"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"moshopserver/models"
	"moshopserver/utils"
)

type GoodsController struct {
	beego.Controller
}

type SkuRtnJson struct {
	ProductList       []models.NideshopProduct   `json:"productList"`
	SpecificationList []models.SpecificationItem `json:"specificationList"`
}

type DetailRtnJson struct {
	SkuRtnJson
	Goods          models.NideshopGoods          `json:"info"`
	Galleries      []models.NideshopGoodsGallery `json:"gallery"`
	Attribute      []Attribute                   `json:"attribute"`
	Issues         []models.NideshopGoodsIssue   `json:"issue"`
	UserHasCollect int                           `json:"userHasCollect"`
	Comment        Comment                       `json:"comment"`
	Brand          models.NideshopBrand          `json:"brand"`
}

type CategoryRtnJson struct {
	CurCategory     models.NideshopCategory   `json:"currentCategory"`
	ParentCategory  models.NideshopCategory   `json:"parentCategory"`
	BrotherCategory []models.NideshopCategory `json:"brotherCategory"`
}

type Attribute struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type CommentUser struct {
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}

type CommentInfo struct {
	Content  string                          `json:"content"`
	AddTime  int64                           `json:"add_time"`
	NickName string                          `json:"nick_name"`
	Avatar   string                          `json:"avatar"`
	PicList  []models.NideshopCommentPicture `json:"pic_list"`
}

type Comment struct {
	Count int64       `json:"count"`
	Data  CommentInfo `json:"data"`
}

type FilterCategory struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

type ListRtnJson struct {
	utils.PageData
	FilterCategories []FilterCategory `json:"filterCategory"`
	GoodsList        []orm.Params     `json:"goodsList"`
}

type Banner struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Img_url string `json:"imgurl"`
}

type NewRtnJson struct {
	BannerInfo Banner `json:"bannerinfo"`
}

type HotRtnJson struct {
	BannerInfo Banner `json:"bannerinfo"`
}

type CountRtnJson struct {
	GoodsCount int64 `json:"goodsCount"`
}

func updateJsonKeysGoods(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "ListPicUrl":
				delete(val, k)
				val["list_pic_url"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			}
		}
	}
}

func (this *GoodsController) Goods_Index() {
	o := orm.NewOrm()

	var goods []models.NideshopGoods
	good := new(models.NideshopGoods)
	o.QueryTable(good).All(&goods)

	utils.ReturnHTTPSuccess(&this.Controller, goods)
	this.ServeJSON()
}

func (this *GoodsController) Goods_Sku() {

	goodsId := this.GetString("id")
	goodsId_int := utils.String2Int(goodsId)

	plist := models.GetProductList(goodsId_int)
	slist := models.GetSpecificationList(goodsId_int)

	utils.ReturnHTTPSuccess(&this.Controller, SkuRtnJson{ProductList: plist, SpecificationList: slist})
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
	qb.Select("ga.value", "a.name").
		From("nideshop_goods_attribute ga").
		InnerJoin("nideshop_attribute a").On("ga.attribute_id = a.id").
		Where("ga.goods_id =" + goodsId).OrderBy("ga.id").Asc()
	sql := qb.String()
	o.Raw(sql).QueryRows(&attributes)

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
	loginuserid := getLoginUserId()

	userhascollect := models.IsUserHasCollect(loginuserid, 0, intGoodsId)

	models.AddFootprint(loginuserid, intGoodsId)

	plist := models.GetProductList(intGoodsId)
	slist := models.GetSpecificationList(intGoodsId)

	utils.ReturnHTTPSuccess(&this.Controller, DetailRtnJson{Goods: goodone, Galleries: galleries, Attribute: attributes,
		UserHasCollect: userhascollect, Issues: issues, Comment: commentval, Brand: *brand,
		SkuRtnJson: SkuRtnJson{ProductList: plist, SpecificationList: slist}})
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
	o.QueryTable(category).Filter("parent_id", curcategory.ParentId).All(&brothercategory)

	utils.ReturnHTTPSuccess(&this.Controller, CategoryRtnJson{CurCategory: curcategory,
		ParentCategory: parentcategory, BrotherCategory: brothercategory})
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

	var intsize int = 10
	if size != "" {
		intsize = utils.String2Int(size)
	}

	var intpage int = 1
	if page != "" {
		intpage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	goodstable := new(models.NideshopGoods)
	rs := o.QueryTable(goodstable)

	if isNew != "" {
		rs = rs.Filter("is_new", isNew)
	}
	if isHot != "" {
		rs = rs.Filter("is_hot", isHot)
	}

	keyword, _ = url.QueryUnescape(keyword)
	if keyword != "" {

		rs = rs.Filter("name__icontains", keyword)
		searchhistory := models.NideshopSearchHistory{Keyword: keyword, UserId: utils.Int2String(getLoginUserId()),
			AddTime: utils.GetTimestamp()}
		o.Insert(&searchhistory)
	}
	if brandId != "" {
		rs = rs.Filter("brand_id", brandId)
	}

	var categoryids []orm.Params
	rs.Limit(10000).Values(&categoryids, "category_id")
	categoryintids := utils.ExactMapValues2Int64Array(categoryids, "CategoryId")

	var filterCategories = []FilterCategory{FilterCategory{Id: 0, Name: "全部", Checked: false}}

	if len(categoryintids) > 0 {
		var parentids []orm.Params
		categorytable := new(models.NideshopCategory)
		o.QueryTable(categorytable).Filter("id__in", categoryintids).Limit(10000).Values(&parentids, "parent_id")
		parentintids := utils.ExactMapValues2Int64Array(parentids, "ParentId")

		var parentcategories []orm.Params
		o.QueryTable(categorytable).Filter("id__in", parentintids).OrderBy("sort_order").Values(&parentcategories, "id", "name")

		for _, value := range parentcategories {
			id := value["Id"].(int64)
			checked := (categoryId == "" && id == 0)

			filterCategories = append(filterCategories, FilterCategory{Id: id, Name: value["Name"].(string), Checked: checked})
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
	updateJsonKeysGoods(rawData)

	pageData := utils.GetPageData(rawData, intpage, intsize)

	utils.ReturnHTTPSuccess(&this.Controller, ListRtnJson{PageData: pageData, FilterCategories: filterCategories, GoodsList: pageData.Data.([]orm.Params)})
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
			id := value["id"].(int64)
			filterCategories = append(filterCategories, FilterCategory{Id: id, Name: value["name"].(string)})
		}
	}

	utils.ReturnHTTPSuccess(&this.Controller, filterCategories)
	this.ServeJSON()
}

func (this *GoodsController) Goods_New() {

	banner := Banner{Url: "", Name: "坚持初心，为你寻觅世间好物", Img_url: "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png"}

	utils.ReturnHTTPSuccess(&this.Controller, NewRtnJson{BannerInfo: banner})
	this.ServeJSON()

}

func (this *GoodsController) Goods_Hot() {

	banner := Banner{Url: "", Name: "大家都在买的严选好物", Img_url: "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png"}

	utils.ReturnHTTPSuccess(&this.Controller, HotRtnJson{BannerInfo: banner})
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
	updateJsonKeysGoods(relatedgoods)
	utils.ReturnHTTPSuccess(&this.Controller, relatedgoods)
	this.ServeJSON()
}

func (this *GoodsController) Goods_Count() {

	o := orm.NewOrm()
	goodstable := new(models.NideshopGoods)

	count, _ := o.QueryTable(goodstable).Filter("is_delete", 0).Filter("is_on_sale", 1).Count()

	utils.ReturnHTTPSuccess(&this.Controller, CountRtnJson{GoodsCount: count})
	this.ServeJSON()
}
