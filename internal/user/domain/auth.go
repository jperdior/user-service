package domain

import "user-service/kit"

type AuthenticatedUser struct {
	ID    string
	Name  string
	Email string
	Roles []string
}

func NewAuthenticatedUser(id, name, email string, roles []string) AuthenticatedUser {
	return AuthenticatedUser{
		ID:    id,
		Name:  name,
		Email: email,
		Roles: roles,
	}
}

func (u AuthenticatedUser) IsSuperAdmin() bool {
	return kit.ContainsString(u.Roles, RoleSuperAdmin)
}
