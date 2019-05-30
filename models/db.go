package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type NideshopAd struct {
	AdPositionId int    `json:"ad_position_id"`
	Content      string `json:"content"`
	Enabled      int    `json:"enabled"`
	EndTime      int    `json:"end_time"`
	Id           int    `json:"id"`
	ImageUrl     string `json:"image_url"`
	Link         string `json:"link"`
	MediaType    int    `json:"media_type"`
	Name         string `json:"name"`
}

type NideshopAdPosition struct {
	Desc   string `json:"desc"`
	Height int    `json:"height"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Width  int    `json:"width"`
}

type NideshopAddress struct {
	Address    string `json:"address"`
	CityId     int    `json:"city_id"`
	CountryId  int    `json:"country_id"`
	DistrictId int    `json:"district_id"`
	Id         int    `json:"id"`
	IsDefault  int    `json:"is_default"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	ProvinceId int    `json:"province_id"`
	UserId     int    `json:"user_id"`
}

type NideshopAdmin struct {
	AddTime       int    `json:"add_time"`
	AdminRoleId   int    `json:"admin_role_id"`
	Avatar        string `json:"avatar"`
	Id            int    `json:"id"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime int    `json:"last_login_time"`
	Password      string `json:"password"`
	PasswordSalt  string `json:"password_salt"`
	UpdateTime    int    `json:"update_time"`
	Username      string `json:"username"`
}

type NideshopAttribute struct {
	AttributeCategoryId int    `json:"attribute_category_id"`
	Id                  int    `json:"id"`
	InputType           int    `json:"input_type"`
	Name                string `json:"name"`
	SortOrder           int    `json:"sort_order"`
	Values              string `json:"values"`
}

type NideshopAttributeCategory struct {
	Enabled int    `json:"enabled"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

type NideshopBrand struct {
	AppListPicUrl string `json:"app_list_pic_url"`
	FloorPrice    string `json:"floor_price"`
	Id            int    `json:"id"`
	IsNew         int    `json:"is_new"`
	IsShow        int    `json:"is_show"`
	ListPicUrl    string `json:"list_pic_url"`
	Name          string `json:"name"`
	NewPicUrl     string `json:"new_pic_url"`
	NewSortOrder  int    `json:"new_sort_order"`
	PicUrl        string `json:"pic_url"`
	SimpleDesc    string `json:"simple_desc"`
	SortOrder     int    `json:"sort_order"`
}

type NideshopCart struct {
	Checked                   int     `json:"checked"`
	GoodsId                   int     `json:"goods_id"`
	GoodsName                 string  `json:"goods_name"`
	GoodsSn                   string  `json:"goods_sn"`
	GoodsSpecifitionIds       string  `json:"goods_specifition_ids"`
	GoodsSpecifitionNameValue string  `json:"goods_specifition_name_value"`
	Id                        int     `json:"id"`
	ListPicUrl                string  `json:"list_pic_url"`
	MarketPrice               float64 `json:"market_price"`
	Number                    int     `json:"number"`
	ProductId                 int     `json:"product_id"`
	RetailPrice               float64 `json:"retail_price"`
	SessionId                 string  `json:"session_id"`
	UserId                    int     `json:"user_id"`
}

type NideshopCategory struct {
	BannerUrl    string `json:"banner_url"`
	FrontDesc    string `json:"front_desc"`
	FrontName    string `json:"front_name"`
	IconUrl      string `json:"icon_url"`
	Id           int    `json:"id"`
	ImgUrl       string `json:"img_url"`
	IsShow       int    `json:"is_show"`
	Keywords     string `json:"keywords"`
	Level        string `json:"level"`
	Name         string `json:"name"`
	ParentId     int    `json:"parent_id"`
	ShowIndex    int    `json:"show_index"`
	SortOrder    int    `json:"sort_order"`
	Type         int    `json:"type"`
	WapBannerUrl string `json:"wap_banner_url"`
}

type NideshopChannel struct {
	IconUrl   string `json:"icon_url"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
	Url       string `json:"url"`
}

type NideshopCollect struct {
	AddTime     int64 `json:"add_time"`
	Id          int   `json:"id"`
	IsAttention int   `json:"is_attention"`
	TypeId      int   `json:"type_id"`
	UserId      int   `json:"user_id"`
	ValueId     int   `json:"value_id"`
}

type NideshopComment struct {
	AddTime    int64  `json:"add_time"`
	Content    string `json:"content"`
	Id         int    `json:"id"`
	NewContent string `json:"new_content"`
	Status     int    `json:"status"`
	TypeId     int    `json:"type_id"`
	UserId     int    `json:"user_id"`
	ValueId    int    `json:"value_id"`
}

type NideshopCommentPicture struct {
	CommentId int    `json:"comment_id"`
	Id        int    `json:"id"`
	PicUrl    string `json:"pic_url"`
	SortOrder int    `json:"sort_order"`
}

type NideshopCoupon struct {
	Id             int    `json:"id"`
	MaxAmount      string `json:"max_amount"`
	MinAmount      string `json:"min_amount"`
	MinGoodsAmount string `json:"min_goods_amount"`
	Name           string `json:"name"`
	SendEndDate    int    `json:"send_end_date"`
	SendStartDate  int    `json:"send_start_date"`
	SendType       int    `json:"send_type"`
	TypeMoney      string `json:"type_money"`
	UseEndDate     int    `json:"use_end_date"`
	UseStartDate   int    `json:"use_start_date"`
}

type NideshopFeedback struct {
	MessageImg string `json:"message_img"`
	MsgArea    int    `json:"msg_area"`
	MsgContent string `json:"msg_content"`
	Id         int    `json:"msg_id"`
	MsgStatus  int    `json:"msg_status"`
	MsgTime    int    `json:"msg_time"`
	MsgTitle   string `json:"msg_title"`
	MsgType    int    `json:"msg_type"`
	OrderId    int    `json:"order_id"`
	ParentId   int    `json:"parent_id"`
	UserEmail  string `json:"user_email"`
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
}

type NideshopFootprint struct {
	AddTime int64 `json:"add_time"`
	GoodsId int   `json:"goods_id"`
	Id      int   `json:"id"`
	UserId  int   `json:"user_id"`
}

type NideshopGoods struct {
	AddTime           int    `json:"add_time"`
	AppExclusivePrice string `json:"app_exclusive_price"`
	AttributeCategory int    `json:"attribute_category"`
	BrandId           int    `json:"brand_id"`
	CategoryId        int    `json:"category_id"`
	CounterPrice      string `json:"counter_price"`
	ExtraPrice        string `json:"extra_price"`
	GoodsBrief        string `json:"goods_brief"`
	GoodsDesc         string `json:"goods_desc"`
	GoodsNumber       int    `json:"goods_number"`
	GoodsSn           string `json:"goods_sn"`
	GoodsUnit         string `json:"goods_unit"`
	Id                int    `json:"id"`
	IsAppExclusive    int    `json:"is_app_exclusive"`
	IsDelete          bool   `json:"is_delete"`
	IsHot             int    `json:"is_hot"`
	IsLimited         int    `json:"is_limited"`
	IsNew             int    `json:"is_new"`
	IsOnSale          int    `json:"is_on_sale"`
	Keywords          string `json:"keywords"`
	ListPicUrl        string `json:"list_pic_url"`
	Name              string `json:"name"`
	PrimaryPicUrl     string `json:"primary_pic_url"`
	PrimaryProductId  int    `json:"primary_product_id"`
	PromotionDesc     string `json:"promotion_desc"`
	PromotionTag      string `json:"promotion_tag"`
	RetailPrice       string `json:"retail_price"`
	SellVolume        int    `json:"sell_volume"`
	SortOrder         int    `json:"sort_order"`
	UnitPrice         string `json:"unit_price"`
}

type NideshopGoodsAttribute struct {
	AttributeId int    `json:"attribute_id"`
	GoodsId     int    `json:"goods_id"`
	Id          int    `json:"id"`
	Value       string `json:"value"`
}

type NideshopGoodsGallery struct {
	GoodsId   int    `json:"goods_id"`
	Id        int    `json:"id"`
	ImgDesc   string `json:"img_desc"`
	ImgUrl    string `json:"img_url"`
	SortOrder int    `json:"sort_order"`
}

type NideshopGoodsIssue struct {
	Answer   string `json:"answer"`
	GoodsId  string `json:"goods_id"`
	Id       int    `json:"id"`
	Question string `json:"question"`
}

type NideshopGoodsSpecification struct {
	GoodsId         int    `json:"goods_id"`
	Id              int    `json:"id"`
	PicUrl          string `json:"pic_url"`
	SpecificationId int    `json:"specification_id"`
	Value           string `json:"value"`
}

type NideshopKeywords struct {
	Id        int    `json:"id"`
	IsDefault int    `json:"is_default"`
	IsHot     int    `json:"is_hot"`
	IsShow    int    `json:"is_show"`
	Keyword   string `json:"keyword"`
	SchemeUrl string `json:"scheme_url"`
	SortOrder int    `json:"sort_order"`
	Type      int    `json:"type"`
}

type NideshopOrder struct {
	ActualPrice    float64 `json:"actual_price"`
	AddTime        int64   `json:"add_time"`
	Address        string  `json:"address"`
	CallbackStatus string  `json:"callback_status"`
	City           int     `json:"city"`
	ConfirmTime    int     `json:"confirm_time"`
	Consignee      string  `json:"consignee"`
	Country        int     `json:"country"`
	CouponId       int     `json:"coupon_id"`
	CouponPrice    float64 `json:"coupon_price"`
	District       int     `json:"district"`
	FreightPrice   float64 `json:"freight_price"`
	GoodsPrice     float64 `json:"goods_price"`
	Id             int     `json:"id"`
	Integral       int     `json:"integral"`
	IntegralMoney  string  `json:"integral_money"`
	Mobile         string  `json:"mobile"`
	OrderPrice     float64 `json:"order_price"`
	OrderSn        string  `json:"order_sn"`
	OrderStatus    int     `json:"order_status"`
	ParentId       int     `json:"parent_id"`
	PayId          int     `json:"pay_id"`
	PayName        string  `json:"pay_name"`
	PayStatus      int     `json:"pay_status"`
	PayTime        int     `json:"pay_time"`
	Postscript     string  `json:"postscript"`
	Province       int     `json:"province"`
	ShippingFee    string  `json:"shipping_fee"`
	ShippingStatus int     `json:"shipping_status"`
	UserId         int     `json:"user_id"`
}

type NideshopOrderExpress struct {
	AddTime      int    `json:"add_time"`
	Id           int    `json:"id"`
	IsFinish     int    `json:"is_finish"`
	LogisticCode string `json:"logistic_code"`
	OrderId      int    `json:"order_id"`
	RequestCount int    `json:"request_count"`
	RequestTime  int    `json:"request_time"`
	ShipperCode  string `json:"shipper_code"`
	ShipperId    int    `json:"shipper_id"`
	ShipperName  string `json:"shipper_name"`
	Traces       string `json:"traces"`
	UpdateTime   int    `json:"update_time"`
}

type NideshopOrderGoods struct {
	GoodsId                   int     `json:"goods_id"`
	GoodsName                 string  `json:"goods_name"`
	GoodsSn                   string  `json:"goods_sn"`
	GoodsSpecifitionIds       string  `json:"goods_specifition_ids"`
	GoodsSpecifitionNameValue string  `json:"goods_specifition_name_value"`
	Id                        int     `json:"id"`
	IsReal                    int     `json:"is_real"`
	ListPicUrl                string  `json:"list_pic_url"`
	MarketPrice               float64 `json:"market_price"`
	Number                    int     `json:"number"`
	OrderId                   int     `json:"order_id"`
	ProductId                 int     `json:"product_id"`
	RetailPrice               float64 `json:"retail_price"`
}

type NideshopProduct struct {
	GoodsId               int     `json:"goods_id"`
	GoodsNumber           int     `json:"goods_number"`
	GoodsSn               string  `json:"goods_sn"`
	GoodsSpecificationIds string  `json:"goods_specification_ids"`
	Id                    int     `json:"id"`
	RetailPrice           float64 `json:"retail_price"`
}

type NideshopRegion struct {
	AgencyId int    `json:"agency_id"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
	Type     int    `json:"type"`
}

type NideshopRelatedGoods struct {
	GoodsId        int `json:"goods_id"`
	Id             int `json:"id"`
	RelatedGoodsId int `json:"related_goods_id"`
}

type NideshopSearchHistory struct {
	AddTime int64  `json:"add_time"`
	From    string `json:"from"`
	Id      int    `json:"id"`
	Keyword string `json:"keyword"`
	UserId  string `json:"user_id"`
}

type NideshopShipper struct {
	Code      string `json:"code"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type NideshopSpecification struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
}

type NideshopTopic struct {
	Avatar          string `json:"avatar"`
	Content         string `json:"content"`
	Id              int    `json:"id"`
	IsShow          int    `json:"is_show"`
	ItemPicUrl      string `json:"item_pic_url"`
	PriceInfo       string `json:"price_info"`
	ReadCount       string `json:"read_count"`
	ScenePicUrl     string `json:"scene_pic_url"`
	SortOrder       int    `json:"sort_order"`
	Subtitle        string `json:"subtitle"`
	Title           string `json:"title"`
	TopicCategoryId int    `json:"topic_category_id"`
	TopicTagId      int    `json:"topic_tag_id"`
	TopicTemplateId int    `json:"topic_template_id"`
}

type NideshopTopicCategory struct {
	Id     int    `json:"id"`
	PicUrl string `json:"pic_url"`
	Title  string `json:"title"`
}

type NideshopUser struct {
	Avatar        string `json:"avatar"`
	Birthday      int    `json:"birthday"`
	Gender        int    `json:"gender"`
	Id            int    `json:"id"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime int64  `json:"last_login_time"`
	Mobile        string `json:"mobile"`
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	RegisterIp    string `json:"register_ip"`
	RegisterTime  int64  `json:"register_time"`
	UserLevelId   int    `json:"user_level_id"`
	Username      string `json:"username"`
	WeixinOpenid  string `json:"weixin_openid"`
}

type NideshopUserCoupon struct {
	CouponId     int    `json:"coupon_id"`
	CouponNumber string `json:"coupon_number"`
	Id           int    `json:"id"`
	OrderId      int    `json:"order_id"`
	UsedTime     int    `json:"used_time"`
	UserId       int    `json:"user_id"`
}

type NideshopUserLevel struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
}

// type User struct {
// 	Id   int    `orm:"not null pk autoincr INT(11)"`
// 	Name string `orm:"not null default '' VARCHAR(100)"`
// }

func init() {

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/nideshop?charset=utf8mb4", 30)

	// register model
	orm.RegisterModel(new(NideshopAd))
	orm.RegisterModel(new(NideshopAdPosition))
	orm.RegisterModel(new(NideshopAddress))
	orm.RegisterModel(new(NideshopAdmin))
	orm.RegisterModel(new(NideshopAttribute))
	orm.RegisterModel(new(NideshopAttributeCategory))

	orm.RegisterModel(new(NideshopBrand))
	orm.RegisterModel(new(NideshopCart))
	orm.RegisterModel(new(NideshopCategory))

	orm.RegisterModel(new(NideshopChannel))
	orm.RegisterModel(new(NideshopCollect))
	orm.RegisterModel(new(NideshopComment))
	orm.RegisterModel(new(NideshopCommentPicture))

	orm.RegisterModel(new(NideshopCoupon))
	orm.RegisterModel(new(NideshopFeedback))

	orm.RegisterModel(new(NideshopFootprint))
	orm.RegisterModel(new(NideshopGoods))

	orm.RegisterModel(new(NideshopGoodsAttribute))

	orm.RegisterModel(new(NideshopGoodsGallery))
	orm.RegisterModel(new(NideshopGoodsIssue))

	orm.RegisterModel(new(NideshopGoodsSpecification))
	orm.RegisterModel(new(NideshopKeywords))

	orm.RegisterModel(new(NideshopOrder))
	orm.RegisterModel(new(NideshopOrderExpress))

	orm.RegisterModel(new(NideshopOrderGoods))

	orm.RegisterModel(new(NideshopProduct))
	orm.RegisterModel(new(NideshopRegion))

	orm.RegisterModel(new(NideshopRelatedGoods))
	orm.RegisterModel(new(NideshopSearchHistory))

	orm.RegisterModel(new(NideshopShipper))
	orm.RegisterModel(new(NideshopSpecification))
	orm.RegisterModel(new(NideshopTopic))
	orm.RegisterModel(new(NideshopTopicCategory))

	orm.RegisterModel(new(NideshopUser))
	orm.RegisterModel(new(NideshopUserCoupon))
	orm.RegisterModel(new(NideshopUserLevel))

}
