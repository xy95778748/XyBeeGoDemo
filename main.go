package main

import (
	_ "XyBeeGoDemo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
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
	fmt.Println(1)
}
