package b24models

import "fmt"

type RegisterBotRequest struct {
	Code       string   `json:"CODE"`
	Type       string   `json:"TYPE"`         // B(default chat-bot) or S(supervisor)
	Handler    string   `json:"HANDLER"`      //handler URL
	AuthUserID int      `json:"AUTH_USER_ID"` //Author id
	Props      BotProps `json:"PROPS"`
}

type BotProps struct {
	Name           string       `json:"NAME"`
	LastName       string       `json:"LAST_NAME,omitempty"`
	Color          string       `json:"COLOR,omitempty"` // HEX-color
	Email          string       `json:"EMAIL,omitempty"`
	WorkPosition   string       `json:"WORK_POSITION,omitempty"`
	PersonalPhoto  string       `json:"PERSONAL_PHOTO,omitempty"` //image URL
	WelcomeMessage string       `json:"MESSAGE_WELCOME_MESSAGE,omitempty"`
	Commands       []BotCommand `json:"COMMANDS,omitempty"`
}

type BotCommand struct {
	Command     string `json:"COMMAND"`
	Handler     string `json:"HANDLER"` //handler URL
	Description string `json:"DESCRIPTION,omitempty"`
}

type RegisterBotResponse struct {
	Result struct {
		BotID int    `json:"BOT_ID"`
		Code  string `json:"CODE"`
		Type  string `json:"TYPE"`
	} `json:"result"`
	Error string `json:"error"`
}

func (c *Client) RegisterBot(params RegisterBotRequest) (*RegisterBotResponse, error) {
	const method = "imbot.register"

	// Validation of mandatory parameters
	if params.Code == "" || params.Handler == "" || params.AuthUserID == 0 {
		return nil, fmt.Errorf("missing required parameters")
	}

	var result RegisterBotResponse
	err := c.CallMethod(method, params, &result)
	if err != nil {
		return nil, fmt.Errorf("registration failed: %v", err)
	}

	if result.Error != "" {
		return nil, fmt.Errorf("bitrix API error: %s", result.Error)
	}

	return &result, nil
}
