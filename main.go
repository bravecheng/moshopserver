package main

import (
	"github.com/astaxie/beego"
	"github.com/harlanc/moshopserver/services"

	_ "github.com/harlanc/moshopserver/models"
	_ "github.com/harlanc/moshopserver/routers"
	_ "github.com/harlanc/moshopserver/utils"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 8080

	beego.InsertFilter("/api/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080

}
