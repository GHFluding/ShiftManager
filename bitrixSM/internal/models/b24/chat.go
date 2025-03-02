package b24models

import "fmt"

// imbot.chat.add
type ChatAddRequest struct {
	Title       string `json:"TITLE"`
	Description string `json:"DESCRIPTION,omitempty"`
	Users       []int  `json:"USERS"`
	Color       string `json:"COLOR,omitempty"`
	Message     string `json:"MESSAGE,omitempty"`
	Avatar      string `json:"AVATAR,omitempty"`
}

type ChatAddResponse struct {
	Result struct {
		ChatID int `json:"CHAT_ID"`
	} `json:"result"`
	Error string `json:"error"`
}

func (c *Client) CreateChat(params ChatAddRequest) (int, error) {
	const method = "imbot.chat.add"

	if params.Title == "" || len(params.Users) == 0 {
		return 0, fmt.Errorf("title and users are required")
	}

	var result ChatAddResponse
	err := c.CallMethod(method, params, &result)
	if err != nil {
		return 0, err
	}

	if result.Error != "" {
		return 0, fmt.Errorf(result.Error)
	}

	return result.Result.ChatID, nil
}

// imbot.chat.get
type ChatGetRequest struct {
	ChatID int `json:"CHAT_ID"`
}

type ChatInfo struct {
	ID          int    `json:"ID"`
	Title       string `json:"TITLE"`
	Description string `json:"DESCRIPTION"`
	Users       []int  `json:"USERS"`
	Avatar      string `json:"AVATAR"`
	Color       string `json:"COLOR"`
}

type ChatGetResponse struct {
	Result struct {
		Chat ChatInfo `json:"CHAT"`
	} `json:"result"`
	Error string `json:"error"`
}

// GetChat get chat information
func (c *Client) GetChat(chatID int) (*ChatInfo, error) {
	const method = "imbot.chat.get"

	if chatID == 0 {
		return nil, fmt.Errorf("chatID is required")
	}

	params := ChatGetRequest{ChatID: chatID}
	var result ChatGetResponse

	err := c.CallMethod(method, params, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, fmt.Errorf(result.Error)
	}

	return &result.Result.Chat, nil
}

// imbot.dialog.get
type DialogGetRequest struct {
	DialogID string `json:"DIALOG_ID"`
}

type DialogInfo struct {
	ChatID        int   `json:"CHAT_ID"`
	LastMessageID int   `json:"LAST_MESSAGE_ID"`
	UnreadCount   int   `json:"UNREAD_COUNT"`
	Users         []int `json:"USERS"`
}

type DialogGetResponse struct {
	Result struct {
		Dialog DialogInfo `json:"DIALOG"`
	} `json:"result"`
	Error string `json:"error"`
}

// GetDialog get dialog information
func (c *Client) GetDialog(dialogID string) (*DialogInfo, error) {
	const method = "imbot.dialog.get"

	if dialogID == "" {
		return nil, fmt.Errorf("dialogID is required")
	}

	params := DialogGetRequest{DialogID: dialogID}
	var result DialogGetResponse

	err := c.CallMethod(method, params, &result)
	if err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, fmt.Errorf(result.Error)
	}

	return &result.Result.Dialog, nil
}

//im.message.add
type SendMessageRequest struct {
	Dialog      string        `json:"DIALOG_ID"` //chat or user id(not dialog)
	Message     string        `json:"MESSAGE"`
	Attachments []interface{} `json:"ATTACH,omitempty"`
}

func (c *Client) SendMessage(dialogId, text string) error {
	reqParams := SendMessageRequest{
		Dialog:  dialogId,
		Message: text,
	}
	return c.CallMethod("im.message.add", reqParams, nil)
}
