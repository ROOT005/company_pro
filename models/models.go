package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Account  string
	PhoneNum string
	Created  string `orm:"index"`
}

//注册数据库
func RegisterDB() {
	//注册模型
	orm.RegisterModel(new(User))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库
	orm.RegisterDataBase("default", "mysql", "root:special005@/user?charset=utf8&loc=Local", 10)
}

//获取所有客户资料
func AddUser(name, account, phonenum string) error {
	o := orm.NewOrm()
	user := &User{
		Name:     name,
		Account:  account,
		PhoneNum: phonenum,
		Created:  time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err := o.Insert(user)
	if err != nil {
		return err
	}
	return err
}
