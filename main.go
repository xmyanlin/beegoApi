package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "beegoApi/model"
	_ "beegoApi/router"
)

func main()  {
	beego.SetLevel(beego.LevelInformational)
	beego.AppConfig.Set("version", "1.0.0")
	beego.Run()
}
