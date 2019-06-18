package model

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"github.com/astaxie/beego/orm"
)

/**
 *
 * 初始化mysql数据
 */
func init(){
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		appConf.String("database::db_user"),
		appConf.String("database::db_passwd"),
		appConf.String("database::db_host"),
		appConf.String("database::db_port"),
		appConf.String("database::db_name"),
		appConf.String("database::db_charset"))
	orm.RegisterDataBase("default", "mysql", conn)


	orm.RegisterModel(new(UserModel))
}