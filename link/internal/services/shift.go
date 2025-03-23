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

type createShiftParams struct {
	Machineid   int64 `json:"machineid" `
	ShiftMaster int64 `json:"shiftmaster"`
}

func CreateShift(data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing shift creation request")

	// map[string]interface{}, because the request may not match the structure
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		log.Error("Failed to unmarshal input data", logger.ErrToAttr(err))
		return nil, fmt.Errorf("input data parsing failed: %w", err)
	}

	// validation fields
	var shift createShiftParams
	var err error

	fields := map[string]interface{}{}
	for key, converter := range map[string]func() error{
		"machineid":   func() error { shift.Machineid, err = getInt64(rawData, "machineid"); return err },
		"shiftmaster": func() error { shift.ShiftMaster, err = getInt64(rawData, "shiftmaster"); return err },
	} {
		if err := converter(); err != nil {
			return nil, fmt.Errorf("validation error for field %s: %w", key, err)
		}
		fields[key] = rawData[key]
	}

	log.Info("Successfully parsed task data", slog.Any("fields", fields))

	requestBody, err := json.Marshal(shift)
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
