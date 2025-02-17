package handler

import (
	"context"
	"errors"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ChangeMachineToRepair change machine status to need repair.
// @Summary     change machine status to need repair
// @Description   change machine status to need repair machine:id from the database.
// @Tags         machine
// @Accept       json
// @Produce      json
// @Param        id   path      int64 true "Machine id"
// @Success      204  {object} gin.H "No connection"
// @Failure      400  {object} gin.H "invalid data"
// @Failure      404  {object}  gin.H "missing id"
// @Router       /api/machine/{id} [put]
func ChangeMachineToRepair(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with machine_to_repair handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		machineIdFromPath := c.Param("id")
		machineId, err := strconv.Atoi(machineIdFromPath)
		if err != nil {
			err := errors.New("invalid user id")
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		err = p.DB.MachineNeedRepair(context.Background(), int64(machineId))
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
	}
}
