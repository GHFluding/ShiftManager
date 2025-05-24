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

type createUserParams struct {
	Bitrixid   *int64 `json:"bitrixid,omitempty"`
	TelegramID string `json:"telegramid"`
	Name       string `json:"name"`
	Role       string `json:"role"`
}

func CreateUser(data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing user creation request")
	user, err := marshalCreateUser(data, log)
	if err != nil {
		return nil, err
	}
	requestBody, err := json.Marshal(user)
	if err != nil {
		log.Error("JSON marshal error", logger.ErrToAttr(err))
		return nil, fmt.Errorf("data encoding failed: %w", err)
	}
	if user.Bitrixid == nil {
		log.Info("BitrixID is not set. This user use only telegram")
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
			slog.String("response", string(body)))
		return nil, fmt.Errorf("service returned %d status: %s", resp.StatusCode, string(body))
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response", logger.ErrToAttr(err))
		return nil, fmt.Errorf("response read failed: %w", err)
	}

	log.Info("User created successfully",
		slog.Int("response_size", len(responseData)),
		slog.Int("status", resp.StatusCode))

	return responseData, nil
}

func marshalCreateUser(data []byte, log *slog.Logger) (createUserParams, error) {
	var user createUserParams
	if err := json.Unmarshal(data, &user); err != nil {
		log.Error("JSON unmarshal error", logger.ErrToAttr(err))
		return user, fmt.Errorf("invalid request format: %w", err)
	}

	log.Info("Parsed user data",
		slog.String("name", user.Name),
		slog.String("role", user.Role),
		slog.String("telegramid", user.TelegramID),
		slog.Any("bitrixid", user.Bitrixid))
	return user, nil
}
