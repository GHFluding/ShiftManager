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
	MachineServer CommandCode = iota
	UserServer
	TaskServer
	ShiftServer
)

var commandName = map[CommandCode]string{
	MachineServer: "machine",
	UserServer:    "user",
	TaskServer:    "task",
	ShiftServer:   "shift",
}

func (cs CommandCode) String() string {
	return commandName[cs]
}

func New(command CommandCode, log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	switch command {
	case MachineServer:
		machine.RegisterServerAPI(gRPCServer)
	case ShiftServer:
		shift.RegisterServerAPI(gRPCServer)
	case UserServer:
		user.RegisterServerAPI(gRPCServer)
	case TaskServer:
		task.RegisterServerAPI(gRPCServer)
	}

	return &App{
		log,
		gRPCServer,
		port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
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
