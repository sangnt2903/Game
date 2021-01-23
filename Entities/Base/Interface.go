package Base

import (
	"gorm.io/gorm"
)

type InterfaceBase interface {
	CreateInstance(instance InterfaceModel) error
	UpdateInstance(instance InterfaceModel) error
	GetInstanceByID(instance InterfaceModel, uuid string) error
	GetInstancesWithConditions(instances interface{}, where string, args ...interface{}) error
	DeleteInstanceWithConditions(instance InterfaceModel, where string, args ...interface{}) error
	AddStatement(where string, args ...interface{}) *gorm.DB
}
