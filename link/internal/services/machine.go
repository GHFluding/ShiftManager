package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"log/slog"

	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type createMachineParams struct {
	Name             string `json:"name"`
	IsRepairRequired *bool  `json:"isrepairrequired,omitempty"`
	IsActive         *bool  `json:"isactive,omitempty"`
}

func CreateMachine(data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing machine creation request")

	machine, err := marshalCreateMachine(data, log)
	if err != nil {
		return nil, err
	}
	requestBody, err := json.Marshal(machine)
	if err != nil {
		log.Error("JSON marshal error", logger.ErrToAttr(err))
		return nil, fmt.Errorf("data encoding failed: %w", err)
	}
	if machine.IsActive == nil {
		log.Info("Isactive is not set. Using default value")
	}
	if machine.IsRepairRequired == nil {
		log.Info("IsRepairRequired is not set. Using default value")
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
		log.Error("Service error",
			slog.Int("status", resp.StatusCode),
			slog.String("response", string(body)))
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

func marshalCreateMachine(data []byte, log *slog.Logger) (createMachineParams, error) {

	var machine createMachineParams
	if err := json.Unmarshal(data, &machine); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return machine, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed machine data",
		slog.String("name", machine.Name),
		slog.Any("isrepairrequired", machine.IsRepairRequired),
		slog.Any("isactive", machine.IsActive))
	return machine, nil

}
