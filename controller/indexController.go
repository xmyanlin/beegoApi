package controller

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

type Response struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func (o *IndexController) Get() {
	o.Data["json"] = Response{0, "success.", "API 1.0"}
	o.ServeJSON()
}
