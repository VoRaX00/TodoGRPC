package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"todoGRPC/internal/app"
	"todoGRPC/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	storagePath := fmt.Sprintf(
		`host=%s port=%d dbname=%s user=%s password=%s sslmode=%s`,
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Database,
		cfg.DB.User, cfg.DB.Password, cfg.DB.SSLMode,
	)

	log.Info("starting todo", slog.String("env", cfg.Env))
	application := app.New(log, cfg.GRPC.Port, storagePath)
	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return log
}
