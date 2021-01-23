package Server

import (
	"MiniGameAPI/Commons/Paging"
	"MiniGameAPI/Logging/Response"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type ServerController struct {
	ss *ServerService
}

func NewServerController(service *ServerService) *ServerController {
	return &ServerController{service}
}

func (sc *ServerController) CreateServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response.NewResponse(c)
		var req *ServerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		var payload Server
		if err := mapstructure.Decode(req, &payload); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		if err := sc.ss.CreateServer(&payload); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		r.Show(http.StatusOK, req, nil)
	}
}

func (sc *ServerController) GetServers() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response.NewResponse(c)
		var servers []Server
		paginator := Paging.NewPaginator(c)
		if err := sc.ss.GetServers(&servers, paginator); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		paginator.Records = PublicServers(servers)
		r.Show(http.StatusOK, paginator, nil)
	}
}

func (sc *ServerController) GetServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := Response.NewResponse(c)
		uuid := c.Param("uuid")
		var server Server
		if err := sc.ss.GetServer(&server, uuid); err != nil {
			r.Show(http.StatusBadRequest, err, nil)
			return
		}
		r.Show(http.StatusOK, server.Public(), nil)
	}
}
