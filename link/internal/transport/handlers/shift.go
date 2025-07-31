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

type Shift struct {
	ID          string `json:"id"`
	MachineID   int64  `json:"machine_id" binding:"required"`
	ShiftMaster int64  `json:"shift_master" binding:"required"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Status      string `json:"status"`
}

type ShiftService struct {
	grpcClient *client.Client
	log        *slog.Logger
}

func NewShiftService(grpcClient *client.Client, log *slog.Logger) *ShiftService {
	return &ShiftService{
		grpcClient: grpcClient,
		log:        log,
	}
}

func (s *ShiftService) CreateShift(c *gin.Context) {
	var req Shift
	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Invalid request format", logger.ErrToAttr(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeout)
	defer cancel()

	resp, err := s.grpcClient.CreateShift(ctx, &entities.CreateShiftParams{
		MachineId:   req.MachineID,
		ShiftMaster: req.ShiftMaster,
	})
	if err != nil {
		s.log.Error("Failed to create shift", logger.ErrToAttr(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create shift"})
		return
	}

	c.JSON(http.StatusCreated, Shift{
		MachineID:   resp.Data.MachineId,
		ShiftMaster: resp.Data.ShiftMaster,
	})
}
