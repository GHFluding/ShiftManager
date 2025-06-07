package handler

import (
	"log/slog"
	"net/http"

	"github.com/GHFluding/ShiftManager/SM/internal/services"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/handler_utils"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

type createShiftDTO struct {
	ID          int64 `json:"id"`
	Machineid   int64 `json:"machineid" `
	ShiftMaster int64 `json:"shiftmaster"`
}

// CreateShift create new shift.
// @Summary      create a shift
// @Description  create new shift in db.
// @Tags         shift
// @Accept       json
// @Produce      json
// @Param        shift  body  createShiftDTO  true  "Shift data"
// @Success      201  {object}  services.Shift
// @Failure      400  {object}  map[string]interface{} "Invalid data"
// @Failure 500 {object} map[string]interface{} "Failed"
// @Router       /api/task/ [post]
func CreateShift(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with create_shift handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		req, err := parseCreateShiftRequest(c, log)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		shiftParams := convertShiftForServices(req)
		shift, err := services.CreateShift(sp, shiftParams)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"shift": shift,
		})
	}
}

func parseCreateShiftRequest(c *gin.Context, log *slog.Logger) (createShiftDTO, error) {
	var req createShiftDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return req, err
	}
	return req, nil
}

func convertShiftForServices(req createShiftDTO) services.Shift {
	return services.Shift{
		ID:          req.ID,
		Machineid:   req.Machineid,
		ShiftMaster: req.ShiftMaster,
	}
}
