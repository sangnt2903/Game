package Specification

import (
	"MiniGameAPI/Commons/constants"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Specification struct {
	UUID      string         `json:"uuid" gorm:"primaryKey"`
	ItemUUID  string         `json:"item_uuid"`
	Attack    int32          `json:"attack"`
	Defense   int32          `json:"defense"`
	Avoid     int32          `json:"avoid"`
	Lucky     int32          `json:"lucky"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (s *Specification) GetUUID() string {
	return s.UUID
}

func (s *Specification) SetUUID() {
	id, _ := uuid.NewV4()
	s.UUID = id.String()
}

func (s *Specification) TableName() string {
	return constants.Specifications
}
