package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user-service/internal/user/application/forgot_password"
	"user-service/internal/user/application/login"
	"user-service/internal/user/application/register"
	"user-service/kit/command"
	"user-service/kit/event"
	"user-service/kit/query"
)

type ServerConfig struct {
	Host            string        `default:""`
	Port            uint          `default:"9091"`
	ShutdownTimeout time.Duration `default:"10s"`
}

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	//deps
	commandBus            command.Bus
	queryBus              query.Bus
	eventBus              event.Bus
	registerService       *register.UserRegisterService
	loginService          *login.UserLoginService
	forgotPasswordService *forgot_password.ForgotPasswordService
}

func New(
	ctx context.Context,
	config ServerConfig,
	commandBus command.Bus,
	queryBus query.Bus,
	eventBus event.Bus,
	registerService *register.UserRegisterService,
	loginService *login.UserLoginService,
	forgotPasswordService *forgot_password.ForgotPasswordService,
) (context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		engine:   gin.Default(),

		shutdownTimeout: config.ShutdownTimeout,

		//deps
		commandBus:            commandBus,
		queryBus:              queryBus,
		eventBus:              eventBus,
		registerService:       registerService,
		loginService:          loginService,
		forgotPasswordService: forgotPasswordService,
	}

	srv.registerRoutes()
	srv.engine.HandleMethodNotAllowed = true
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
