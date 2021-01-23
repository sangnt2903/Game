package Server

import (
	"MiniGameAPI/Commons/Paging"
)

type ServerServiceInterface interface {
	CreateServer(server *Server) error
	GetServers(servers *[]Server, paginator ...*Paging.Paginator) error
	GetServer(server *Server, uuid string) error
}

type ServerService struct {
	sr *ServerRepository
}

func (ss *ServerService) GetServer(server *Server, uuid string) error {
	return ss.sr.GetInstanceByID(server, uuid)
}

func (ss *ServerService) GetServers(servers *[]Server, paginator ...*Paging.Paginator) error {
	if len(paginator) > 0 {
		return paginator[0].Paging(ss.sr.Db, servers)
	}
	return ss.sr.GetInstancesWithConditions(servers, "")
}

func (ss *ServerService) CreateServer(server *Server) error {
	return ss.sr.CreateInstance(server)
}

var _ ServerServiceInterface = &ServerService{}

func NewServerService(base *ServerRepository) *ServerService {
	return &ServerService{base}
}
