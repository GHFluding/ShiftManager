package b24models

import (
	"encoding/json"
	"time"
)

type WebhookEvent struct {
	Event string          `json:"event"`
	TS    int64           `json:"ts"`
	Data  json.RawMessage `json:"data"`
	Auth  struct {
		Domain   string `json:"domain"`
		MemberID string `json:"member_id"`
	} `json:"auth"`
}

type ImMessage struct {
	ID         string    `json:"id"`
	DialogID   string    `json:"dialog_id"`
	Message    string    `json:"message"`
	Date       time.Time `json:"date"`
	FromUserID string    `json:"from_user_id"`
}

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Responsible int       `json:"responsibleId"`
	Created     time.Time `json:"createdDate"`
	Deadline    time.Time `json:"deadline"`
	Status      string    `json:"status"`
}

type Response struct {
	Result json.RawMessage `json:"result"`
	Time   struct {
		Start    float64 `json:"start"`
		Finish   float64 `json:"finish"`
		Duration float64 `json:"duration"`
	} `json:"time"`
	Error string `json:"error"`
}
