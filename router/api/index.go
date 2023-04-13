package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pigeon/router/api/userRouter"
)

func InitApi(r *gin.Engine) {
	api := r.Group("/api")
	userRouter.UserInitRouter(api)
}
