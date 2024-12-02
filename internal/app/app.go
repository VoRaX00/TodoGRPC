package app

import (
	"log/slog"
	grpcapp "todoGRPC/internal/app/grpc"
	"todoGRPC/internal/services/storage/postgres"
	"todoGRPC/internal/services/tasks"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string) *App {
	_, err := postgres.New(storagePath)
	if err != nil {
		panic(err)
	}

	taskService := tasks.New(log)
	grpcApp := grpcapp.New(log, taskService, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
