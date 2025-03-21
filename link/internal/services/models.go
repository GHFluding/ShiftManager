package services

import "log/slog"

type WebhookTransformDataFunc func([]byte, *slog.Logger) error

func ExampleFunc(data []byte, log *slog.Logger) error {
	log.Info("This is example function")
	return nil
}
