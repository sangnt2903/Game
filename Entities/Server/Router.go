package Server

import "github.com/gin-gonic/gin"

func MakeServerHandlerFunc(rootGroup *gin.RouterGroup, controller *ServerController) {
	servers := rootGroup.Group("/servers")
	{
		servers.POST("/", controller.CreateServer())
		servers.GET("/", controller.GetServers())
		servers.GET("/:uuid", controller.GetServer())
	}
}
