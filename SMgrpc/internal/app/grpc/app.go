package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/machine"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/shift"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/task"
	"github.com/GHFluding/ShiftManager/SMgrpc/internal/grpc/sm/user"
	machine_service "github.com/GHFluding/ShiftManager/SMgrpc/internal/services/machine"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"

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

func New(command CommandCode, log *slog.Logger, db_app models.DBFunction, port int) *App {
	gRPCServer := grpc.NewServer()
	switch command {
	case MachineServer:
		service := machine_service.New(log, db_app.Machine)
		machine.RegisterServerAPI(gRPCServer, service)
	case ShiftServer:
		shift.RegisterServerAPI(gRPCServer, db_app.Shift)
	case UserServer:
		user.RegisterServerAPI(gRPCServer, db_app.User)
	case TaskServer:
		task.RegisterServerAPI(gRPCServer, db_app.Task)
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
