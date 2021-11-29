package main

import (
	"anime_com/connections"
	"anime_com/global"
	"anime_com/router"
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// 初始化全局配置
	global.InitGlobalConfig()

	// init mysql
	err := connections.InitMysql()
	if err != nil {
		log.Info(err)
		os.Exit(1)
	}

	app := iris.New()
	app.ConfigureHost()
	// init router
	router.InitRouter(app)
	// listen port
	app.Listen(":7070")
}
