package vtsb_interface

import "github.com/aaron2198/vts_broker/model"

type User interface {
	// GetAll Users and return the list
	GetAll() ([]*model.User, error)
	// create a new User
	Create(*model.User) error
	// update an existing User
	Update(*model.User) error
	// delete an existing User
	Delete(id uint) error
	// FindByID returns a user found by its ID
	FindByID(id uint) (*model.User, error)
	// FindByEmailFirstLast returns a user matching the provided email and first and last name
	FindByEmailFirstLast(email, first, last string) (*model.User, error)
	// FindByDbId returns all users matching the provides database id
	FindByDbId(id uint) (*model.User, error)
	// FuzzyFind returns a list of users matching the provided search term as a substring of Email
	FuzzyFind(email string) ([]*model.User, error)
}
