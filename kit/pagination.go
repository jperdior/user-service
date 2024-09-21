package kit

import "user-service/internal/user/domain"

// PaginationDTO represents the pagination metadata and the results
type PaginationDTO struct {
	Page       int           `json:"page"`
	PageSize   int           `json:"pageSize"`
	TotalRows  int64         `json:"totalRows"`
	TotalPages int           `json:"totalPages"`
	Data       []domain.User `json:"data"`
}
