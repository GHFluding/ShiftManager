package services

import (
	"context"
	"net/http"

	"telegramSM/internal/telegramapi/commands"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService commands.UserService
}

type MasterHandler struct {
	masterService MasterService
}

type MasterService interface {
	ListMasters(ctx context.Context) ([]commands.MasterIcon, error)
}

func NewUserHandler(us commands.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func NewMasterHandler(ms MasterService) *MasterHandler {
	return &MasterHandler{masterService: ms}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var req struct {
		TelegramID int `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), req.TelegramID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) SaveUser(c *gin.Context) {
	var user commands.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.SaveUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *MasterHandler) ListMasters(c *gin.Context) {
	masters, err := h.masterService.ListMasters(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, masters)
}
