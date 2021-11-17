package repo

import (
	"github.com/aaron2198/vts_broker/model"
	"gorm.io/gorm"
)

type User struct {
	Db *gorm.DB
}

func (r *User) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := r.Db.Find(&users).Error
	return users, err
}

func (r *User) Create(user *model.User) error {
	return r.Db.Create(user).Error
}

// func (r *User) Update(id uint, user *model.User) error {
// 	if err := r.Db.Model(&model.User{}).Where("ID = ?", id).Updates(user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
func (r *User) Update(user *model.User) error {
	if err := r.Db.Model(&model.User{}).Where("ID = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
func (r *User) Delete(id uint) error {
	if err := r.Db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *User) FindByID(id uint) (*model.User, error) {
	user := &model.User{}
	if err := r.Db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *User) FindByEmailFirstLast(email, first, last string) (*model.User, error) {
	user := &model.User{}
	if err := r.Db.Where("email = ? AND first_name = ? AND last_name = ?", email, first, last).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindByDbId returns all users matching the provides database id
func (r *User) FindByDbId(id uint) (*model.User, error) {
	user := &model.User{}
	if err := r.Db.Where("DatabaseId = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *User) FuzzyFind(email string) ([]*model.User, error) {
	users := []*model.User{}
	if err := r.Db.Where("email LIKE ?", email+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
