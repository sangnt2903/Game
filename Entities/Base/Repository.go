package Base

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RepositoryBase struct {
	Db *gorm.DB
}

func (r RepositoryBase) CreateInstance(instance InterfaceModel) error {
	instance.SetUUID()
	return r.Db.Create(instance).Error
}

func (r RepositoryBase) UpdateInstance(instance InterfaceModel) error {
	return r.Db.Updates(instance).Error
}

func (r RepositoryBase) GetInstanceByID(instance InterfaceModel, uuid string) error {
	return r.Db.Preload(clause.Associations).First(instance, "uuid = ?", uuid).Error
}

func (r RepositoryBase) GetInstancesWithConditions(instances interface{}, where string, args ...interface{}) error {
	return r.Db.Preload(clause.Associations).Where(where, args...).Find(instances).Error
}

func (r RepositoryBase) DeleteInstanceWithConditions(instance InterfaceModel, where string, args ...interface{}) error {
	return r.Db.Where(where, args...).Delete(instance).Error
}

func (r RepositoryBase) AddStatement(where string, args ...interface{}) (statement *gorm.DB) {
	if statement == nil {
		statement = r.Db.Where(where, args...)
	}
	statement = statement.Where(where, args...)
	return statement
}

var _ InterfaceBase = &RepositoryBase{}

func NewBaseRepository(db *gorm.DB) *RepositoryBase {
	return &RepositoryBase{Db: db}
}
