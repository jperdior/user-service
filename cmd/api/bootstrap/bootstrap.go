package bootstrap

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"time"
	"user-service/internal/platform/bus/inmemory"
	"user-service/internal/platform/database/mysql"
	"user-service/internal/platform/mailer"
	"user-service/internal/platform/server"
	"user-service/internal/platform/token"
	"user-service/internal/user/application/login"
	"user-service/internal/user/application/register"
	"user-service/internal/user/infrastructure"
)

func Run() error {

	var cfg config
	err := envconfig.Process("user", &cfg)
	if err != nil {
		return err
	}

	mysql.ConnectDB(mysql.DatabaseConfig{
		User:     cfg.DatabaseUser,
		Password: cfg.DatabasePassword,
		Host:     cfg.DatabaseHost,
		Port:     cfg.DatabasePort,
		Name:     cfg.DatabaseName,
	})

	mailer.NewMailer(mailer.MailerConfig{
		Host:     cfg.MailerHost,
		Port:     cfg.MailerPort,
		User:     cfg.MailerUser,
		Password: cfg.MailerPassword,
	})

	var (
		commandBus      = inmemory.NewCommandBus()
		queryBus        = inmemory.NewQueryBus()
		eventBus        = inmemory.NewEventBus()
		tokenService    = token.NewJwtService(cfg.JwtSecret, cfg.JwtExpiration)
		userRepository  = infrastructure.NewUserRepository(mysql.DB)
		registerService = register.NewUserRegisterService(userRepository)
		loginService    = login.NewUserLoginService(userRepository, tokenService)
	)

	ctx, srv := server.New(
		context.Background(),
		server.ServerConfig{
			Host:            cfg.Host,
			Port:            cfg.Port,
			ShutdownTimeout: cfg.ShutdownTimeout,
		},
		commandBus,
		queryBus,
		eventBus,
		registerService,
		loginService,
	)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:""`
	Port            uint          `default:"9091"`
	ShutdownTimeout time.Duration `default:"10s"`

	// JWT configuration
	JwtSecret     string `required:"true"`
	JwtExpiration int    `default:"15"`

	// Database configuration
	DatabaseUser     string `required:"true"`
	DatabasePassword string `required:"true"`
	DatabaseHost     string `required:"true"`
	DatabasePort     int    `required:"true"`
	DatabaseName     string `required:"true"`

	// Mailer configuration
	MailerHost     string `required:"true"`
	MailerPort     int    `required:"true"`
	MailerUser     string `required:"true"`
	MailerPassword string `required:"true"`
}
