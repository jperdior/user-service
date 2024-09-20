package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/platform/server/handler/status"
	"user-service/internal/user/domain"
)

func setClaimsMiddleware(claims jwt.MapClaims) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("claims", claims)
		c.Next()
	}
}

func TestRoleMiddleware(t *testing.T) {

	t.Run("Token with required roles", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		engine := gin.New()

		claims := jwt.MapClaims{
			"roles": []interface{}{domain.RoleSuperAdmin},
		}
		engine.Use(setClaimsMiddleware(claims))
		engine.Use(RoleMiddleware([]string{domain.RoleSuperAdmin}))
		engine.GET("/status", status.StatusHandler())

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/status", nil)
		engine.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("Token without required roles", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		engine := gin.New()

		claims := jwt.MapClaims{
			"roles": []interface{}{domain.RoleUser},
		}
		engine.Use(setClaimsMiddleware(claims))
		engine.Use(RoleMiddleware([]string{domain.RoleSuperAdmin}))
		engine.GET("/status", status.StatusHandler())

		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/status", nil)
		engine.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusForbidden, recorder.Code)
	})
}
