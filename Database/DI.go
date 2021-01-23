package Database

import (
	. "MiniGameAPI/Entities/Base"
	. "MiniGameAPI/Entities/Character"
	. "MiniGameAPI/Entities/Server"
	"github.com/gin-gonic/gin"
)

func InitializeDI(r *gin.Engine) {
	db := GetDatabase()

	// base
	baseRepo := NewBaseRepository(db)
	//base router
	apiV1 := r.Group("/api/v1/")

	// server
	serverCtl := NewServerController(NewServerService(NewServerRepository(baseRepo)))
	MakeServerHandlerFunc(apiV1, serverCtl)

	// server
	characterCtl := NewCharacterController(NewCharacterService(NewCharacterRepository(baseRepo)))
	MakeCharacterHandlerFunc(apiV1, characterCtl)
}
