package helpers

import (
	"user-service/internal/platform/token"
	"user-service/internal/user/domain"
	kitDomain "user-service/kit/domain"
)

// TODO, allow to specify uid
func GenerateJwtToken(roles []string, secretKey string) (string, error) {
	authenticatedUser := domain.User{
		ID:    kitDomain.RandomUuidValueObject(),
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
