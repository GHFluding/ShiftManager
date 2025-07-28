package handler

import (
	"net/http"
	"strconv"

	"telegramSM/internal/telegramapi/commands"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService commands.TaskService
}

func NewTaskHandler(ts commands.TaskService) *TaskHandler {
	return &TaskHandler{taskService: ts}
}

func (h *TaskHandler) GetTodayTask(c *gin.Context) {
	telegramIDStr := c.Query("telegram_id")
	if telegramIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "telegram_id parameter is required"})
		return
	}

	telegramID, err := strconv.Atoi(telegramIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "telegram_id must be an integer"})
		return
	}

	task, err := h.taskService.GetTaskToday(c.Request.Context(), telegramID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task commands.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.taskService.SaveTask(c.Request.Context(), &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}
