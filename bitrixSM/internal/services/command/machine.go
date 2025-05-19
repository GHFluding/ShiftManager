package bot_command

import (
	"log/slog"

	config "github.com/GHFluding/ShiftManager/bitrixSM/internal/config/loadconfig"
	logger "github.com/GHFluding/ShiftManager/bitrixSM/internal/utils"
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
