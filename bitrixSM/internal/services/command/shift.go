package bot_command

import (
	config "bsm/internal/config/loadconfig"
	"bsm/internal/utils/logger"
	"log/slog"
)

func CreateShift(baseURL string, args []string, log *slog.Logger) (string, error) {
	webhookURL := baseURL + "/shift"
	resp, err := sendPostRequest(webhookURL, args)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	return resp, err
}

func ShiftList(baseURL string, log *slog.Logger) (string, error) {
	webhookURL := baseURL + config.Routes.Shift
	resp, err := sendGetListRequest(webhookURL)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	log.Info("ShiftList request was successful.")
	return resp, nil
}
