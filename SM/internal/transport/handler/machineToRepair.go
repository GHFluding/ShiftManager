package handler

import (
	"context"
	"errors"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
