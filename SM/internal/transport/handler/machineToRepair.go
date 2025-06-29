package handler

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/GHFluding/ShiftManager/SM/internal/services"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/handler_utils"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// ChangeMachineToRepair change machine status to need repair.
// @Summary     change machine status to need repair
// @Description   change machine status to need repair machine:id from the database.
// @Tags         machine
// @Accept       json
// @Produce      json
// @Param        id   path      int64 true "Machine id"
// @Success      204  {object} map[string]interface{} "No connection"
// @Failure      400  {object} map[string]interface{} "invalid data"
// @Failure      404  {object}  map[string]interface{} "missing id"
// @Router       /api/machine/{id} [put]
func ChangeMachineToRepair(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with machine_to_repair handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		machineIdFromPath := c.Param("id")
		machineId, err := strconv.Atoi(machineIdFromPath)
		if err != nil {
			err := errors.New("invalid user id")
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		err = services.MachineNeedRepair(sp, int64(machineId))
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
	}
}
