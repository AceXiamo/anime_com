package global

import (
	"github.com/kataras/iris/v12"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func InitGlobalConfig() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	log.Info("Global Config Init Success！")
}

func Result(ctx iris.Context, code int, msg interface{}) {
	if code == 500 {
		log.Error("update failed", msg)
	}
	ctx.JSON(iris.Map{
		"code": code,
		"msg":  msg,
	})
}
