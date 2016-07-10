package main

import (
	_ "USC-HA-User/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("console", `{"level":7}`) // debug 级别以上的信息会被输出
	beego.SetLogger("file", `{"filename":"logs/`+beego.AppConfig.String("appname")+`.log"}`)
	beego.Run()
}
