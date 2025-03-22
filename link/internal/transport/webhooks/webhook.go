package webhooks

import (
	"linkSM/internal/services"
	logger "linkSM/internal/utils"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebhookMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// not webhook url is request db url
func WebhookHandler(log *slog.Logger, processingFunc services.WebhookProcessingFunc, url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			log.Error("Failed to read request data", logger.ErrToAttr(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		//outData is result of json.Marshal function
		outData, err := processingFunc(data, log, url)
		if err != nil {
			log.Error("Data processing failed", logger.ErrToAttr(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Data processing error",
				"details": err.Error(),
			})
			return
		}

		c.Data(http.StatusOK, "application/json", outData)

	}
}
