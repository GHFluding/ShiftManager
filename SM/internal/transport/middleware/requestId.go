package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("X-Request-Id")
		if requestId == "" {
			requestId = c.GetHeader("RequestId")
		}
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set("RequestId", requestId)
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}

func RequestIdFromContext(c *gin.Context) string {
	if requestId, exists := c.Get("RequestId"); exists {
		return requestId.(string)
	}
	return ""
}
