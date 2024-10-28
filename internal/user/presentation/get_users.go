package presentation

import (
	"net/http"
	"strconv"
	"user-service/internal/user/application/find_users"
	"user-service/kit/query"

	"github.com/gin-gonic/gin"
)

// GetUsersHandler is a handler for getting a user by ID.
// @Summary Get a user by ID
// @Schemes
// @Description Retrieves a user by their UUID
// @Accept json
// @Produce json
// @Param id query string false "User ID"
// @Param name query string false "User Name"
// @Param email query string false "User Email"
// @Param role query string false "User Role"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Param sort query string false "Sort field"
// @Param sortDir query string false "Sort direction"
// @Success 200 {object} model.PaginationDTO "Users pagination"
// @Failure 400 {object} kit.ErrorResponse "Invalid Filters"
// @Failure 500 {object} kit.ErrorResponse "Internal server error"
// @Router /users [get]
// @Tags user
// @Security Bearer
func GetUsersHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")
		email := c.Query("email")
		role := c.Query("role")
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		sort := c.Query("sort")
		sortDir := c.Query("sortDir")

		findUsersQuery := find_users.NewFindUserQuery(
			id, name, email, role, sort, sortDir, page, pageSize)

		userPagination, err := queryBus.Ask(c, findUsersQuery)
		if err != nil {
			MapErrorToHTTP(c, err)
			return
		}
		c.JSON(http.StatusOK, userPagination)
	}
}
