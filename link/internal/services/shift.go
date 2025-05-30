package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type createShiftParams struct {
	Machineid   int64 `json:"machineid"`
	ShiftMaster int64 `json:"shiftmaster"`
}

func CreateShift(data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing shift creation request")

	shift, err := marshalCreateShift(data, log)
	if err != nil {
		return nil, err
	}
	requestBody, err := json.Marshal(shift)
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

func CreateShiftGRPC(c *client.Client, data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing shift creation request")

	shift, err := marshalCreateShift(data, log)
	if err != nil {
		return nil, err
	}

	resp, err := c.CreateShift(context.Background(), shift.Machineid, shift.ShiftMaster)
	if err != nil {
		log.Error("GRPC request failed", logger.ErrToAttr(err))
		return nil, fmt.Errorf("service unavailable: %w", err)
	}

	//using only response data for marshaling
	responseData, err := json.Marshal(resp.Data)
	if err != nil {
		log.Error("Failed to marshal response", logger.ErrToAttr(err))
		return nil, fmt.Errorf("response marshal failed: %w", err)
	}
	return responseData, nil
}

func marshalCreateShift(data []byte, log *slog.Logger) (createShiftParams, error) {
	var shift createShiftParams
	if err := json.Unmarshal(data, &shift); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return shift, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed shift data",
		slog.Int64("machineid", shift.Machineid),
		slog.Int64("shiftmaster", shift.ShiftMaster))
	return shift, nil

}
