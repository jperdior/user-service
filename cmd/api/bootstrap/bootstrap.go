package bootstrap

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"time"
	"user-service/internal/platform/bus/inmemory"
	"user-service/internal/platform/database/mysql"
	"user-service/internal/platform/server"
	"user-service/internal/user/application"
	"user-service/internal/user/infrastructure"
)

func Run() error {

	var cfg config
	err := envconfig.Process("pooling", &cfg)
	if err != nil {
		return err
	}

	var (
		commandBus     = inmemory.NewCommandBus()
		queryBus       = inmemory.NewQueryBus()
		eventBus       = inmemory.NewEventBus()
		userRepository = infrastructure.NewUserRepository(mysql.DB)
		userService    = application.NewUserService(userRepository)
	)

	ctx, srv := server.New(
		context.Background(),
		cfg.Host,
		cfg.Port,
		cfg.ShutdownTimeout,
		commandBus,
		queryBus,
		eventBus,
		userService)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:""`
	Port            uint          `default:"9091"`
	ShutdownTimeout time.Duration `default:"10s"`
}
