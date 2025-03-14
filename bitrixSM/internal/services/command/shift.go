package bot_command

import (
	"bsm/internal/services/logger"
	"log/slog"
)

func CreateShift(webhookURL string, args []string, log *slog.Logger) (string, error) {
	resp, err := sendPostRequest(webhookURL, args)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	return resp, err
}

func ShiftList(webhookURL string, log *slog.Logger) (string, error) {
	resp, err := sendGetListRequest(webhookURL)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	log.Info("ShiftList request was successful.")
	return resp, nil
}
