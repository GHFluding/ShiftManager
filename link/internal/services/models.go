package services

import (
	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
)

type WebhookProcessingFunc func(data []byte, log *slog.Logger, url string) ([]byte, error)

type WebhookProcessingFuncGRPC func(c *client.Client, data []byte, log *slog.Logger) ([]byte, error)
