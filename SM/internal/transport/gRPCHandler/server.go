package grpchandler

import (
	"log/slog"

	grpc_storage "github.com/GHFluding/ShiftManager/SM/internal/grpcStorage"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"
)

type gRPCServer struct {
	log  *slog.Logger
	port int
	db   models.DBFunction
}

func New(log *slog.Logger, port int, sp *grpc_storage.Services) *gRPCServer {
	machine := models.MachineDB{
		Saver:    sp,
		Provider: sp,
	}
	user := models.UserDB{
		Saver:    sp,
		Provider: sp,
	}
	task := models.TaskDB{
		Saver:    sp,
		Provider: sp,
	}
	shift := models.ShiftDB{
		Saver:    sp,
		Provider: sp,
	}
	db := models.DBFunction{
		Machine: machine,
		User:    user,
		Shift:   shift,
		Task:    task,
	}
	return &gRPCServer{
		log:  log,
		port: port,
		db:   db,
	}
}
