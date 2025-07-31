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

type Machine struct {
	ID               string `json:"id"`
	Name             string `json:"name" binding:"required"`
	IsRepairRequired *bool  `json:"is_repair_required,omitempty"`
	IsActive         *bool  `json:"is_active,omitempty"`
}

type MachineService struct {
	grpcClient *client.Client
	log        *slog.Logger
}

func NewMachineService(grpcClient *client.Client, log *slog.Logger) *MachineService {
	return &MachineService{
		grpcClient: grpcClient,
		log:        log,
	}
}

func (s *MachineService) CreateMachine(c *gin.Context) {
	var req Machine
	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Invalid request format", logger.ErrToAttr(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeout)
	defer cancel()

	resp, err := s.grpcClient.CreateMachine(ctx, &entities.CreateMachine{
		Name:             req.Name,
		IsRepairRequired: req.IsRepairRequired,
		IsActive:         req.IsActive,
	})
	if err != nil {
		s.log.Error("Failed to create machine", logger.ErrToAttr(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create machine"})
		return
	}

	c.JSON(http.StatusCreated, Machine{
		Name:             resp.Data.Name,
		IsRepairRequired: resp.Data.IsRepairRequired,
		IsActive:         resp.Data.IsActive,
	})
}
