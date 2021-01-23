package Character

import (
	"MiniGameAPI/Commons/Paging"
	"MiniGameAPI/Logging/Response"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type CharacterController struct {
	ss *CharacterService
}

func NewCharacterController(service *CharacterService) *CharacterController {
	return &CharacterController{service}
}

func (sc *CharacterController) CreateCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response.NewResponse(c)
		var req *CharacterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		var payload Character
		if err := mapstructure.Decode(req, &payload); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		if err := sc.ss.CreateCharacter(&payload); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		r.Show(http.StatusOK, req, nil)
	}
}

func (sc *CharacterController) GetCharacters() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response.NewResponse(c)
		var Characters []Character
		paginator := Paging.NewPaginator(c)
		if err := sc.ss.GetCharacters(&Characters, paginator); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		paginator.Records = PublicCharacters(Characters)
		r.Show(http.StatusOK, paginator, nil)
	}
}

func (sc *CharacterController) GetCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response.NewResponse(c)
		uuid := c.Param("uuid")
		var Character Character
		if err := sc.ss.GetCharacter(&Character, uuid); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		r.Show(http.StatusOK, Character.Public(), nil)
	}
}
