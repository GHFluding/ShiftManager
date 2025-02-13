package handler

import (
	"context"
	"errors"
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
func DeleteUser(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		userIdFromPath := c.Param("id")
		userId, err := strconv.Atoi(userIdFromPath)
		if err != nil {
			err := errors.New("invalid user id")
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		err = p.DB.DeleteUser(context.Background(), int64(userId))
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
	}
}
