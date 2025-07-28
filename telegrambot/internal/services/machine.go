package services

import (
	"net/http"

	"telegramSM/internal/telegramapi/commands"

	"github.com/gin-gonic/gin"
)

type MachineHandler struct {
	machineService commands.MachineService
}

func NewMachineHandler(ms commands.MachineService) *MachineHandler {
	return &MachineHandler{machineService: ms}
}

func (h *MachineHandler) ListMachines(c *gin.Context) {
	machines, err := h.machineService.ListMachines(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, machines)
}
