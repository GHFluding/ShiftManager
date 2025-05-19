package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/GHFluding/ShiftManager/SM/internal/services"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/handler_utils"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// GetShiftTaskList Return list of shifts task by shift id.
// @Summary      Get list of shifts task by shift id
// @Description  Return list of shifts task by shift id.
// @Tags         shift task
// @Accept       json
// @Produce      json
// @Param        id   path      int64  true  "Shift id" format(id)
// @Success 	 200 {array} services.ShiftTask "List of task by shift id"
// @Failure      400  {object}  map[string]interface{} "invalid data"
// @Router       /api/users/{id} [get]
func GetShiftTaskList(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with shift_task_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		shiftidFromContext := c.Param("id")
		shiftid, err := strconv.Atoi(shiftidFromContext)
		if err != nil {
			err := errors.New("invalid shift id")
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		shiftTasksService, err := services.ShiftTasksList(sp, int64(shiftid))
		if err != nil {
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"Tasks": shiftTasksService,
		})
	}
}
