package main

import (
	"github.com/astaxie/beego"
	"github.com/study-demo-proj/beego-demo/common/logger"
	_ "github.com/study-demo-proj/beego-demo/routers"
)

func main() {
	loge := logger.Get_logger()
	loge.Info("The Program is starting...")
	beego.Run()
}

