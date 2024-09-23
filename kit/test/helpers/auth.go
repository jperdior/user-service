package helpers

import (
	"github.com/google/uuid"
	"user-service/internal/platform/token"
	"user-service/internal/user/domain"
	"user-service/kit/model"
)

// TODO, allow to specify uid
func GenerateJwtToken(roles []string, secretKey string) (string, error) {
	uid := uuid.New()
	baseUser := model.Base{
		ID: uid[:],
	}
	authenticatedUser := domain.User{
		Base:  baseUser,
		Name:  "Tester",
		Email: "tester@federation.com",
		Roles: roles,
	}
	tokenService := token.NewJwtService(secretKey, 1)
	jwtToken, err := tokenService.GenerateToken(&authenticatedUser)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
