package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pigeon/common/response"
	"github.com/pigeon/internal/service"
	"net/http"
)

func InitApi(r *gin.Engine) {
	group := r.Group("/user")
	{
		group.GET("/:uid", GetUserDetail)
		group.GET("f/:username", FindUser)
	}

}

func GetUserDetail(c *gin.Context) {
	uid := c.Param("uid")
	c.JSON(http.StatusOK, response.Success(service.GetUserDetail(uid)))
}

func FindUser(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, response.Success(service.FindUserByUserName(username)))
}
