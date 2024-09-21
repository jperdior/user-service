package find_users

import (
	"errors"
	"user-service/internal/user/domain"
)

type FindUsersService struct {
	repo domain.UserRepository
}

// NewFindUsersService creates a new service for finding users
func NewFindUsersService(repo domain.UserRepository) *FindUsersService {
	return &FindUsersService{repo: repo}
}

// Execute fetches users based on the provided query (which includes pagination)
func (s *FindUsersService) FindUsers(query FindUsersQuery) (FindUsersResult, error) {
	if query.Page < 1 || query.PageSize < 1 {
		return FindUsersResult{}, errors.New("invalid pagination parameters")
	}

	// Fetch paginated users from repository
	users, totalRows, err := s.repo.FindPaginated(query.Page, query.PageSize)
	if err != nil {
		return FindUsersResult{}, err
	}

	// Calculate total pages
	totalPages := int((totalRows + int64(query.PageSize) - 1) / int64(query.PageSize))

	// Build pagination DTO
	pagination := dto.PaginationDTO{
		Page:       query.Page,
		PageSize:   query.PageSize,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Items:      users,
	}

	// Return result
	return FindUsersResult{
		Users:      users,
		Pagination: pagination,
	}, nil
}
