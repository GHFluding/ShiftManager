package task

import (
	"context"

	entities "github.com/GHFluding/ShiftManager/SMgrpc/internal/gen"
	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/domain/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskResponse struct {
	MachineId    int64
	ShiftId      int64
	Frequency    string
	TaskPriority string
	Description  string
}

type serverAPI struct {
	entities.UnimplementedTaskServiceServer
	task models.TaskDB
}

func RegisterServerAPI(gRPC *grpc.Server, task models.TaskDB) {
	entities.RegisterTaskServiceServer(gRPC, &serverAPI{task: task})
}
func (s *serverAPI) Create(ctx context.Context, req *entities.CreateTaskParams) (*entities.TaskResponse, error) {
	if req.GetDescription() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if req.GetFrequency() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	if req.GetTaskPriority() == "" {
		return nil, status.Error(codes.InvalidArgument, "name is empty")
	}
	taskID, err := s.task.Saver.SaveTask(ctx, req.MachineId, req.ShiftId, req.Frequency, req.TaskPriority, req.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	task, err := s.task.Provider.GetTask(ctx, taskID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &entities.TaskResponse{
		Data: &entities.CreateTaskParams{
			MachineId:    task.MachineId,
			ShiftId:      task.ShiftId,
			Frequency:    task.Frequency,
			TaskPriority: task.TaskPriority,
			Description:  task.Description,
		},
	}, nil
}
