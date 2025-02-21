package handler

import (
	"errors"
	"log/slog"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteUser Delete user by id.
// @Summary      Delete user
// @Description   Delete a user:id from the database.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int64 true "User id"
// @Success      204  {object} gin.H "No connection"
// @Failure      400  {object} gin.H "invalid data"
// @Failure      404  {object}  gin.H "missing id"
// @Router       /api/users/{id} [delete]
func DeleteUser(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		userIdFromPath := c.Param("id")
		userId, err := strconv.Atoi(userIdFromPath)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", errors.New("invalid user id"))
			return
		}
		err = services.DeleteUser(sp, int64(userId))
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
	}
}
