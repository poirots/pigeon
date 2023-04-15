package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pigeon/common/response"
	"pigeon/internal/service"
)

func GetUserList(c *gin.Context) {
	uuid := c.Query("uuid")
	c.JSON(http.StatusOK, response.Success(service.UserService.GetUserList(uuid)))
}

func GetUserDetail(c *gin.Context) {
	uid := c.Param("uuid")
	c.JSON(http.StatusOK, response.Success(service.GetUserDetail(uid)))
}

func FindUser(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, response.Success(service.FindUserByUserName(username)))
}
