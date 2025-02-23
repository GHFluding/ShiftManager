package handler

import (
	"log/slog"
	"net/http"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type createMachineDTO struct {
	ID               int64  `json:"id"`
	Name             string `json:"name" `
	Isrepairrequired bool   `json:"Isrepairrequired"`
	Isactive         bool   `json:"Isactive"`
}

func CreateMachine(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with create_machine handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		req, err := parseCreateMachineRequest(c, log)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		machineParams := convertMachineForServices(req)
		machine, err := services.CreateMachine(sp, machineParams)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"machine": machine,
		})
	}
}

func parseCreateMachineRequest(c *gin.Context, log *slog.Logger) (createMachineDTO, error) {
	var req createMachineDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return req, err
	}
	return req, nil
}

func convertMachineForServices(req createMachineDTO) services.Machine {
	return services.Machine{
		ID:               req.ID,
		Name:             req.Name,
		Isrepairrequired: req.Isrepairrequired,
		Isactive:         req.Isactive,
	}
}
