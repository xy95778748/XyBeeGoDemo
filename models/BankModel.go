package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type BankModel struct {
	Id int
	UserId string
	Balance float32
}

func (self *BankModel) InsertBalance() {

	mORM := orm.NewOrm()
	mORM.Insert(self)
}

func (self *BankModel) GetBalance() {

	mORM := orm.NewOrm()
	mORM.Using("default")
	err := mORM.Read(self, "userId")
	//var userId string
	//var balance string
	//err := mORM.Raw("select balance, id from bank_model where user_id = ?", self.UserId).QueryRow(&balance, &userId)
	if err != nil {
		beego.Error(err)
	}
}

func (self *BankModel) UpdateBalance() {

	mORM := orm.NewOrm()
	//_, err := mORM.Update(&self, "balance")
	_, err :=mORM.Raw("Update bank_model set balance = ? where user_id = ?", self.Balance, self.UserId).Exec()
	if err != nil {
		beego.Error(err)
	}
}