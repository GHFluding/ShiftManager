package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
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

type CommandCode int

const (
	machineServer CommandCode = iota
	userServer
	taskServer
	shiftServer
)

var commandName = map[CommandCode]string{
	machineServer: "machine",
	userServer:    "user",
	taskServer:    "task",
	shiftServer:   "shift",
}

func (cs CommandCode) String() string {
	return commandName[cs]
}

func New(command string, log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	switch command {
	case machineServer.String():
		machine.RegisterServerAPI(gRPCServer)
	case shiftServer.String():
		shift.RegisterServerAPI(gRPCServer)
	case userServer.String():
		user.RegisterServerAPI(gRPCServer)
	case taskServer.String():
		task.RegisterServerAPI(gRPCServer)
	}

	return &App{
		log,
		gRPCServer,
		port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("operation", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Starting grpc server", slog.String("addr", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
