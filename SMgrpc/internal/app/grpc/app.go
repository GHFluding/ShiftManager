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

func New(command models.CommandCode, log *slog.Logger, db_app models.DBFunction, port int) *App {
	gRPCServer := grpc.NewServer()
	switch command {
	case models.MachineServer:
		service := machine_service.New(log, db_app.Machine)
		machine.RegisterServerAPI(gRPCServer, service)
	case models.ShiftServer:
		shift.RegisterServerAPI(gRPCServer, db_app.Shift)
	case models.UserServer:
		user.RegisterServerAPI(gRPCServer, db_app.User)
	case models.TaskServer:
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
