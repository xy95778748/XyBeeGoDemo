package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"XyBeeGoDemo/models"
	"XyBeeGoDemo/Commons"
	"strconv"
)

type BankMoney struct {
	beego.Controller
}


func (self *BankMoney) SaveMoney() {

	var res Commons.ResponseModel
	var saveMoney float32

	mBankModel := new(models.BankModel)

	self.Ctx.Input.Bind(&mBankModel.UserId, "userId")
	self.Ctx.Input.Bind(&saveMoney, "money")

	if len(mBankModel.UserId) == 0 {
		res.Failed("userId 不能为空")
	} else if saveMoney <= 0 {
		res.Failed("money 不能为0")
	} else {
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

	self.Data["json"] = &res
	self.ServeJSON()
}

func (self *BankMoney) WithdrawMoney () {

	var res Commons.ResponseModel
	var saveMoney float32

	mBankModel := new(models.BankModel)

	self.Ctx.Input.Bind(&mBankModel.UserId, "userId")
	self.Ctx.Input.Bind(&saveMoney, "money")

	if len(mBankModel.UserId) == 0 {
		res.Failed("userId 不能为空")
	} else if saveMoney <= 0 {
		res.Failed("money 不能为0")
	} else {
		mUserModel := new(models.UserModel)
		mUserModel.UserId = mBankModel.UserId
		if mUserModel.UserIdIsExist() {
			mBankModel.GetBalance()
			if mBankModel.Balance < saveMoney {
				res.Failed("余额不足")
			} else {
				mBankModel.Balance = mBankModel.Balance - saveMoney
				mBankModel.UpdateBalance()
				res.Success(mBankModel)
			}
		} else {
			res.Failed("账号不存在")
		}
	}

	self.Data["json"] = &res
	self.ServeJSON()
}

// private
func floatToDoublePoint (value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 32)
	return value
}