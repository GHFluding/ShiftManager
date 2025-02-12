package handler

import (
	"log/slog"
	"sm/internal/database/postgres"
	"sm/internal/transport/middleware"
	"sm/internal/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func createStartData(c *gin.Context) requestParams {
	requestID := middleware.RequestIdFromContext(c)
	startTime := time.Now()
	return requestParams{
		startTime: startTime,
		requestId: requestID,
	}
}

// logger command string: Start, Error, Successfully
func requestLogger(log *slog.Logger, reqParams requestParams, handlerName string, command string, err error) {
	switch command {
	case "Start":
		log.Info(reqParams.startTime.String(), "id: ", reqParams.requestId, "Start ", handlerName)
	case "Error":
		log.Info(reqParams.startTime.String(), "id: ", reqParams.requestId, "Failed ", handlerName, logger.ErrToAttr(err))
	case "Successfully":
		log.Info(reqParams.startTime.String(), "id: ", reqParams.requestId, "Successfully ",
			handlerName, "Request duration", time.Since(reqParams.startTime).String())
	}

}

func DetectUserRole(sRole string) (postgres.Userrole, bool) {
	switch sRole {
	case "engineer":
		return postgres.UserroleEngineer, true
	case "worker":
		return postgres.UserroleWorker, true
	case "master":
		return postgres.UserroleMaster, true
	case "manager":
		return postgres.UserroleManager, true
	case "admin":
		return postgres.UserroleAdmin, true
	default:
		return postgres.Userrole(""), false
	}
}
