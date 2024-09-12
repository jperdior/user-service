package domain

//go:generate mockery --case=snake --outpkg=domainmocks --output=domainmocks --name=TokenService

type TokenService interface {
	GenerateToken(user *User) (string, error)
}
