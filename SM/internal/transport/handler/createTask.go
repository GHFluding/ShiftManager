package handler

import (
	"log/slog"
	"net/http"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type createTaskDTO struct {
	ID           int64  `json:"id"`
	Machineid    int64  `json:"machineid"`
	Shiftid      int64  `json:"shiftid" `
	Frequency    string `json:"frequency"`
	Taskpriority string `json:"taskpriority"`
	Description  string `json:"description"`
	Createdby    int64  `json:"createdby"`
}

// CreateTask create new task.
// @Summary      create a task
// @Description  create new task in db.
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        user  body  createTaskDTO  true  "Task data"
// @Success      201  {object}  services.Task
// @Failure      400  {object}  map[string]interface{} "Invalid data"
// @Failure 500 {object} map[string]interface{} "Failed"
// @Router       /api/task/ [post]

func CreateTask(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with create_task handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		req, err := parseCreateTaskRequest(c, log)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		taskParams := convertTaskForServices(req)
		task, err := services.CreateTask(sp, taskParams)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"task": task,
		})
	}
}

func parseCreateTaskRequest(c *gin.Context, log *slog.Logger) (createTaskDTO, error) {
	var req createTaskDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return req, err
	}
	return req, nil
}
func convertTaskForServices(req createTaskDTO) services.Task {
	return services.Task{
		ID:           req.ID,
		Machineid:    req.Machineid,
		Shiftid:      req.Shiftid,
		Frequency:    req.Frequency,
		Taskpriority: req.Taskpriority,
		Description:  req.Description,
		Createdby:    req.Createdby,
	}
}
