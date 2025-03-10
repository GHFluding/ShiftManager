package bot_command

import (
	"bsm/internal/services/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func ShiftList(args []string, log *slog.Logger) (string, error) {
	resp, err := sendShiftListRequest(args)
	if err != nil {
		log.Info("error in receiving ShiftList response: ", logger.ErrToAttr(err))
		return "", err
	}
	log.Info("ShiftList request was successful.")
	return resp, nil
}

func sendShiftListRequest(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no data to send request")
	}

	jsonData, err := json.Marshal(args)
	if err != nil {
		return "", err
	}

	webhookURL := "https://example.com/api/shift"

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error sending data, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var responseData []string
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", err
	}

	// create a string to send to the client
	result := strings.Join(responseData, "\n")

	return result, nil
}
