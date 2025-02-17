package handler

import (
	"context"
	"errors"
	"net/http"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetShiftTaskList(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with shift_task_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		shiftidFromContext := c.Param("id")
		shiftid, err := strconv.Atoi(shiftidFromContext)
		if err != nil {
			err := errors.New("invalid shift id")
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		shiftTasks, err := p.DB.ShiftTasksList(context.Background(), int64(shiftid))
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"Tasks": shiftTasks,
		})
	}
}
