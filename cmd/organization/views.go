package organization


type Organization struct {
	Id uint
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