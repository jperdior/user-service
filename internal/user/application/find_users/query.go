package find_users

import (
	"user-service/kit"
	"user-service/kit/query"
)

const FindUsersQueryType query.Type = "find_users"

type FindUsersQuery struct {
	ID       string
	Email    string
	Page     int
	PageSize int
}

func NewFindUserQuery(id, email string, page, pageSize int) FindUsersQuery {
	return FindUsersQuery{
		ID:       id,
		Email:    email,
		Page:     page,
		PageSize: pageSize,
	}
}

func (c FindUsersQuery) Type() query.Type {
	return FindUsersQueryType
}

type FindUsersQueryHandler struct {
	service FindUsersService
}

func NewFindUsersQueryHandler(service FindUsersService) FindUsersQueryHandler {
	return FindUsersQueryHandler{service: service}
}

// Handle implements the query.Handler interface
func (h FindUsersQueryHandler) Handle(findUserQuery query.Query) (interface{}, *kit.DomainError) {
	fuq, ok := findUserQuery.(FindUsersQuery)
	if !ok {
		return nil, kit.NewDomainError("unexpected query", "user.find_user.error", 500)
	}
	user, err := h.service.FindUsers(fuq.ID, fuq.Email, fuq.Page, fuq.PageSize)
	if err != nil {
		return nil, err
	}
	return user, nil
}
