package grpcapp

import (
	"log/slog"
	"smgrpc/internal/grpc/sm/machine"
	"smgrpc/internal/grpc/sm/shift"
	"smgrpc/internal/grpc/sm/task"
	"smgrpc/internal/grpc/sm/user"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

const (
	machineServer = "machine"
	userServer    = "user"
	taskServer    = "task"
	shiftServer   = "shift"
)

func New(command string, log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	switch command {
	case machineServer:
		machine.RegisterServerAPI(gRPCServer)
	case shiftServer:
		shift.RegisterServerAPI(gRPCServer)
	case userServer:
		user.RegisterServerAPI(gRPCServer)
	case taskServer:
		task.RegisterServerAPI(gRPCServer)
	}
	return &App{
		log,
		gRPCServer,
		port,
	}
}
