package api

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string
	Description string
	Default bool
	Color string
	OrganizationID int
	// TODO search vector
}

type GetAllProjectsRequestOptions struct {
	Paingation int
}


type CreateProjectRequestOptions struct {
	Name string
	Description string
}

type UpdateProjectRequestOptions struct {
	gorm.Model
	Name string
	Description string
	Default bool
	Color string
	OrganizationID int
}

