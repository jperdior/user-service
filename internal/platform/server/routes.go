package server

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"user-service/internal/platform/server/handler/status"
	"user-service/internal/user/presentation"
)

// @title           Golang Template API
// @version         1.0
// @description     This is a Golang template API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://julioperdiguer.es
// @contact.email  julio.perdiguer@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9091
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func (s *Server) registerRoutes() {
	api := s.engine.Group("/api/v1")
	{
		api.GET("/status", status.StatusHandler())
		api.POST("/register", presentation.RegisterUserHandler(s.userService))
	}

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
