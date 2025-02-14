package handler

import (
	"context"
	"net/http"
	"sm/internal/database/postgres"
	"sm/internal/utils/handler_utils"
	handler_output "sm/internal/utils/handler_utils/output"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

func GetShiftList(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with shift_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		shifts, err := p.DB.ShiftList(context.Background())
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		shiftsOut, err := shiftsListCovert(shifts)
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"shifts": shiftsOut,
		})
	}
}

func GetActiveShiftList(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with active_shift_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		shifts, err := p.DB.ActiveShiftList(context.Background())
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		shiftsOut, err := shiftsListCovert(shifts)
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"shifts": shiftsOut,
		})
	}
}

func shiftsListCovert(shifts []postgres.Shift) ([]interface{}, error) {
	var shiftsOut []interface{}
	for i := range shifts {
		userOut, err := handler_output.ConvertToOut(shifts[i])
		shiftsOut = append(shiftsOut, userOut)
		if err != nil {
			return nil, err
		}
	}
	return shiftsOut, nil
}
