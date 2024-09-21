package find_user

import (
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
	service UserFinderService
}

func NewFindUserQueryHandler(service UserFinderService) FindUserQueryHandler {
	return FindUserQueryHandler{service: service}
}

// Handle implements the query.Handler interface
func (h FindUserQueryHandler) Handle(findUserQuery query.Query) (interface{}, *kit.DomainError) {
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
