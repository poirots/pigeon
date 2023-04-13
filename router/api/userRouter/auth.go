package userRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/pigeon/config"
	"net/http"
)

func UserAuthRouter(g *gin.RouterGroup) {
	g.GET("/login", func(c *gin.Context) {
		//name := c.Param("name") //通过Context的Param方法来获取API参数
		//action := c.Param("action")
		////截取
		//action = strings.Trim(action, "/")
		c.JSON(http.StatusOK, gin.H{
			"message": config.GetConfig().MySQL.Host,
		})
	})
}
