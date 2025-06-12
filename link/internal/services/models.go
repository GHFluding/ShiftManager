package services

import (
	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	"github.com/gin-gonic/gin"
)

type WebhookProcessingFunc func(data []byte, log *slog.Logger, url string) ([]byte, error)

type WebhookProcessingFuncGRPC func(c *client.Client, ctx *gin.Context, log *slog.Logger) ([]byte, error)
