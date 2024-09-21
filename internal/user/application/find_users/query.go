package find_user

import (
	"user-service/kit"
	"user-service/kit/query"
)

const FindUsersQueryType query.Type = "find_users"

type FindUserQuery struct {
	ID       string
	Email    string
	Page     int
	PageSize int
}

func NewFindUserQuery(id, email string, page, pageSize int) FindUserQuery {
	return FindUserQuery{
		ID:       id,
		Email:    email,
		Page:     page,
		PageSize: pageSize,
	}
}

func (c FindUserQuery) Type() query.Type {
	return FindUsersQueryType
}

type FindUsersQueryHandler struct {
	service FindUsersService
}

func NewFindUserQueryHandler(service FindUsersService) FindUserQueryHandler {
	return FindUsersQueryHandler{service: service}
}

// Handle implements the query.Handler interface
func (h FindUsersQueryHandler) Handle(findUserQuery query.Query) (interface{}, *kit.DomainError) {
	fuq, ok := findUserQuery.(FindUserQuery)
	if !ok {
		return nil, kit.NewDomainError("unexpected query", "user.find_user.error", 500)
	}
	user, err := h.service.FindUser(fuq.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
