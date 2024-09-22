package find_user

import (
	"context"
	"user-service/internal/user/application"
	"user-service/kit"
	"user-service/kit/query"
)

const FindUserQueryType query.Type = "find_user"

type FindUserQuery struct {
	ID string
}

func NewFindUserQuery(ID string) FindUserQuery {
	return FindUserQuery{
		ID: ID,
	}
}

func (c FindUserQuery) Type() query.Type {
	return FindUserQueryType
}

type FindUserQueryHandler struct {
	service *UserFinderService
}

func NewFindUserQueryHandler(service *UserFinderService) FindUserQueryHandler {
	return FindUserQueryHandler{service: service}
}

// Handle implements the query.Handler interface
func (h FindUserQueryHandler) Handle(ctx context.Context, findUserQuery query.Query) (interface{}, error) {
	authenticatedUser, err := application.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}
	fuq, ok := findUserQuery.(FindUserQuery)
	if !ok {
		return nil, kit.NewDomainError("unexpected query", "user.find_user.error", 500)
	}
	user, err := h.service.FindUser(authenticatedUser, fuq.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
