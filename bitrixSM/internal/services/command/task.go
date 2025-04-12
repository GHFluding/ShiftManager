package bot_command

import (
	config "bsm/internal/config/loadconfig"
	logger "bsm/internal/utils"
	"log/slog"
)

func CreateTask(baseURL string, args []string, log *slog.Logger) error {
	webhookURL := baseURL + config.Routes.Task
	resp, err := sendPostRequest(webhookURL, args)
	if err != nil {
		log.Info("error in receiving CreateTask response: ", logger.ErrToAttr(err))
		return err
	}
	_ = resp
	log.Info("CreateTask request was successful.")
	return nil
}

func GetTaskList(webhookURL string, log *slog.Logger) (string, error) {
	resp, err := sendGetListRequest(webhookURL)
	if err != nil {
		log.Info("error in receiving TaskList")
		return "", err
	}
	return resp, nil
}
