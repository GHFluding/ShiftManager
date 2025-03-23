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

type createMachineParams struct {
	Name             string `json:"name" `
	Isrepairrequired bool   `json:"Isrepairrequired",omitempty`
	Isactive         bool   `json:"Isactive",omitempty`
}

func CreateMachine(data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing shift creation request")

	// map[string]interface{}, because the request may not match the structure
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		log.Error("Failed to unmarshal input data", logger.ErrToAttr(err))
		return nil, fmt.Errorf("input data parsing failed: %w", err)
	}

	// validation required fields
	name, err := getString(rawData, "name")
	if err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	machine := createMachineParams{
		Name: name,
	}
	//validation optional fields
	if val, ok := rawData["isrepairrequired"]; ok {
		if isRepair, err := getBool(val); err == nil {
			machine.Isrepairrequired = isRepair
		} else {
			log.Warn("Invalid isrepairrequired value", logger.ErrToAttr(err))
		}
	}

	if val, ok := rawData["isactive"]; ok {
		if isActive, err := getBool(val); err == nil {
			machine.Isactive = isActive
		} else {
			log.Warn("Invalid isactive value", logger.ErrToAttr(err))
		}
	}

	log.Info("Parsed machine data",
		slog.String("name", machine.Name),
		slog.Bool("isrepairrequired", machine.Isrepairrequired),
		slog.Bool("isactive", machine.Isactive))

	requestBody, err := json.Marshal(machine)
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
