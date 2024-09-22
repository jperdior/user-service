package kit

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"user-service/internal/user/domain"
)

func GetAuthenticatedUser(ctx context.Context) (*domain.AuthenticatedUser, error) {
	ginCtx, ok := ctx.Value("ginContext").(*gin.Context)
	if !ok {
		return nil, NewDomainError("invalid context", "user.find_user.error", 500)
	}
	id, idExists := ginCtx.Get("ID")
	roles, rolesExists := ginCtx.Get("roles")
	name, nameExists := ginCtx.Get("name")
	email, emailExists := ginCtx.Get("email")

	if !idExists || !rolesExists || !nameExists || !emailExists {
		return nil, NewDomainError("unauthorized", "user.find_user.error", 401)
	}
	authenticatedUser := domain.NewAuthenticatedUser(
		id.(string),
		name.(string),
		email.(string),
		roles.([]string),
	)
	return &authenticatedUser, nil
}

func UuidStringToBinary(id string) ([]byte, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}
	return parsedUUID.MarshalBinary()
}

func ContainsString(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
