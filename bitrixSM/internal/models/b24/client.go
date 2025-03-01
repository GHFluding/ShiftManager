package b24models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	WebhookURL string
	HTTPClient *http.Client
}

func NewClient(webhookURL string) *Client {
	return &Client{
		WebhookURL: webhookURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) CallMethod(method string, params interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s", c.WebhookURL, method)

	body, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %s", resp.Status)
	}

	var apiResponse Response
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return fmt.Errorf("response decoding failed: %w", err)
	}

	if apiResponse.Error != "" {
		return fmt.Errorf("API error: %s", apiResponse.Error)
	}

	return json.Unmarshal(apiResponse.Result, result)
}
