package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

type UserModel struct {
	Id         int
	Name       string
	Pass       string
	CreateTime int64
	UpdateTime int64
}

func init() {

	orm.RegisterDataBase("default", "mysql", "root:95778748@tcp(127.0.0.1:3306)/User?charset=utf8", 30)

	orm.RegisterModel(new(UserModel))

	orm.RunSyncdb("default", false, true)

	orm.Debug = true
}

func (self *UserModel) UserIsExist() bool {

	mORM := orm.NewOrm()

	return mORM.QueryTable(self).Filter("name", self.Name).Exist()
}

func (self *UserModel) InsetUser() bool {

	mORM := orm.NewOrm()
	_, err := mORM.Insert(self)
	if err != nil {
		beego.Error(err)
	}
	return err == nil
}

func (self *UserModel) ResetPassword() bool {

	fmt.Println(self)
	mORM := orm.NewOrm()
	_, err := mORM.Raw("UPDATE user_model SET pass = ? WHERE name = ?", self.Pass, self.Name).Exec()
	return err == nil
}
