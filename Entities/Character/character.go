package Character

import (
	"MiniGameAPI/Commons/constants"
	. "MiniGameAPI/Entities/Item"
	. "MiniGameAPI/Entities/Server"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Character struct {
	UUID           string         `json:"uuid" gorm:"primaryKey"`
	IdentifierName string         `json:"identifier_name"`
	ServerUUID     string         `json:"server_uuid"`
	Server         Server         `json:"server" gorm:"foreignKey:ServerUUID"`
	Level          int            `json:"level"`
	Items          []Item         `json:"items" gorm:"foreignKey:OwnerUUID"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

func (c *Character) GetUUID() string {
	return c.UUID
}

func (c *Character) SetUUID() {
	id, _ := uuid.NewV4()
	c.UUID = id.String()
}

func (c *Character) TableName() string {
	return constants.Characters
}
