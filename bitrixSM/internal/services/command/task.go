package bot_command

import (
	"bsm/internal/services/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func CreateTask(args []string, log *slog.Logger) error {
	err := sendCreateTaskRequest(args)
	if err != nil {
		log.Info("error in receiving CreateTask response: ", logger.ErrToAttr(err))
		return err
	}
	log.Info("CreateTask request was successful.")
	return nil
}

func sendCreateTaskRequest(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no data to send request")
	}

	//maybe make params struct
	jsonData, err := json.Marshal(args)
	if err != nil {
		return err
	}

	webhookURL := "https://example.com/api/tasks"

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error sending data, status code: %d", resp.StatusCode)
	}

	return nil
}
