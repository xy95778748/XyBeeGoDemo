package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserModel struct {
	Id         int
	Name       string
	Pass       string
	CreateTime int64
	UpdateTime int64
}

func init() {

	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/UserCenter?charset=utf8", 30)

	orm.RegisterModel(new(UserModel))

	orm.RunSyncdb("default", false, true)
}

func UserIsExist(name string) (*UserModel, bool) {

	mORM := orm.NewOrm()
	user := new(UserModel)
	user.Name = name
	if err := mORM.Read(user, "Name"); err == nil {
		return user, true
	} else {
		return nil, false
	}
}

func InsetUser(user *UserModel) bool {

	mORM := orm.NewOrm()
	_, err := mORM.Insert(user)
	return err == nil
}

