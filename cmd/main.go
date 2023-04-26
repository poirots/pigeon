package main

import (
	"flag"
	"fmt"
	"pigeon/common/log"
	"pigeon/config"
	"pigeon/router"
)

func main() {
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.SugarLogger.Debugf("start server %s", config.GetConfig().Log.Path)
	// 注册路由
	r := router.InitRouter()
	r.StaticFile("/favicon.ico", "./favicon.ico")
	port := flag.Int("port", 8080, "服务器端口")
	flag.Parse()
	err := r.Run(fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
