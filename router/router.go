package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pigeon/router/api"
	"github.com/pigeon/router/user"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api.InitApi(r)
	user.InitApi(r)
	return r
}
