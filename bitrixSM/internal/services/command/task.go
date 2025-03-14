package bot_command

import (
	"bsm/internal/services/logger"
	"log/slog"
)

func CreateTask(webhookURL string, args []string, log *slog.Logger) error {
	resp, err := sendPostRequest(webhookURL, args)
	if err != nil {
		log.Info("error in receiving CreateTask response: ", logger.ErrToAttr(err))
		return err
	}
	_ = resp
	log.Info("CreateTask request was successful.")
	return nil
}
