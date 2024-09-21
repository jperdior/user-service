package server

import (
	"user-service/internal/platform/server/handler/status"
	"user-service/internal/platform/server/middleware/auth"
	"user-service/internal/user/presentation"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @securityDefinitions.bearerAuth  BearerAuth
// @securityDefinitions.bearerAuth.type  apiKey
// @securityDefinitions.bearerAuth.name  Authorization
// @securityDefinitions.bearerAuth.in    header

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func (s *Server) registerRoutes() {
	api := s.engine.Group("/api/v1")
	{
		api.GET("/status", status.StatusHandler())
		api.POST("/register", presentation.RegisterUserHandler(s.registerService))
		api.POST("/login", presentation.LoginUserHandler(s.loginService))
		api.POST("/forgot-password", presentation.ForgotPasswordHandler(s.forgotPasswordService))
	}

	protected := api.Group("")
	protected.Use(auth.JWTMiddleware(s.config.JwtSecret))
	protected.GET("/user/:uuid", presentation.GetUserHandler(s.queryBus))

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
