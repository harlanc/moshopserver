package routers

import (
	"github.com/astaxie/beego"
	"moshopserver/controllers"
)

func init() {

	beego.Router("api/index/index", &controllers.IndexController{}, "get:Index_Index")

	beego.Router("api/catalog/index", &controllers.CatalogController{}, "get:Catalog_Index")
	beego.Router("api/catalog/current", &controllers.CatalogController{}, "get:Catalog_Current")

	beego.Router("api/auth/loginByWeixin", &controllers.AuthController{}, "post:Auth_LoginByWeixin")

	beego.Router("api/goods/count", &controllers.GoodsController{}, "get:Goods_Count")
	beego.Router("api/goods/list", &controllers.GoodsController{}, "get:Goods_List")
	beego.Router("api/goods/category", &controllers.GoodsController{}, "get:Goods_Category")
	beego.Router("api/goods/detail", &controllers.GoodsController{}, "get:Goods_Detail")
	beego.Router("api/goods/new", &controllers.GoodsController{}, "get:Goods_New")
	beego.Router("api/goods/hot", &controllers.GoodsController{}, "get:Goods_Hot")
	beego.Router("api/goods/related", &controllers.GoodsController{}, "get:Goods_Related")

	beego.Router("api/brand/list", &controllers.BrandController{}, "get:Brand_List")
	beego.Router("api/brand/detail", &controllers.BrandController{}, "get:Brand_Detail")

	beego.Router("api/cart/index", &controllers.CartController{}, "get:Cart_Index")
	beego.Router("api/cart/add", &controllers.CartController{}, "post:Cart_Add")
	beego.Router("api/cart/update", &controllers.CartController{}, "post:Cart_Update")
	beego.Router("api/cart/delete", &controllers.CartController{}, "post:Cart_Delete")
	beego.Router("api/cart/checked", &controllers.CartController{}, "post:Cart_Checked")
	beego.Router("api/cart/goodscount", &controllers.CartController{}, "get:Cart_GoodsCount")
	beego.Router("api/cart/checkout", &controllers.CartController{}, "get:Cart_Checkout")

	beego.Router("api/pay/prepay", &controllers.PayController{}, "get:Pay_Prepay")

	beego.Router("api/collect/list", &controllers.CollectController{}, "get:Collect_List")
	beego.Router("api/collect/addordelete", &controllers.CollectController{}, "post:Collect_AddorDelete")

	beego.Router("api/comment/list", &controllers.CommentController{}, "get:Comment_List")
	beego.Router("api/comment/count", &controllers.CommentController{}, "get:Comment_Count")
	beego.Router("api/comment/post", &controllers.CommentController{}, "post:Comment_Post")

	beego.Router("api/topic/list", &controllers.TopicController{}, "get:Topic_List")
	beego.Router("api/topic/detail", &controllers.TopicController{}, "get:Topic_Detail")
	beego.Router("api/topic/related", &controllers.TopicController{}, "get:Topic_Related")

	beego.Router("api/search/index", &controllers.SearchController{}, "get:Search_Index")
	//beego.Router("api/search/result", &controllers.SearchController{}, "get:Topic_Detail")
	beego.Router("api/search/helper", &controllers.SearchController{}, "get:Search_Helper")
	beego.Router("api/search/clearhistory", &controllers.SearchController{}, "post:Search_Clearhistory")

	beego.Router("api/address/list", &controllers.AddressController{}, "get:Address_List")
	beego.Router("api/address/detail", &controllers.AddressController{}, "get:Address_Detail")
	beego.Router("api/address/save", &controllers.AddressController{}, "post:Address_Save")
	beego.Router("api/address/delete", &controllers.AddressController{}, "post:Address_Delete")

	beego.Router("api/region/list", &controllers.RegionController{}, "get:Region_List")

	beego.Router("api/order/submit", &controllers.OrderController{}, "post:Order_Submit")
	beego.Router("api/order/list", &controllers.OrderController{}, "get:Order_List")
	beego.Router("api/order/detail", &controllers.OrderController{}, "get:Order_Detail")
	//beego.Router("api/order/cancel", &controllers.OrderController{}, "post:Address_Save")
	beego.Router("api/order/express", &controllers.OrderController{}, "get:Order_Express")

	beego.Router("api/footprint/list", &controllers.FootprintController{}, "get:Footprint_List")
	beego.Router("api/footprint/delete", &controllers.FootprintController{}, "post:Footprint_Delete")

}
