package application

import (
	"context"
	"user-service/internal/user/domain"
	"user-service/kit"
)

func GetAuthenticatedUser(ctx context.Context) (*domain.AuthenticatedUser, error) {
	id, idExists := ctx.Value("ID").(string)
	name, nameExists := ctx.Value("name").(string)
	email, emailExists := ctx.Value("email").(string)
	rolesValue := ctx.Value("roles")

	if !idExists || !nameExists || !emailExists || rolesValue == nil {
		return nil, kit.NewDomainError("required fields not found", "user.auth.error")
	}

	var rolesSlice []string
	switch v := rolesValue.(type) {
	case string:
		rolesSlice = []string{v}
	case []interface{}:
		rolesSlice = make([]string, len(v))
		for i, role := range v {
			rolesSlice[i] = role.(string)
		}
	default:
		return nil, kit.NewDomainError("unexpected type for roles", "user.auth.error")
	}

	authenticatedUser := domain.NewAuthenticatedUser(
		id,
		name,
		email,
		rolesSlice,
	)
	return &authenticatedUser, nil
}
