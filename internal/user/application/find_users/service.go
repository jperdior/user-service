package find_users

import (
	"user-service/internal/user/application/dto"
	"user-service/internal/user/domain"
	"user-service/kit"
	"user-service/kit/model"
)

type FindUsersService struct {
	repo domain.UserRepository
}

// NewFindUsersService creates a new service for finding users
func NewFindUsersService(repo domain.UserRepository) *FindUsersService {
	return &FindUsersService{repo: repo}
}

// FindUsers fetches users based on the provided query (which includes pagination)
func (s *FindUsersService) FindUsers(
	ID, name, email, role string, page, pageSize int, sort, sortDir string,
) (model.PaginationDTO, error) {
	pageValueObject, err := model.NewPageValueObject(page)
	if err != nil {
		return model.PaginationDTO{}, err
	}
	pageSizeValueObject, err := model.NewPageSizeValueObject(pageSize)
	if err != nil {
		return model.PaginationDTO{}, err
	}
	sortDirValueObject, err := model.NewSortDirValueObject(sortDir)
	if err != nil {
		return model.PaginationDTO{}, err
	}

	filters := map[string]interface{}{}
	if ID != "" {
		filters["id"] = ID
	}
	if name != "" {
		filters["name"] = name
	}
	if email != "" {
		filters["email"] = email
	}
	if role != "" {
		filters["role"] = role
	}

	criteria := domain.NewCriteria(filters, sort, sortDirValueObject.Value(), pageValueObject.Value(), pageSizeValueObject.Value())

	users, totalRows, err := s.repo.Find(criteria)
	if err != nil {
		return model.PaginationDTO{}, kit.NewDomainError(err.Error(), "user.find_users.error", 500)
	}

	// Calculate total pages
	totalPages := int((totalRows + int64(pageSizeValueObject.Value()) - 1) / int64(pageSizeValueObject.Value()))

	// Build pagination DTO
	pagination := model.PaginationDTO{
		Page:       pageValueObject.Value(),
		PageSize:   pageSizeValueObject.Value(),
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       mapUsersToDTOs(users),
	}

	return pagination, nil
}

func mapUsersToDTOs(users []*domain.User) []*dto.UserDTO {
	userDTOs := make([]*dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.NewUserDTO(user)
	}
	return userDTOs
}
