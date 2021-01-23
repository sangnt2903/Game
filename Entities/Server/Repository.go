package Server

import (
	"MiniGameAPI/Entities/Base"
)

type ServerRepository struct {
	*Base.RepositoryBase
}

func NewServerRepository(base *Base.RepositoryBase) *ServerRepository {
	return &ServerRepository{base}
}
