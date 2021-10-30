package auth

import (
	"net/http"
)

const (
	errorMsg string = "Forbidden."
	// todo probably don't need this with the http library
	statusCode string = "HTTP_403_FORBIDDEN"
)

type BasePermission interface{
	HasRequiredPermissions(r http.Request) bool
}

type PermissionsDependency struct {
	PermissionsClasses map[string]string
}

type OrganizationOwnerPermission struct{

}

func (o *OrganizationOwnerPermission) HasRequiredPermissions(r http.Request) bool {
	return false
}


