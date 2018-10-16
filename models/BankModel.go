package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func init() {

	orm.RegisterDataBase("default", "mysql", "root:95778748@tcp(127.0.0.1:3306)/User?charset=utf8", 30)

	orm.RegisterModel(new(BankModel))

	orm.RunSyncdb("default", false, true)

	orm.Debug = true
}

type BankModel struct {
	Id int
	UserId string
	Balance float32
}

func (self *BankModel) GetBalance() {

	mORM := orm.NewOrm()
	err := mORM.Read(self, "userId")
	if err != nil {
		beego.Error(err)
	}
}

func (self *BankModel) UpdateBalance() {

	mORM := orm.NewOrm()
	mORM.Update(&self, "Balance")
}