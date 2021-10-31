package organization

import (
	"github.com/gofrs/uuid"
	// "gorm.io/gorm"
	// "gorm.io/driver/postgres"
)

type Organization struct {
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string
	Default bool
	Description string
	// TODO modify once projects struct is done
	Projects interface{}
}

type OrganizationBase struct {
	Id uint
	Name string
	Description string
	Default bool
}

type OrganizationPagination struct {
	Total uint
	Items []OrganizationBase
}