package main

import "github.com/astaxie/beego"

func main() {

	beego.Router("api/index/index")

	beego.Run() // listen and serve on 0.0.0.0:8080

}
