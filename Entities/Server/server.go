package Server

import (
	"MiniGameAPI/Commons/constants"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Server struct {
	UUID      string         `json:"uuid" gorm:"primaryKey"`
	Host      string         `json:"host"`
	Port      int            `json:"port"`
	Name      string         `json:"name"`
	UserName  string         `json:"user_name"`
	Password  string         `json:"password"`
	DBName    string         `json:"db_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (s *Server) GetUUID() string {
	return s.UUID
}

func (s *Server) SetUUID() {
	id, _ := uuid.NewV4()
	s.UUID = id.String()
}

func (s *Server) TableName() string {
	return constants.Servers
}
