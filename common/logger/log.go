package logger

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/study-demo-proj/beego-demo/common/models"
	"github.com/study-demo-proj/beego-demo/common/tool"
	"golang.org/x/net/context"
	"os"
	"runtime/debug"
)

var logger seelog.LoggerInterface

func init() {
	var err error
	logger, err = seelog.LoggerFromConfigAsFile("./conf/seelog.xml")
	if err != nil {
		debug.PrintStack()
		os.Exit(3)
	}
	seelog.ReplaceLogger(logger)
}

// DisableLog disables all library log output
func DisableLog() {
	logger = seelog.Disabled
}

// Call this before app shutdown
func FlushLog() {
	logger.Flush()
}

//日志规范
func LogFormatInfo(ctx context.Context, msg string) {
	//msg = tool.ConvertNewLineToSpace(msg)
	logger.Info(logFormat(ctx, msg))
}
func LogFormatDebug(ctx context.Context, msg string) {
	//msg = tool.ConvertNewLineToSpace(msg)
	logger.Debug(logFormat(ctx, msg))
}
func LogFormatError(ctx context.Context, msg string) {
	//msg = tool.ConvertNewLineToSpace(msg)
	_ = logger.Error(logFormat(ctx, msg))
}
func LogFormatWarn(ctx context.Context, msg string) {
	//msg = tool.ConvertNewLineToSpace(msg)
	_ = logger.Warn(logFormat(ctx, msg))
}
func LogFormatCritical(ctx context.Context, msg string) {
	//msg = tool.ConvertNewLineToSpace(msg)
	_ = logger.Critical(logFormat(ctx, msg))
}
func logFormat(ctx context.Context, msg string) string {
	msg = tool.ConvertNewLineToSpace(msg)
	info,ok:=ctx.Value("common_info").(*models.CommonInfo)
	if !ok{return fmt.Sprintf("%s",msg)}
	//时间戳|ip|user|msg
	return fmt.Sprintf("%s|%s|%s", info.RemoteIP, info.Username, msg)
}
func Get_logger() seelog.LoggerInterface {
	return logger
}
