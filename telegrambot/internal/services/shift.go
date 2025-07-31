package services

import (
	"net/http"

	"telegramSM/internal/telegramapi/commands"

	"github.com/gin-gonic/gin"
)

type ShiftHandler struct {
	shiftService commands.ShiftService
}

func NewShiftHandler(ss commands.ShiftService) *ShiftHandler {
	return &ShiftHandler{shiftService: ss}
}

func (h *ShiftHandler) CreateShift(c *gin.Context) {
	var shift commands.Shift

	if err := c.ShouldBindJSON(&shift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.shiftService.SaveShift(c.Request.Context(), &shift); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shift)
}

func (h *ShiftHandler) ListShifts(c *gin.Context) {
	shifts, err := h.shiftService.ListShifts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shifts)
}
