package bot_command

import (
	config "bsm/internal/config/loadconfig"
	logger "bsm/internal/utils"
	"log/slog"
)

func AddShiftWorker(baseURL string, args []string, log *slog.Logger) (string, error) {
	webhookURL := baseURL + config.Routes.ShiftWorker
	resp, err := sendPostRequest(webhookURL, args)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	return resp, err
}
