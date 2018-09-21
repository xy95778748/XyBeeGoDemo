package controllers

import (
	"XyBeeGoDemo/Commons"
	"XyBeeGoDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
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

	self.Ctx.Input.Bind(&mUserModel.Name, "Name")
	self.Ctx.Input.Bind(&mUserModel.Pass, "Pass")
	pass := mUserModel.Pass

	var res Commons.ResponseModel

	if mUserModel.UserIsExist() {
		beego.Debug("exist")
		if pass == mUserModel.Pass {
			res.Success(nil)
		} else {
			res.Failed("密码错误")
		}
	} else {
		res.Failed("账号不存在")
	}
	fmt.Println(time.Now().Unix())

	self.Data["json"] = &res
	self.ServeJSON()
}

func (self *UserCenterController) Regist() {

	mUserModel := new(models.UserModel)

	self.Ctx.Input.Bind(&mUserModel.Name, "Name")
	self.Ctx.Input.Bind(&mUserModel.Pass, "Pass")

	var res Commons.ResponseModel

	if mUserModel.UserIsExist() {
		res.Failed("账号存在, 可以使用该账号直接登录")
	} else {
		if mUserModel.InsetUser() { // success
			logs.Debug("注册成功")
			res.Success(nil)
		} else {
			res.Failed("注册失败, 清稍后再试")
		}
	}
	self.Data["json"] = &res
	self.ServeJSON()
}

func (self *UserCenterController) Reset() {

	mUserModel := new(models.UserModel)

	self.Ctx.Input.Bind(&mUserModel.Name, "Name")
	self.Ctx.Input.Bind(&mUserModel.Pass, "Pass")

	var res Commons.ResponseModel

	if mUserModel.UserIsExist() {
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


