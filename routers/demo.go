package routers

import (
	"github.com/astaxie/beego"
	"github.com/study-demo-proj/beego-demo/controllers"
)

func init() {
	demoContro := &controllers.DemoController{}
	beego.Router("/ping",demoContro,controllers.Ping)
}
