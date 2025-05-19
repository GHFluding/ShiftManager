package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/GHFluding/ShiftManager/SM/internal/services"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/handler_utils"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// Check User Role by id
// @Summary    check role
// @Description  Return user role by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int64 true "user id"
// @Success      204  {object} map[string]interface{} "No connection"
// @Failure      400  {object} map[string]interface{} "invalid data"
// @Failure      404  {object} map[string]interface{} "missing id"
// @Router       /api/users/{id} [get]
func CheckUserRoleHandler(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		userIdFromPath := c.Param("id")
		userid, err := strconv.Atoi(userIdFromPath)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", errors.New("invalid user id"))
			return
		}
		role, err := services.CheckUserRole(sp, int64(userid))
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"role": role,
		})
	}
}
