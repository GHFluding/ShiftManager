package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type createTaskParams struct {
	Machineid    int64
	Shiftid      int64
	Frequency    string
	Taskpriority string
	Description  string
	Createdby    int64
}

func CreateTask(c *gin.Context) error {
	rawUserID, exists := c.Get("userID")
	if !exists {
		return fmt.Errorf("userID not found in context")
	}
	message := c.GetString("message")

	userID, ok := rawUserID.(int64)
	if !ok {
		if intUserID, ok := rawUserID.(int); ok {
			userID = int64(intUserID)
		} else {
			return fmt.Errorf("invalid userID type: %T", rawUserID)
		}
	}
	role, err := CheckUserRole(userID)
	if err != nil {
		return err
	}
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have enough rights"})
		return fmt.Errorf("no admin role")
	}

	params, err := parseTaskParams(userID, message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	data, err := json.Marshal(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to serialize request"})
		return err
	}

	resp, err := http.Post("https://example.url/api/users", "application/json", bytes.NewBuffer(data))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return err
	}
	defer resp.Body.Close()

	c.JSON(resp.StatusCode, gin.H{"message": "task created successfully"})
	return nil
}

func parseTaskParams(userID int64, msg string) (*createTaskParams, error) {
	params := make(map[string]string)
	parts := strings.Fields(msg)

	if len(parts) < 1 {
		return nil, fmt.Errorf("empty command")
	}

	for _, part := range parts[1:] {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key := strings.ToLower(kv[0])
		params[key] = strings.Trim(kv[1], `"'`)
	}

	required := []string{"machineid", "shiftid", "taskpriority", "frequency"}
	for _, field := range required {
		if _, ok := params[field]; !ok {
			return nil, fmt.Errorf("missing required parameter: %s", field)
		}
	}

	machineID, err := strconv.ParseInt(params["machineid"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid machineid: must be a number")
	}

	shiftID, err := strconv.ParseInt(params["shiftid"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid shiftid: must be a number")
	}

	validPriorities := map[string]bool{
		"low":      true,
		"middle":   true,
		"hot_task": true,
		"hot_fix":  true,
	}
	if !validPriorities[strings.ToLower(params["taskpriority"])] {
		return nil, fmt.Errorf("invalid taskpriority: allowed values are low, medium, high")
	}

	validFrequencies := map[string]bool{
		"one_time":  true,
		"daily":     true,
		"weekly":    true,
		"monthly":   true,
		"quarterly": true,
	}
	if !validFrequencies[strings.ToLower(params["frequency"])] {
		return nil, fmt.Errorf("invalid frequency: allowed values are hourly, daily, weekly, monthly")
	}

	return &createTaskParams{
		Machineid:    machineID,
		Shiftid:      shiftID,
		Frequency:    strings.ToLower(params["frequency"]),
		Taskpriority: strings.ToLower(params["taskpriority"]),
		Description:  params["description"],
		Createdby:    userID,
	}, nil
}
