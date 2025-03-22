package bot_command

import (
	config "bsm/internal/config/loadconfig"
	"bsm/internal/utils/logger"
	"log/slog"
)

func AddMachine(baseURL string, args []string, log *slog.Logger) (string, error) {
	webhookURL := baseURL + config.Routes.Machine
	resp, err := sendPostRequest(webhookURL, args)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	return resp, err
}
