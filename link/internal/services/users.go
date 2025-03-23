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

type createUserParams struct {
	Bitrixid   *int64 `json:"bitrixid,omitempty"` // Используем указатель и omitempty
	TelegramID string `json:"telegramid"`
	Name       string `json:"name"`
	Role       string `json:"role"`
}

func CreateUser(data []byte, log *slog.Logger, url string) ([]byte, error) {
	log.Info("Start processing user creation request")

	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		log.Error("Failed to unmarshal input data", logger.ErrToAttr(err))
		return nil, fmt.Errorf("input data parsing failed: %w", err)
	}

	user := createUserParams{}
	var err error

	//  validation required fields
	requiredFields := map[string]func() error{
		"name":       func() error { user.Name, err = getString(rawData, "name"); return err },
		"role":       func() error { user.Role, err = getString(rawData, "role"); return err },
		"telegramid": func() error { user.TelegramID, err = getString(rawData, "telegramid"); return err },
	}

	for key, validator := range requiredFields {
		if err := validator(); err != nil {
			return nil, fmt.Errorf("validation error for field %s: %w", key, err)
		}
	}

	//  validation optional fields
	if val, exists := rawData["bitrixid"]; exists {
		bitrixID, err := getInt64Optional(val)
		if err != nil {
			log.Warn("Invalid bitrixid format", logger.ErrToAttr(err))
		} else {
			user.Bitrixid = &bitrixID
		}
	}

	log.Info("Successfully parsed user data",
		slog.String("name", user.Name),
		slog.String("role", user.Role),
		slog.String("telegramid", user.TelegramID),
		slog.Any("bitrixid", user.Bitrixid))

	requestBody, err := json.Marshal(user)
	if err != nil {
		log.Error("Failed to marshal user data", logger.ErrToAttr(err))
		return nil, fmt.Errorf("user data encoding failed: %w", err)
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

	log.Info("Successfully processed user creation",
		slog.Int("response_size", len(responseData)),
		slog.Int("status_code", resp.StatusCode))

	return responseData, nil
}
