package application

import (
	"context"
	"github.com/gin-gonic/gin"
	"user-service/internal/user/domain"
	"user-service/kit"
)

func GetAuthenticatedUser(ctx context.Context) (*domain.AuthenticatedUser, error) {
	ginCtx, ok := ctx.Value("ginContext").(*gin.Context)
	if !ok {
		return nil, kit.NewDomainError("invalid context", "user.find_user.error", 500)
	}
	id, idExists := ginCtx.Get("ID")
	roles, rolesExists := ginCtx.Get("roles")
	name, nameExists := ginCtx.Get("name")
	email, emailExists := ginCtx.Get("email")

	if !idExists || !rolesExists || !nameExists || !emailExists {
		return nil, kit.NewDomainError("unauthorized", "user.find_user.error", 401)
	}
	authenticatedUser := domain.NewAuthenticatedUser(
		id.(string),
		name.(string),
		email.(string),
		roles.([]string),
	)
	return &authenticatedUser, nil
}
