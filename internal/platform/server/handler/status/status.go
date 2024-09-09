package status

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Status endpoint
// @Summary Provides the status of the service
// @Schemes
// @Description Provides the status of the service
// @Success 200 {string} OK
// @Router /status [get]
func StatusHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	}
}
