package Item

import (
	"MiniGameAPI/Commons/constants"
	Specification2 "MiniGameAPI/Entities/Specification"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Item struct {
	UUID          string                       `json:"uuid" gorm:"primaryKey"`
	Name          string                       `json:"name"`
	Type          string                       `json:"type"`
	RequireLevel  string                       `json:"require_level"`
	IsLocked      string                       `json:"IsLocked"`
	OwnerUUID     string                       `json:"owner_uuid"`
	Specification Specification2.Specification `json:"specification"`
	CreatedAt     time.Time                    `json:"created_at"`
	UpdatedAt     time.Time                    `json:"updated_at"`
	DeletedAt     gorm.DeletedAt               `json:"deleted_at"`
}

func (i *Item) GetUUID() string {
	return i.UUID
}

func (i *Item) SetUUID() {
	id, _ := uuid.NewV4()
	i.UUID = id.String()
}

func (i *Item) TableName() string {
	return constants.Items
}
