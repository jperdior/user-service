package server

import (
	"user-service/internal/platform/server/handler/status"
	"user-service/internal/platform/server/middleware/auth"
	"user-service/internal/user/domain"
	"user-service/internal/user/presentation"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           User Service API
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

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

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
	protected.GET("/users/:uuid", presentation.GetUserHandler(s.queryBus))
	protected.PUT("/users/:uuid", presentation.UpdateUserHandler(s.updateUserService))

	adminProtected := protected.Group("")
	adminProtected.Use(auth.RoleMiddleware([]string{domain.RoleSuperAdmin}))
	adminProtected.GET("/users", presentation.GetUsersHandler(s.queryBus))

	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
