package token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"user-service/internal/user/domain"
)

type JWTService struct {
	secretKey  string
	expiration int
}

func NewJwtService(secretKey string, expiration int) *JWTService {
	return &JWTService{secretKey: secretKey, expiration: expiration}
}

func (s *JWTService) GenerateToken(user *domain.User) (string, error) {
	duration := time.Duration(s.expiration) * time.Hour
	claims := jwt.MapClaims{
		"uid":   user.GetID(),
		"roles": user.Roles,
		"iss":   "user-service",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}
