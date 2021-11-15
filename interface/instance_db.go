package vtsb_interface

import "github.com/aaron2198/vts_broker/model"

type InstanceDb interface {
	// GetAll Database Instances and return the list
	GetAll() ([]*model.InstanceDb, error)
	// GetByID Search for Database Instance with the provided id
	GetById(id uint) (*model.InstanceDb, error)
	// Create a new Database Instance
	Create(*model.InstanceDb) error
	// Update an existing Database Instance
	Update(*model.InstanceDb) error
	// Delete an existing Database Instance
	Delete(id uint) error
}
