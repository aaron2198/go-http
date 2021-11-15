package repo

import (
	"github.com/aaron2198/vts_broker/model"
	"gorm.io/gorm"
)

type Community struct {
	Db *gorm.DB
}

func (r *Community) GetAll() ([]*model.Community, error) {
	var communities []*model.Community
	err := r.Db.Find(&communities).Error
	return communities, err
}

func (r *Community) Create(community *model.Community) error {
	return r.Db.Create(community).Error
}

func (r *Community) FindByID(id uint) (*model.Community, error) {
	community := &model.Community{}
	if err := r.Db.First(community, id).Error; err != nil {
		return nil, err
	}
	return community, nil
}

func (r *Community) FindBySubdomain(sub string) (*model.Community, error) {
	community := &model.Community{}
	if err := r.Db.Where("subdomain = ?", sub).First(community).Error; err != nil {
		return nil, err
	}
	return community, nil
}

func (r *Community) FindByDbId(id uint) (*model.Community, error) {
	community := &model.Community{}
	if err := r.Db.Where("DatabaseId = ?", id).First(community).Error; err != nil {
		return nil, err
	}
	return community, nil
}

func (r *Community) FuzzyFind(sub string) ([]*model.Community, error) {
	communities := []*model.Community{}
	if err := r.Db.Where("subdomain LIKE ?", sub+"%").Find(&communities).Error; err != nil {
		return nil, err
	}
	return communities, nil
}
