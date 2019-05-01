package routers

import (
	"github.com/astaxie/beego"
	"github.com/moshopserver/controllers"
)

func init() {

	beego.Router("api/index/index", &controllers.IndexController{}, "get:Get")
	// Register routers.
	//beego.Router("/", &controllers.AppController{})
	// Indicate AppController.Join method to handle POST requests.
	// beego.Router("/join", &controllers.AppController{}, "post:Join")

	// // Long polling.
	// beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	// beego.Router("/lp/post", &controllers.LongPollingController{})
	// beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// // WebSocket.
	// beego.Router("/ws", &controllers.WebSocketController{})
	// beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
