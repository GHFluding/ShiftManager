package handlers

import (
	"context"
	"net/http"

	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
	"github.com/gin-gonic/gin"
)

type Task struct {
	ID            string `json:"id"`
	MachineID     int64  `json:"machine_id" binding:"required"`
	ShiftID       int64  `json:"shift_id" binding:"required"`
	Frequency     string `json:"frequency" binding:"required"`
	TaskPriority  string `json:"task_priority" binding:"required"`
	Description   string `json:"description" binding:"required"`
	CreatedBy     int64  `json:"created_by" binding:"required"`
	AssignedTo    int64  `json:"assigned_to"`
	Status        string `json:"status"`
	DueDate       string `json:"due_date"`
	CompletedDate string `json:"completed_date"`
}

type TaskService struct {
	grpcClient *client.Client
	log        *slog.Logger
}

func NewTaskService(grpcClient *client.Client, log *slog.Logger) *TaskService {
	return &TaskService{
		grpcClient: grpcClient,
		log:        log,
	}
}

func (s *TaskService) CreateTask(c *gin.Context) {
	var req Task
	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Invalid request format", logger.ErrToAttr(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeout)
	defer cancel()

	resp, err := s.grpcClient.CreateTask(ctx, &entities.CreateTaskParams{
		MachineId:    req.MachineID,
		ShiftId:      req.ShiftID,
		Frequency:    req.Frequency,
		TaskPriority: req.TaskPriority,
		Description:  req.Description,
		CreatedBy:    req.CreatedBy,
	})
	if err != nil {
		s.log.Error("Failed to create task", logger.ErrToAttr(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, Task{
		MachineID:    resp.Data.MachineId,
		ShiftID:      resp.Data.ShiftId,
		Frequency:    resp.Data.Frequency,
		TaskPriority: resp.Data.TaskPriority,
		Description:  resp.Data.Description,
		CreatedBy:    resp.Data.CreatedBy,
	})
}
