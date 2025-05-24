package logger

import (
	"log/slog"
	"time"

	"github.com/GHFluding/ShiftManager/SM/internal/utils/handler_utils"
)

// logger command string: Start, Error, Successfully
func RequestLogger(log *slog.Logger, reqParams handler_utils.RequestParams, handlerName string, command string, err error) {
	switch command {
	case "Start":
		log.Info(reqParams.StartTime.String(), "id: ", reqParams.RequestId, "Start ", handlerName)
	case "Error":
		log.Info(reqParams.StartTime.String(), "id: ", reqParams.RequestId, "Failed ", handlerName, ErrToAttr(err))
	case "Successfully":
		log.Info(reqParams.StartTime.String(), "id: ", reqParams.RequestId, "Successfully ",
			handlerName, "Request duration", time.Since(reqParams.StartTime).String())
	}

}
