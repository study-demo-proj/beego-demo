package controllers

import (
	"github.com/study-demo-proj/beego-demo/common/conntrollers"
)

type DemoController struct {
	conntrollers.CommonContro
}

const (
	Ping="GET:Ping"
)

func (c *DemoController) Ping() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
