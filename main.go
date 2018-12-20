package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/setting"
	"github.com/jweboy/api-server/router"
)

func main() {
	// 初始化项目基础配置
	setting.Setup()
	// 初始化数据库连接
	model.DB.Init()
	defer model.DB.Close()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)

	server := endless.NewServer(endPoint, router.Load())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	// 初始化gin
	// runmode := cfg.Section("").Key("app_mode").String()
	// gin.SetMode(runmode)

	// g := gin.New()
	// // middlewares := []gin.HandlerFunc{}
	// server := cfg.Section("server")
	// port := server.Key("port").String()

	// http.ListenAndServe(port, g).Error()
	// fmt.Println()

	// fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
}
