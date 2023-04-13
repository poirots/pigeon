package main

import (
	"fmt"
	"github.com/pigeon/router"
)

func main() {

	// 注册路由
	r := router.InitRouter()
	r.StaticFile("/favicon.ico", "./favicon.ico")
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
