package app

import (
	"log/slog"
	grpcapp "smgrpc/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, command grpcapp.CommandCode) *App {
	grpcApp := grpcapp.New(command, log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
