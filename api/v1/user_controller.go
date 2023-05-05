package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pigeon/common/log"
	"pigeon/common/response"
	"pigeon/internal/model"
	"pigeon/internal/service"
)

func Register(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	err := service.UserService.Register(&user)
	if err != nil {
		log.SugarLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.Failed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(nil))
}

func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	log.SugarLogger.Debugf("username=%s", user.Username)
	selectUser := service.UserService.GetUserByUserName(&user)
	if selectUser != nil {
		c.JSON(http.StatusOK, response.Success("success login,uid="+selectUser.Uuid))
	} else {
		c.JSON(http.StatusInternalServerError, response.Failed("error"))
	}
}

func GetUserList(c *gin.Context) {
	uuid := c.Query("uuid")
	c.JSON(http.StatusOK, response.Success(service.UserService.GetUserList(uuid)))
}

func ListAllUser(c *gin.Context) {
	// user := service.UserService.ListAllUser()
	c.JSON(http.StatusOK, response.Success(service.UserService.ListAllUser()))
}

func GetUserDetail(c *gin.Context) {
	uid := c.Param("uuid")
	c.JSON(http.StatusOK, response.Success(service.GetUserDetail(uid)))
}

func FindUser(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, response.Success(service.FindUserByUserName(username)))
}
