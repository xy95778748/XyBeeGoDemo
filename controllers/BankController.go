package controllers

import (
	"github.com/astaxie/beego"
	"XyBeeGoDemo/models"
	"XyBeeGoDemo/Commons"
)

type BankController struct {
	beego.Controller
}


func (self *BankController) SaveMoney() {

	var res Commons.ResponseModel
	var saveMoney float32

	mBankModel := new(models.BankModel)

	self.Ctx.Input.Bind(&mBankModel.UserId, "userId")
	self.Ctx.Input.Bind(&saveMoney, "money")

	mUserModel := new(models.UserModel)
	mUserModel.UserId = mBankModel.UserId
	if mUserModel.UserIdIsExist() {
		mBankModel.GetBalance()
		mBankModel.Balance = mBankModel.Balance + saveMoney
		mBankModel.UpdateBalance()
		res.Success(mBankModel)
	} else {
		res.Failed("账号不存在")
	}
}

func (self *BankController) WithdrawMoney () {


}