package handler

import (
	"errors"
	"log/slog"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteTask Delete task by id.
// @Summary      Delete task
// @Description   Delete a task:id from the database.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int64 true "User id"
// @Success      204  {object} map[string]interface{} "No connection"
// @Failure      400  {object} map[string]interface{} "invalid data"
// @Failure      404  {object} map[string]interface{} "missing id"
// @Router       /api/task/{id} [delete]
func DeleteTask(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		taskIdFromPath := c.Param("id")
		taskId, err := strconv.Atoi(taskIdFromPath)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", errors.New("invalid user id"))
			return
		}
		err = services.DeleteTask(sp, int64(taskId))
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
	}
}
