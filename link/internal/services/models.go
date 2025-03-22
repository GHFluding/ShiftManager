package services

import (
	"log/slog"
)

type WebhookProcessingFunc func(data []byte, log *slog.Logger, url string) ([]byte, error)
