package b24models

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const SignatureHeader = "X-Bitrix-Signature"

func VerifyRequest(r *http.Request, secret []byte) error {
	signature := r.Header.Get(SignatureHeader)
	if signature == "" {
		return errors.New("missing signature header")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("body read error: %w", err)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	mac := hmac.New(sha256.New, secret)
	mac.Write(body)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return errors.New("invalid signature")
	}

	return nil
}

func ParseWebhook(r *http.Request) (WebhookEvent, error) {
	var event WebhookEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return WebhookEvent{}, fmt.Errorf("decode error: %w", err)
	}
	return event, nil
}
