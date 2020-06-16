package routers

import (
	"github.com/astaxie/beego"
	"github.com/study-demo-proj/beego-demo/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
