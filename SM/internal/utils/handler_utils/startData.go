package handler_utils

import (
	"sm/internal/transport/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestParams struct {
	StartTime time.Time
	RequestId string
}

func CreateStartData(c *gin.Context) RequestParams {
	requestID := middleware.RequestIdFromContext(c)
	startTime := time.Now()
	return RequestParams{
		StartTime: startTime,
		RequestId: requestID,
	}
}
