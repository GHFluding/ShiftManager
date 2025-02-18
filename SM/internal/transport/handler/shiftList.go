package handler

import (
	"log/slog"
	"net/http"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// GetActiveShiftList get all shifts.
// @Summary    get all shifts.
// @Description    get all shifts.
// @Tags         shifts
// @Produce json
// @Success 200 {array} handler_output.ShiftOutput "List of shifts"
// @Failure 500 {object} gin.H "Server error"
// @Router /api/shifts [get]
func GetShiftList(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with shift_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		shiftsService := services.ShiftList(sp)
		if !shiftsService.Valid {
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"shifts": shiftsService.ShiftListDTO,
		})
	}
}

// GetActiveShiftList get out shifts that are active.
// @Summary    get out shifts that are active
// @Description   get out shifts that are active.
// @Tags         shifts
// @Produce json
// @Success 200 {array} handler_output.ShiftOutput "List of active shifts"
// @Failure 500 {object} gin.H "Server error"
// @Router /api/shifts [get]
func GetActiveShiftList(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with active_shift_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		shiftsService := services.ActiveShiftList(sp)
		if !shiftsService.Valid {
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"shifts": shiftsService.ShiftListDTO,
		})
	}
}
