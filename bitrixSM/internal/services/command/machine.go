package bot_command

import (
	"bsm/internal/services/logger"
	"log/slog"
)

func AddMachine(baseURL string, args []string, log *slog.Logger) (string, error) {
	webhookURL := baseURL + "/machine"
	resp, err := sendPostRequest(webhookURL, args)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	return resp, err
}
