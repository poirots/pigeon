package userRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserInitRouter 初始化模块路由
func UserInitRouter(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "user list")
	})
	userAuth := r.Group("/user-auth")
	UserAuthRouter(userAuth)
}
