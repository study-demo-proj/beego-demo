package conntrollers

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	wsl_log "github.com/study-demo-proj/beego-demo/common/logger"
	"github.com/study-demo-proj/beego-demo/common/models"
)

type CommonContro struct {
	beego.Controller
}

type CommonInfo struct {
	Username string
	RemoteIP string
}

func (this *CommonContro) Context() context.Context {
	//return models.WithValue(this.Ctx.Request.Context(), models.COMMOM_INFO, this.GetCommonInfo())
	return context.WithValue(this.Ctx.Request.Context(),"common_info",CommonInfo{"test","1.1.1.1"})
}

func (this *CommonContro) ResponseBody(code int, msg string, data interface{}) {
	fmt.Println("this is an response : ",msg)
	var r models.ResponseBody
	r.Code = code
	r.Msg = msg
	r.Data = data
	msgLog := fmt.Sprintf("response body:%+v", r)
	if code == models.CODE_SYS_ERROR {
		wsl_log.LogFormatError(this.Context(),msgLog)
	} /*else {
		//wsl_log.LogFormatInfo(this.Context(),msgLog)
	}*/
	if this.Data == nil {
		this.Data = make(map[interface{}]interface{}, 0)
	}
	this.Data["json"] = r
	this.ServeJSON()
	//如果是异常返回则停止往下执行
	if code != 0 {
		this.StopRun()
	}
}