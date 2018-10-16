package main

import (
	"XyBeeGoDemo/models"
	_ "XyBeeGoDemo/routers"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	beego.SetLogger(logs.AdapterFile, `{"filename":"project.log"}`)
	beego.SetLogFuncCall(true)
	beego.BeeLogger.DelLogger(logs.AdapterConsole)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()

	//beego.Async()
}

func init() {

	registDataBase()
}

func registDataBase () {

	orm.RegisterDataBase("default", "mysql", "root:95778748@tcp(127.0.0.1:3306)/User?charset=utf8", 30)
	orm.RegisterModel(new(models.UserModel))
	orm.RegisterModel(new(models.BankModel))
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
