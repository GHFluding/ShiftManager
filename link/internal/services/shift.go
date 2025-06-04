package services

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
)

type createShiftParams struct {
	Machineid   int64 `json:"machineid"`
	ShiftMaster int64 `json:"shiftmaster"`
}

// func CreateShift(data []byte, log *slog.Logger, url string) ([]byte, error) {
// 	log.Info("Start processing shift creation request")

// 	shift, err := marshalCreateShift(data, log)
// 	if err != nil {
// 		return nil, err
// 	}
// 	requestBody, err := json.Marshal(shift)
// 	if err != nil {
// 		log.Error("JSON marshal error", logger.ErrToAttr(err))
// 		return nil, fmt.Errorf("data encoding failed: %w", err)
// 	}

// 	client := &http.Client{Timeout: 10 * time.Second}
// 	resp, err := client.Post(url, "application/json", bytes.NewReader(requestBody))
// 	if err != nil {
// 		log.Error("HTTP request failed", logger.ErrToAttr(err))
// 		return nil, fmt.Errorf("service unavailable: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode >= 400 {
// 		body, _ := io.ReadAll(resp.Body)
// 		log.Error("Service error response",
// 			slog.Int("status", resp.StatusCode),
// 			slog.String("body", string(body)))
// 		return nil, fmt.Errorf("service returned %d status", resp.StatusCode)
// 	}

// 	responseData, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Error("Failed to read response", logger.ErrToAttr(err))
// 		return nil, fmt.Errorf("response read failed: %w", err)
// 	}

// 	log.Info("Request processed successfully",
// 		slog.Int("response_size", len(responseData)),
// 		slog.Int("status", resp.StatusCode))

// 	return responseData, nil
// }

func CreateShiftGRPC(c *client.Client, data *entities.CreateShiftParams, log *slog.Logger) (*entities.ShiftResponse, error) {
	log.Info("Start processing shift creation request")

	resp, err := c.CreateShift(context.Background(), data)
	if err != nil {
		log.Error("GRPC request failed", logger.ErrToAttr(err))
		return nil, fmt.Errorf("service unavailable: %w", err)
	}

	return resp, nil
}
