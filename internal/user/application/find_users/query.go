package find_users

import (
	"context"
	"user-service/kit"
	"user-service/kit/query"
)

const FindUsersQueryType query.Type = "find_users"

type FindUsersQuery struct {
	ID       string
	Name     string
	Email    string
	Role     string
	Page     int
	PageSize int
	Sort     string
	SortDir  string
}

func NewFindUserQuery(id, name, email, role, sort, sortDir string, page, pageSize int) FindUsersQuery {
	return FindUsersQuery{
		ID:       id,
		Name:     name,
		Email:    email,
		Role:     role,
		Page:     page,
		PageSize: pageSize,
		Sort:     sort,
		SortDir:  sortDir,
	}
}

func (c FindUsersQuery) Type() query.Type {
	return FindUsersQueryType
}

type FindUsersQueryHandler struct {
	service *FindUsersService
}

func NewFindUsersQueryHandler(service *FindUsersService) FindUsersQueryHandler {
	return FindUsersQueryHandler{service: service}
}

// Handle implements the query.Handler interface
func (h FindUsersQueryHandler) Handle(ctx context.Context, findUserQuery query.Query) (interface{}, error) {
	fuq, ok := findUserQuery.(FindUsersQuery)
	if !ok {
		return nil, kit.NewDomainError("unexpected query", "user.find_user.error")
	}
	user, err := h.service.FindUsers(fuq.ID, fuq.Name, fuq.Email, fuq.Role, fuq.Page, fuq.PageSize, fuq.Sort, fuq.SortDir)
	if err != nil {
		return nil, err
	}
	return user, nil
}
