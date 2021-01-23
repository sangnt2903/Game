package Character

import "github.com/gin-gonic/gin"

func MakeCharacterHandlerFunc(rootGroup *gin.RouterGroup, controller *CharacterController) {
	characters := rootGroup.Group("/characters")
	{
		characters.POST("/", controller.CreateCharacter())
		characters.GET("/", controller.GetCharacters())
		characters.GET("/:uuid", controller.GetCharacter())
	}
}
