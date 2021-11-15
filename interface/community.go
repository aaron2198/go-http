package vtsb_interface

import "github.com/aaron2198/vts_broker/model"

type Community interface {
	// GetAll returns all communities
	GetAll() ([]*model.Community, error)
	// Create a new community
	Create(*model.Community) error
	// FindByID returns a community found by its ID
	FindByID(id uint) (*model.Community, error)
	// FindBySubdomain returns a community matching the provided subdomain
	FindBySubdomain(sub string) (*model.Community, error)
	// FindByDbId returns all communities matching the provided Database ID
	FindByDbId(id uint) (*model.Community, error)
	// FuzzyFind returns a list of communities matching the provided search term as sunstring of Subdomain
	FuzzyFind(sub string) ([]*model.Community, error)
}
