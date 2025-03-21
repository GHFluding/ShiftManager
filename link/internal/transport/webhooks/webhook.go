package webhooks

import (
	"linkSM/internal/services"
	logger "linkSM/internal/utils"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func WebhookMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		servicesMap := map[string]services.WebhookTransformDataFunc{
			"example": services.ExampleFunc,
		}
		urlPath := c.Request.URL.Path
		data, err := c.GetRawData()
		if err != nil {
			log.Info("Error return data from webhook", logger.ErrToAttr(err))
			return
		}
		f, exists := servicesMap[urlPath]
		if !exists {
			log.Info("Missing processing function for this url")
			return
		}
		err = f(data, log)
		if err != nil {
			log.Info("Request is failed ", logger.ErrToAttr(err))
			return
		}
	}
}

func WebhookHandler(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// all processing in middleware
	}
}
