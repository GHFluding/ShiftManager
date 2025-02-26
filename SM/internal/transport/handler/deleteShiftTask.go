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

// DeleteShiftWorker Delete shift worker by id.
// @Summary      Delete shift worker
// @Description   Delete a shift worker:id from the database.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int64 true "Task id"
// @Param        id   path      int64 true "shift id"
// @Success      204  {object} map[string]interface{} "No connection"
// @Failure      400  {object} map[string]interface{} "invalid data"
// @Failure      404  {object} map[string]interface{} "missing id"
// @Router       /api/shift/{shiftid}/{taskid} [delete]
func DeleteShiftTask(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		taskIdFromPath := c.Param("taskid")
		shiftIdFromPath := c.Param("shiftid")
		taskid, err := strconv.Atoi(taskIdFromPath)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", errors.New("invalid user id"))
			return
		}
		shiftId, err := strconv.Atoi(shiftIdFromPath)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", errors.New("invalid shift id"))
			return
		}
		err = services.DeleteShiftTask(sp, int64(taskid), int64(shiftId))
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{})
	}
}
