package router

import (
	"github.com/astaxie/beego"
	"beegoApi/controller"
	"beegoApi/middleware"
)

func init(){
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/user", &controller.IndexController{}, "Get:Get"),
	)

	beego.Router("/user", &controller.IndexController{}, "Get:Get")

	beego.InsertFilter("/v1/*",beego.BeforeRouter, middleware.Token)
	beego.AddNamespace(ns)
}
