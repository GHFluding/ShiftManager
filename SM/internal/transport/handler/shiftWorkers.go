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

// GetShiftTaskList Return list of shift workers  by shift id.
// @Summary      Get list of shift workers  by shift id
// @Description  Return list of shift workers  by shift id.
// @Tags         shift worker
// @Accept       json
// @Produce      json
// @Param        id   path      int64  true  "Shift id" format(id)
// @Success 	 200 {array} postgres.ShiftWorker "List of shift workers by shift id"
// @Failure      400  {object}  gin.H "invalid data"
// @Router       /api/users/{id} [get]
func GetShiftWorkersList(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with shift_workers_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		shiftidFromContext := c.Param("id")
		shiftid, err := strconv.Atoi(shiftidFromContext)
		if err != nil {
			err := errors.New("invalid shift id")
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		shiftWorkers, err := p.DB.ShiftWorkersList(context.Background(), int64(shiftid))
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"Workers": shiftWorkers,
		})
	}
}
