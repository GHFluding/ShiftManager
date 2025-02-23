package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updateTaskDTO struct {
	UserId  int64  `json:"userid"`
	Comment string `json:"comment" `
	Command string `json:"command"`
}

// UpdateTask update task.
// @Summary      update task task by command
// @Description  commands for update task: inProgress, completed, verified, failed
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        user  body  updateTaskDTO  true  "Task data"
// @Success      201  {object}  map[string]interface{} "Successfully"
// @Failure      400  {object}    map[string]interface{} "Invalid data"
// @Failure 500 {object}   map[string]interface{} "Failed"
// @Router       /api/task/{id} [patch]
func UpdateTask(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		taskIdFromPath := c.Param("id")
		userId, err := strconv.Atoi(taskIdFromPath)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", errors.New("invalid user id"))
			return
		}
		taskDTO, err := parseUpdateTaskRequest(c, log)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		taskParams := convertUpdateTaskToService(taskDTO)
		err = services.UpdateTask(sp, int64(userId), taskParams)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{})
	}
}

func parseUpdateTaskRequest(c *gin.Context, log *slog.Logger) (updateTaskDTO, error) {
	var req updateTaskDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return req, err
	}
	return req, nil
}

func convertUpdateTaskToService(req updateTaskDTO) services.UpdateTaskParams {
	return services.UpdateTaskParams{
		UserID:  req.UserId,
		Comment: req.Comment,
		Command: req.Command,
	}
}
