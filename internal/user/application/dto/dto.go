package dto

import "user-service/internal/user/domain"

type UserDTO struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func NewUserDTO(user *domain.User) *UserDTO {
	return &UserDTO{
		ID:        user.GetID(),
		Name:      user.Name,
		Email:     user.Email.Value(),
		Roles:     user.Roles,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}
