package controllers

import (
	"XyBeeGoDemo/Commons"
	"XyBeeGoDemo/models"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseControllerInterface interface {
	CheckRequestMethod(method string)
}

type UserCenterController struct  {
	beego.Controller
}

/* 登录接口 */
func (self *UserCenterController) Login() {

	mUserModel := new(models.UserModel)

	self.Ctx.Input.Bind(&mUserModel.Name, "Name")
	self.Ctx.Input.Bind(&mUserModel.Pass, "Pass")

	var res Commons.ResponseModel

	if findUser, isExist := models.UserIsExist(mUserModel.Name); isExist == true {
		fmt.Println("exist")
		if findUser.Pass == mUserModel.Pass {
			fmt.Println("login success")
			res.Success(nil)
		} else {
			fmt.Println("login failed")
			res.Failed("密码错误")
		}
	} else {
		fmt.Println("not exist")
		res.Failed("账号不存在")
	}

	self.Data["json"] = &res
	self.ServeJSON()
}

func (self *UserCenterController)Regist() {

	mUserModel := new(models.UserModel)

	self.Ctx.Input.Bind(&mUserModel.Name, "Name")
	self.Ctx.Input.Bind(&mUserModel.Pass, "Pass")

	var res Commons.ResponseModel

	if _, isExist := models.UserIsExist(mUserModel.Name); isExist == true {
		fmt.Println("存在, 可以直接登录")
		res.Failed("账号存在, 可以使用该账号直接登录")
	} else {
		fmt.Println("查如数据库")
		if models.InsetUser(mUserModel) == true {// success
			fmt.Println("注册成功")
			res.Success(nil)
		} else {
			fmt.Println("注册失败")
			res.Failed("注册失败, 清稍后再试")
		}
	}
	self.Data["json"] = &res
	self.ServeJSON()
}

