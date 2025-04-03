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

	task, err := marshalCreateTask(data, log)
	if err != nil {
		return nil, err
	}
	requestBody, err := json.Marshal(task)
	if err != nil {
		log.Error("JSON marshal error", logger.ErrToAttr(err))
		return nil, fmt.Errorf("data encoding failed: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		log.Error("HTTP request failed", logger.ErrToAttr(err))
		return nil, fmt.Errorf("service unavailable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		log.Error("Service error response",
			slog.Int("status", resp.StatusCode),
			slog.String("body", string(body)))
		return nil, fmt.Errorf("service returned %d status", resp.StatusCode)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response", logger.ErrToAttr(err))
		return nil, fmt.Errorf("response read failed: %w", err)
	}

	log.Info("Request processed successfully",
		slog.Int("response_size", len(responseData)),
		slog.Int("status", resp.StatusCode))

	return responseData, nil
}

func marshalCreateTask(data []byte, log *slog.Logger) (createTaskParams, error) {
	var task createTaskParams
	if err := json.Unmarshal(data, &task); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return task, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed task data",
		slog.Int64("machineid", task.Machineid),
		slog.Int64("shiftid", task.Shiftid),
		slog.String("frequency", task.Frequency),
		slog.String("taskpriority", task.Taskpriority),
		slog.String("description", task.Description),
		slog.Int64("createdby", task.Createdby))

	return task, nil
}
