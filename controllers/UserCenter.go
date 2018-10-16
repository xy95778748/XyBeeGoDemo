package controllers

import (
	"XyBeeGoDemo/Commons"
	"XyBeeGoDemo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseControllerInterface interface {
	CheckRequestMethod(method string)
}

type UserCenterController struct {
	beego.Controller
}

/* 登录接口 */
func (self *UserCenterController) Login() {

	mUserModel := new(models.UserModel)

	self.Ctx.Input.Bind(&mUserModel.Name, "name")
	self.Ctx.Input.Bind(&mUserModel.Pass, "pass")

	var res Commons.ResponseModel

	if len(mUserModel.Name) == 0 || len(mUserModel.Pass) == 0 {

		res.Failed("账号密码不能为空")
	}

	if mUserModel.UserIsExist() {
		if mUserModel.GetUserPass() == mUserModel.Pass {
			res.Success(nil)
		} else {
			res.Failed("密码错误")
		}
	} else {
		res.Failed("账号不存在")
	}

	self.Data["json"] = &res
	self.ServeJSON()

}

func (self *UserCenterController) Regist() {

	mUserModel := new(models.UserModel)

	self.Ctx.Input.Bind(&mUserModel.Name, "name")
	self.Ctx.Input.Bind(&mUserModel.Pass, "pass")

	var res Commons.ResponseModel

	if len(mUserModel.Name) < 0 || len(mUserModel.Pass) < 0 {

		res.Failed("账号或密码不能为空")
	}


	if mUserModel.UserIsExist() {
		res.Failed("账号存在, 可以使用该账号直接登录")
	} else {
		if mUserModel.InsetUser() { // success
			logs.Debug("注册成功")
			bankModel := new(models.BankModel)
			bankModel.UserId = mUserModel.UserId
			bankModel.InsertBalance()
			res.Success(map[string]string{"userId": mUserModel.UserId})
		} else {
			res.Failed("注册失败, 清稍后再试")
		}
	}
	self.Data["json"] = &res
	self.ServeJSON()
}

func (self *UserCenterController) Reset() {

	mUserModel := new(models.UserModel)

	self.Ctx.Input.Bind(&mUserModel.UserId, "userId")
	self.Ctx.Input.Bind(&mUserModel.Pass, "pass")

	var res Commons.ResponseModel

	if mUserModel.UserIdIsExist() {
		if mUserModel.ResetPassword() {
			res.Success(nil)
		} else {
			res.Failed("修改失败, 请稍后再试")
		}
	} else {
		res.Failed("账号不存在")
	}
	self.Data["json"] = &res
	self.ServeJSON()
}



