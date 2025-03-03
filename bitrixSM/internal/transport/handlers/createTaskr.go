package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createTaskParams struct {
	Machineid    int64  `json:"machineid"`
	Shiftid      int64  `json:"shiftid" `
	Frequency    string `json:"frequency"`
	Taskpriority string `json:"taskpriority"`
	Description  string `json:"description"`
	Createdby    int64  `json:"userid"`
}

func CreateTask(c *gin.Context) error {

	var req createTaskParams
	//maybe not work this context from bot.WebhookHandler
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return err
	}
	role, err := CheckUserRole(req.Createdby)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check user role"})
		return err
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have enough rights"})
		return fmt.Errorf("no admin role")
	}

	data, err := json.Marshal(req)
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
