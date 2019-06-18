package model

import (
		"github.com/astaxie/beego/orm"
	)

type UserModel struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Phone    string    `json:"phone" orm:"column(phone);unique;size(11)"`
	Nickname string    `json:"nickname" orm:"column(nickname);unique;size(40);"`
	Password string    `json:"-" orm:"column(password);size(40)"`
	Created  int `json:"created" orm:"column(created)"`
	Updated  int `json:"-" orm:"column(updated)"`
}

func (u *UserModel) TableName() string {
	return "user"
}

func (usr *UserModel) Insert() error{
	usr.Phone = "129379"
	if _,err:=Select("default").Insert(usr);err != nil{
		return err
	}
	return  nil
}

func (usr *UserModel) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}
