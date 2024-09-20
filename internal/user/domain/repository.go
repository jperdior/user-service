package domain

//go:generate mockery --case=snake --outpkg=domainmocks --output=domainmocks --name=UserRepository

type UserRepository interface {
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(user *User) error
}
