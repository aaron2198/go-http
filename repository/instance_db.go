package repo

import (
	"github.com/aaron2198/vts_broker/model"
	"gorm.io/gorm"
)

type InstanceDb struct {
	Db *gorm.DB
}

func (r *InstanceDb) GetAll() ([]*model.InstanceDb, error) {
	var instances []*model.InstanceDb
	err := r.Db.Find(&instances).Error
	return instances, err
}

func (r *InstanceDb) GetById(id uint) (*model.InstanceDb, error) {
	var instance model.InstanceDb
	err := r.Db.First(&instance, id).Error
	return &instance, err
}

func (r *InstanceDb) Create(instance *model.InstanceDb) error {
	err := r.Db.Create(instance).Error
	return err
}

func (r *InstanceDb) Update(instance *model.InstanceDb) error {
	err := r.Db.Save(instance).Error
	return err
}

func (r *InstanceDb) Delete(id uint) error {
	err := r.Db.Delete(&model.InstanceDb{}, id).Error
	return err
}
