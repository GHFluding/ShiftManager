package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	logger "linkSM/internal/utils"
	"log/slog"
)

type createTaskParams struct {
	Machineid    int64  `json:"machineid"`
	Shiftid      int64  `json:"shiftid"`
	Frequency    string `json:"frequency"`
	Taskpriority string `json:"taskpriority"`
	Description  string `json:"description"`
	Createdby    int64  `json:"createdby"`
}

func CreateTask(data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing task creation request")

	// map[string]interface{}, because the request may not match the structure
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		log.Error("Failed to unmarshal input data", logger.ErrToAttr(err))
		return nil, fmt.Errorf("input data parsing failed: %w", err)
	}

	// validation fields
	var task createTaskParams
	var err error

	fields := map[string]interface{}{}
	for key, converter := range map[string]func() error{
		"machineid":    func() error { task.Machineid, err = getInt64(rawData, "machineid"); return err },
		"shiftid":      func() error { task.Shiftid, err = getInt64(rawData, "shiftid"); return err },
		"frequency":    func() error { task.Frequency, err = getString(rawData, "frequency"); return err },
		"taskpriority": func() error { task.Taskpriority, err = getString(rawData, "taskpriority"); return err },
		"description":  func() error { task.Description, err = getString(rawData, "description"); return err },
		"createdby":    func() error { task.Createdby, err = getInt64(rawData, "createdby"); return err },
	} {
		if err := converter(); err != nil {
			return nil, fmt.Errorf("validation error for field %s: %w", key, err)
		}
		fields[key] = rawData[key]
	}

	log.Info("Successfully parsed task data", slog.Any("fields", fields))

	requestBody, err := json.Marshal(task)
	if err != nil {
		log.Error("Failed to marshal task data", logger.ErrToAttr(err))
		return nil, fmt.Errorf("task data encoding failed: %w", err)
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Post(url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		log.Error("HTTP request failed", logger.ErrToAttr(err))
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Error("Unexpected HTTP status",
			slog.Int("status", resp.StatusCode),
			slog.String("status_text", resp.Status))
		return nil, fmt.Errorf("received bad status: %d %s", resp.StatusCode, resp.Status)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body", logger.ErrToAttr(err))
		return nil, fmt.Errorf("response reading failed: %w", err)
	}

	log.Info("Successfully processed request",
		slog.Int("response_size", len(responseData)),
		slog.Int("status_code", resp.StatusCode))

	return responseData, nil
}
