package webhooks

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	"github.com/GHFluding/ShiftManager/link/internal/services"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
	"github.com/GHFluding/ShiftManager/link/internal/validator"
	"github.com/gin-gonic/gin"
)

func ProcessWebhookGRPC(log *slog.Logger, url string) gin.HandlerFunc {
	cl, err := client.New(url)
	if err != nil {
		panic(err)
	}
	handlers := map[string]map[string]services.WebhookProcessingFuncGRPC{
		"machine": {
			"create": processMachineCreate,
		},
		"user": {
			"create": processUserCreate,
		},
		"task": {
			"create": processTaskCreate,
		},
		"shift": {
			"create": processShiftCreate,
		},
	}
	return func(c *gin.Context) {
		entity := c.Param("entity")
		action := c.Param("action")

		reqLog := log.With(
			slog.String("entity", entity),
			slog.String("action", action),
		)
		reqLog.Info("Incoming webhook request")

		entityHandlers, ok := handlers[entity]
		if !ok {
			reqLog.Error("Unknown entity")
			c.JSON(http.StatusNotFound, gin.H{"error": "Unknown entity"})
			return
		}
		processingFunc, ok := entityHandlers[action]
		if !ok {
			reqLog.Error("Unsupported action for entity")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported action"})
			return
		}
		data, err := c.GetRawData()
		if err != nil {
			reqLog.Error("Failed to read request body", logger.ErrToAttr(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		outData, err := processingFunc(cl, data, reqLog)
		if err != nil {
			reqLog.Error("Processing failed", logger.ErrToAttr(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Processing error",
				"details": err.Error(),
			})
			return
		}
		reqLog.Info("Request processed successfully")
		c.Data(http.StatusOK, "application/json", outData)
	}
}

func processMachineCreate(c *client.Client, data []byte, log *slog.Logger) ([]byte, error) {
	req, err := validator.Machine(data, log)
	if err != nil {
		return nil, err
	}
	respGRPC, err := services.CreateMachineGRPC(c, req.ToGRPCCreateParams(), log)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(respGRPC)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func processTaskCreate(c *client.Client, data []byte, log *slog.Logger) ([]byte, error) {
	req, err := validator.Task(data, log)
	if err != nil {
		return nil, err
	}
	respGRPC, err := services.CreateTaskGRPC(c, req.ToGRPCCreateParams(), log)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(respGRPC)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func processUserCreate(c *client.Client, data []byte, log *slog.Logger) ([]byte, error) {
	req, err := validator.User(data, log)
	if err != nil {
		return nil, err
	}
	respGRPC, err := services.CreateUserGRPC(c, req.ToGRPCCreateParams(), log)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(respGRPC)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func processShiftCreate(c *client.Client, data []byte, log *slog.Logger) ([]byte, error) {
	req, err := validator.Shift(data, log)
	if err != nil {
		return nil, err
	}
	respGRPC, err := services.CreateShiftGRPC(c, req.ToGRPCCreateParams(), log)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(respGRPC)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
