package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pigeon/common/log"
	"pigeon/common/response"
	"pigeon/config"
	"pigeon/internal/model"
	"pigeon/internal/service"
	"strings"
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

func ModifyUserInfo(c *gin.Context) {
	var paramUser model.User
	c.ShouldBindJSON(&paramUser)
	log.SugarLogger.Infof("update user: %s", paramUser.Username)
	err := service.UserService.UpdateUser(&paramUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Failed("error"))
	} else {
		c.JSON(http.StatusOK, response.Success("sucess"))
	}
}

func AddFriend(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("sucess"))
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

func SaveFile(c *gin.Context) {
	newId := uuid.New().String()
	userId := c.PostForm("uuid")
	file, _ := c.FormFile("file")
	filename := file.Filename
	idx := strings.LastIndex(filename, ".")
	extName := filename[idx:]
	randomFileName := newId + extName

	log.SugarLogger.Info("file=%s", filename)
	log.SugarLogger.Info("userUuid=%s", userId)

	err := c.SaveUploadedFile(file, config.GetConfig().StaticPath.FilePath+randomFileName)
	if err != nil {
		log.SugarLogger.Error("error", err)
		return
	}
	err2 := service.UserService.UpdateUserAvatar(userId, randomFileName)
	if err2 != nil {
		log.SugarLogger.Error("error", err2)
		return
	}
}

func GetGroup(c *gin.Context) {
	c.PostForm("uuid")
}
