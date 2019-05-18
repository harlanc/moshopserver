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
	Desc   string `orm:"not null default '' VARCHAR(255)"`
	Height int    `orm:"not null default 0 SMALLINT(5)"`
	Id     int    `orm:"not null pk autoincr TINYINT(3)"`
	Name   string `orm:"not null default '' VARCHAR(60)"`
	Width  int    `orm:"not null default 0 SMALLINT(5)"`
}

type NideshopAddress struct {
	Address    string `orm:"not null default '' VARCHAR(120)"`
	CityId     int    `orm:"not null default 0 SMALLINT(5)"`
	CountryId  int    `orm:"not null default 0 SMALLINT(5)"`
	DistrictId int    `orm:"not null default 0 SMALLINT(5)"`
	Id         int    `orm:"not null pk autoincr MEDIUMINT(8)"`
	IsDefault  int    `orm:"not null default 0 TINYINT(1)"`
	Mobile     string `orm:"not null default '' VARCHAR(60)"`
	Name       string `orm:"not null default '' VARCHAR(50)"`
	ProvinceId int    `orm:"not null default 0 SMALLINT(5)"`
	UserId     int    `orm:"not null default 0 index MEDIUMINT(8)"`
}

type NideshopAdmin struct {
	AddTime       int    `orm:"not null default 0 INT(11)"`
	AdminRoleId   int    `orm:"not null default 0 INT(11)"`
	Avatar        string `orm:"not null default '''' VARCHAR(255)"`
	Id            int    `orm:"not null pk autoincr INT(11)"`
	LastLoginIp   string `orm:"not null default '''' VARCHAR(60)"`
	LastLoginTime int    `orm:"not null default 0 INT(11)"`
	Password      string `orm:"not null default '''' VARCHAR(255)"`
	PasswordSalt  string `orm:"not null default '''' VARCHAR(255)"`
	UpdateTime    int    `orm:"not null default 0 INT(11)"`
	Username      string `orm:"not null default '''' VARCHAR(10)"`
}

type NideshopAttribute struct {
	AttributeCategoryId int    `orm:"not null default 0 index INT(11)"`
	Id                  int    `orm:"not null pk autoincr INT(11)"`
	InputType           int    `orm:"not null default 1 TINYINT(1)"`
	Name                string `orm:"not null default '' VARCHAR(60)"`
	SortOrder           int    `orm:"not null default 0 TINYINT(3)"`
	Values              string `orm:"not null TEXT"`
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
	Checked                   int     `orm:"not null default 1 TINYINT(1)"`
	GoodsId                   int     `orm:"not null default 0 MEDIUMINT(8)"`
	GoodsName                 string  `orm:"not null default '' VARCHAR(120)"`
	GoodsSn                   string  `orm:"not null default '' VARCHAR(60)"`
	GoodsSpecifitionIds       string  `orm:"not null default '' comment('product表对应的goods_specifition_ids') VARCHAR(60)"`
	GoodsSpecifitionNameValue string  `orm:"not null comment('规格属性组成的字符串，用来显示用') TEXT"`
	Id                        int     `orm:"not null pk autoincr MEDIUMINT(8)"`
	ListPicUrl                string  `orm:"not null default '' VARCHAR(255)"`
	MarketPrice               float64 `orm:"not null default 0.00 DECIMAL(10,2)"`
	Number                    int     `orm:"not null default 0 SMALLINT(5)"`
	ProductId                 int     `orm:"not null default 0 MEDIUMINT(8)"`
	RetailPrice               float64 `orm:"not null default 0.00 DECIMAL(10,2)"`
	SessionId                 string  `orm:"not null default '' index CHAR(32)"`
	UserId                    int     `orm:"not null default 0 MEDIUMINT(8)"`
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
	AddTime     int `orm:"not null default 0 INT(11)"`
	Id          int `orm:"not null pk autoincr MEDIUMINT(8)"`
	IsAttention int `orm:"not null default 0 comment('是否是关注') index TINYINT(1)"`
	TypeId      int `orm:"not null default 0 INT(2)"`
	UserId      int `orm:"not null default 0 index MEDIUMINT(8)"`
	ValueId     int `orm:"not null default 0 index MEDIUMINT(8)"`
}

type NideshopComment struct {
	AddTime    int64  `orm:"not null default 0 BIGINT(12)"`
	Content    string `orm:"not null comment('储存为base64编码') VARCHAR(6550)"`
	Id         int    `orm:"not null pk autoincr INT(11)"`
	NewContent string `orm:"not null default '' VARCHAR(6550)"`
	Status     int    `orm:"not null default 0 TINYINT(3)"`
	TypeId     int    `orm:"not null default 0 TINYINT(3)"`
	UserId     int    `orm:"not null default 0 INT(11)"`
	ValueId    int    `orm:"not null default 0 index INT(11)"`
}

type NideshopCommentPicture struct {
	CommentId int    `orm:"not null default 0 INT(11)"`
	Id        int    `orm:"not null pk autoincr INT(11)"`
	PicUrl    string `orm:"not null default '' VARCHAR(255)"`
	SortOrder int    `orm:"not null default 5 TINYINT(1)"`
}

type NideshopCoupon struct {
	Id             int    `orm:"not null pk autoincr SMALLINT(5)"`
	MaxAmount      string `orm:"not null default 0.00 DECIMAL(10,2)"`
	MinAmount      string `orm:"not null default 0.00 DECIMAL(10,2)"`
	MinGoodsAmount string `orm:"not null default 0.00 DECIMAL(10,2)"`
	Name           string `orm:"not null default '' VARCHAR(60)"`
	SendEndDate    int    `orm:"not null default 0 INT(11)"`
	SendStartDate  int    `orm:"not null default 0 INT(11)"`
	SendType       int    `orm:"not null default 0 TINYINT(3)"`
	TypeMoney      string `orm:"not null default 0.00 DECIMAL(10,2)"`
	UseEndDate     int    `orm:"not null default 0 INT(11)"`
	UseStartDate   int    `orm:"not null default 0 INT(11)"`
}

type NideshopFeedback struct {
	MessageImg string `orm:"not null default '0' VARCHAR(255)"`
	MsgArea    int    `orm:"not null default 0 TINYINT(1)"`
	MsgContent string `orm:"not null TEXT"`
	MsgId      int    `orm:"not null pk autoincr MEDIUMINT(8)"`
	MsgStatus  int    `orm:"not null default 0 TINYINT(1)"`
	MsgTime    int    `orm:"not null default 0 INT(10)"`
	MsgTitle   string `orm:"not null default '' VARCHAR(200)"`
	MsgType    int    `orm:"not null default 0 TINYINT(1)"`
	OrderId    int    `orm:"not null default 0 INT(11)"`
	ParentId   int    `orm:"not null default 0 MEDIUMINT(8)"`
	UserEmail  string `orm:"not null default '' VARCHAR(60)"`
	UserId     int    `orm:"not null default 0 index MEDIUMINT(8)"`
	UserName   string `orm:"not null default '' VARCHAR(60)"`
}

type NideshopFootprint struct {
	AddTime int64 `orm:"not null default 0 INT(11)"`
	GoodsId int   `orm:"not null default 0 INT(11)"`
	Id      int   `orm:"not null pk autoincr INT(11)"`
	UserId  int   `orm:"not null default 0 INT(11)"`
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
	AttributeId int    `orm:"not null default 0 index INT(11)"`
	GoodsId     int    `orm:"not null default 0 index INT(11)"`
	Id          int    `orm:"not null pk autoincr INT(11)"`
	Value       string `orm:"not null TEXT"`
}

type NideshopGoodsGallery struct {
	GoodsId   int    `orm:"not null default 0 index INT(11)"`
	Id        int    `orm:"not null pk autoincr INT(11)"`
	ImgDesc   string `orm:"not null default '' VARCHAR(255)"`
	ImgUrl    string `orm:"not null default '' VARCHAR(255)"`
	SortOrder int    `orm:"not null default 5 INT(11)"`
}

type NideshopGoodsIssue struct {
	Answer   string `orm:"VARCHAR(45)"`
	GoodsId  string `orm:"TEXT"`
	Id       int    `orm:"not null pk autoincr INT(11)"`
	Question string `orm:"VARCHAR(255)"`
}

type NideshopGoodsSpecification struct {
	GoodsId         int    `orm:"not null default 0 index INT(11)"`
	Id              int    `orm:"not null pk autoincr INT(11)"`
	PicUrl          string `orm:"not null default '' VARCHAR(255)"`
	SpecificationId int    `orm:"not null default 0 index INT(11)"`
	Value           string `orm:"not null default '' VARCHAR(50)"`
}

type NideshopKeywords struct {
	Id        int    `orm:"not null pk autoincr INT(11)"`
	IsDefault int    `orm:"not null default 0 TINYINT(1)"`
	IsHot     int    `orm:"not null default 0 TINYINT(1)"`
	IsShow    int    `orm:"not null default 1 TINYINT(1)"`
	Keyword   string `orm:"not null default '' VARCHAR(90)"`
	SchemeUrl string `orm:"not null default '' comment('关键词的跳转链接') VARCHAR(255)"`
	SortOrder int    `orm:"not null default 100 INT(11)"`
	Type      int    `orm:"not null default 0 INT(11)"`
}

type NideshopOrder struct {
	ActualPrice    float64 `orm:"not null default 0.00 comment('实际需要支付的金额') DECIMAL(10,2)"`
	AddTime        int64   `orm:"not null default 0 INT(11)"`
	Address        string  `orm:"not null default '' VARCHAR(255)"`
	CallbackStatus string  `orm:"default 'true' ENUM('false','true')"`
	City           int     `orm:"not null default 0 SMALLINT(5)"`
	ConfirmTime    int     `orm:"not null default 0 INT(11)"`
	Consignee      string  `orm:"not null default '' VARCHAR(60)"`
	Country        int     `orm:"not null default 0 SMALLINT(5)"`
	CouponId       int     `orm:"not null default 0 comment('使用的优惠券id') MEDIUMINT(8)"`
	CouponPrice    float64 `orm:"not null default 0.00 DECIMAL(10,2)"`
	District       int     `orm:"not null default 0 SMALLINT(5)"`
	FreightPrice   float64 `orm:"not null default 0 comment('配送费用') INT(10)"`
	GoodsPrice     float64 `orm:"not null default 0.00 comment('商品总价') DECIMAL(10,2)"`
	Id             int     `orm:"not null pk autoincr MEDIUMINT(8)"`
	Integral       int     `orm:"not null default 0 INT(10)"`
	IntegralMoney  string  `orm:"not null default 0.00 DECIMAL(10,2)"`
	Mobile         string  `orm:"not null default '' VARCHAR(60)"`
	OrderPrice     float64 `orm:"not null default 0.00 comment('订单总价') DECIMAL(10,2)"`
	OrderSn        string  `orm:"not null default '' unique VARCHAR(20)"`
	OrderStatus    int     `orm:"not null default 0 index TINYINT(1)"`
	ParentId       int     `orm:"not null default 0 MEDIUMINT(8)"`
	PayId          int     `orm:"not null default 0 index TINYINT(3)"`
	PayName        string  `orm:"not null default '' VARCHAR(120)"`
	PayStatus      int     `orm:"not null default 0 index TINYINT(1)"`
	PayTime        int     `orm:"not null default 0 INT(11)"`
	Postscript     string  `orm:"not null default '' VARCHAR(255)"`
	Province       int     `orm:"not null default 0 SMALLINT(5)"`
	ShippingFee    string  `orm:"not null default 0.00 DECIMAL(10,2)"`
	ShippingStatus int     `orm:"not null default 0 index TINYINT(1)"`
	UserId         int     `orm:"not null default 0 index MEDIUMINT(8)"`
}

type NideshopOrderExpress struct {
	AddTime      int    `orm:"not null default 0 comment('添加时间') INT(11)"`
	Id           int    `orm:"not null pk autoincr MEDIUMINT(8)"`
	IsFinish     int    `orm:"not null default 0 TINYINT(1)"`
	LogisticCode string `orm:"not null default '' comment('快递单号') VARCHAR(20)"`
	OrderId      int    `orm:"not null default 0 index MEDIUMINT(8)"`
	RequestCount int    `orm:"default 0 comment('总查询次数') INT(11)"`
	RequestTime  int    `orm:"default 0 comment('最近一次向第三方查询物流信息时间') INT(11)"`
	ShipperCode  string `orm:"not null default '' comment('物流公司代码') VARCHAR(60)"`
	ShipperId    int    `orm:"not null default 0 MEDIUMINT(8)"`
	ShipperName  string `orm:"not null default '' comment('物流公司名称') VARCHAR(120)"`
	Traces       string `orm:"not null default '' comment('物流跟踪信息') VARCHAR(2000)"`
	UpdateTime   int    `orm:"not null default 0 comment('更新时间') INT(11)"`
}

type NideshopOrderGoods struct {
	GoodsId                   int     `orm:"not null default 0 index MEDIUMINT(8)"`
	GoodsName                 string  `orm:"not null default '' VARCHAR(120)"`
	GoodsSn                   string  `orm:"not null default '' VARCHAR(60)"`
	GoodsSpecifitionIds       string  `orm:"not null default '' VARCHAR(255)"`
	GoodsSpecifitionNameValue string  `orm:"not null TEXT"`
	Id                        int     `orm:"not null pk autoincr MEDIUMINT(8)"`
	IsReal                    int     `orm:"not null default 0 TINYINT(1)"`
	ListPicUrl                string  `orm:"not null default '' VARCHAR(255)"`
	MarketPrice               float64 `orm:"not null default 0.00 DECIMAL(10,2)"`
	Number                    int     `orm:"not null default 1 SMALLINT(5)"`
	OrderId                   int     `orm:"not null default 0 index MEDIUMINT(8)"`
	ProductId                 int     `orm:"not null default 0 MEDIUMINT(8)"`
	RetailPrice               float64 `orm:"not null default 0.00 DECIMAL(10,2)"`
}

type NideshopProduct struct {
	GoodsId               int     `orm:"not null default 0 MEDIUMINT(8)"`
	GoodsNumber           int     `orm:"not null default 0 MEDIUMINT(8)"`
	GoodsSn               string  `orm:"not null default '' VARCHAR(60)"`
	GoodsSpecificationIds string  `orm:"not null default '' VARCHAR(50)"`
	Id                    int     `orm:"not null pk autoincr MEDIUMINT(8)"`
	RetailPrice           float64 `orm:"not null default 0.00 DECIMAL(10,2)"`
}

type NideshopRegion struct {
	AgencyId int    `orm:"not null default 0 index SMALLINT(5)"`
	Id       int    `orm:"not null pk autoincr SMALLINT(5)"`
	Name     string `orm:"not null default '' VARCHAR(120)"`
	ParentId int    `orm:"not null default 0 index SMALLINT(5)"`
	Type     int    `orm:"not null default 2 index TINYINT(1)"`
}

type NideshopRelatedGoods struct {
	GoodsId        int `orm:"not null default 0 INT(11)"`
	Id             int `orm:"not null pk autoincr INT(11)"`
	RelatedGoodsId int `orm:"not null default 0 INT(11)"`
}

type NideshopSearchHistory struct {
	AddTime int64  `orm:"not null default 0 comment('搜索时间') INT(11)"`
	From    string `orm:"not null default '' comment('搜索来源，如PC、小程序、APP等') VARCHAR(45)"`
	Id      int    `orm:"not null pk autoincr INT(10)"`
	Keyword string `orm:"not null CHAR(50)"`
	UserId  string `orm:"VARCHAR(45)"`
}

type NideshopShipper struct {
	Code      string `orm:"not null default '' comment('快递公司代码') VARCHAR(10)"`
	Id        int    `orm:"not null pk autoincr unique INT(11)"`
	Name      string `orm:"not null default '' comment('快递公司名称') VARCHAR(20)"`
	SortOrder int    `orm:"not null default 10 comment('排序') INT(11)"`
}

type NideshopSpecification struct {
	Id        int    `orm:"not null pk autoincr INT(11)"`
	Name      string `orm:"not null default '' VARCHAR(60)"`
	SortOrder int    `orm:"not null default 0 TINYINT(3)"`
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
	Id     int    `orm:"not null pk autoincr INT(11)"`
	PicUrl string `orm:"not null default '' VARCHAR(255)"`
	Title  string `orm:"not null default '' VARCHAR(255)"`
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
	CouponId     int    `orm:"not null default 0 TINYINT(3)"`
	CouponNumber string `orm:"not null default '' VARCHAR(20)"`
	Id           int    `orm:"not null pk autoincr MEDIUMINT(8)"`
	OrderId      int    `orm:"not null default 0 MEDIUMINT(8)"`
	UsedTime     int    `orm:"not null default 0 INT(10)"`
	UserId       int    `orm:"not null default 0 index INT(11)"`
}

type NideshopUserLevel struct {
	Description string `orm:"not null default '' VARCHAR(255)"`
	Id          int    `orm:"not null pk autoincr TINYINT(3)"`
	Name        string `orm:"not null default '' VARCHAR(30)"`
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
	orm.RegisterModel(new(NideshopChannel))
	orm.RegisterModel(new(NideshopGoods))
	orm.RegisterModel(new(NideshopBrand))
	orm.RegisterModel(new(NideshopTopic))
	orm.RegisterModel(new(NideshopCategory))

	orm.RegisterModel(new(NideshopUser))
	orm.RegisterModel(new(NideshopAttributeCategory))
}
