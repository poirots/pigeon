package router

import (
	"github.com/gin-gonic/gin"
	"pigeon/api/v1"
)

func InitUserApi(r *gin.Engine) {
	group := r.Group("")
	{
		//group.GET("/:uid", v1.GetUserDetail)
		//group.GET("f/:username", v1.FindUser)

		group.GET("/users", v1.ListAllUser)
		group.GET("/user/:uuid", v1.GetUserDetail)
		group.GET("/user", v1.GetUserList)
		group.POST("/user/register", v1.Register)
		group.POST("/user/login", v1.Login)
		group.PUT("/user", v1.ModifyUserInfo)

		// group.GET("/user/name", v1.GetUserOrGroupByName)
		//group.PUT("/user", v1.ModifyUserInfo)
		//
		group.POST("/friend", v1.AddFriend)
		//
		//group.GET("/message", v1.GetMessage)
		//
		//group.GET("/file/:fileName", v1.GetFile)
		group.POST("/file", v1.SaveFile)
		//
		group.GET("/group/:uuid", v1.GetGroup)
		group.POST("/group/:uuid", v1.SaveGroup)
		//group.POST("/group/join/:userUuid/:groupUuid", v1.JoinGroup)
		//group.GET("/group/user/:uuid", v1.GetGroupUsers)

	}

}
