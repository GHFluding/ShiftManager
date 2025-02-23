package handler

import (
	"log/slog"
	"net/http"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type addWorkerDTO struct {
	ShiftId  int64
	Workerid int64
}

func AddShiftWorker(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with add_shift_worker handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		req, err := parseAddShiftWorkerRequest(c, log)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		shiftWorkerParams := convertShiftWorkerToService(req)
		shiftWorker, err := services.AddShiftWorker(sp, shiftWorkerParams)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"shiftWorker": shiftWorker,
		})
	}
}

func parseAddShiftWorkerRequest(c *gin.Context, log *slog.Logger) (addWorkerDTO, error) {
	var req addWorkerDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return req, err
	}
	return req, nil
}

func convertShiftWorkerToService(req addWorkerDTO) services.ShiftWorker {
	return services.ShiftWorker{
		Shiftid: req.ShiftId,
		Userid:  req.Workerid,
	}
}
