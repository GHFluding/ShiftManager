package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"sm/internal/services"
	"sm/internal/utils/handler_utils"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// ListStudentsHandler - return list of users
// @Summary Get list of users
// @Description Get all users
// @Produce json
// @Success 200 {array} handler_output.UserOutput "List of users"
// @Failure 500 {object} gin.H "Server error"
// @Router /api/users [get]
func GetUserList(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		usersService := services.UsersList(sp)
		if !usersService.Valid {
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"users": usersService.UserListDTO,
		})
	}
}

// GetStudentByIdHandler Return list of users with role.
// @Summary      Get list of users with role
// @Description  Return list of users with role.
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        role   path      string  true  "Users role" format(id)
// @Success 	 200 {array} handler_output.UserOutput "List of users with role"
// @Failure      400  {object}  gin.H "invalid data"
// @Router       /api/users/{role} [get]
func GetUserListByRole(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list_by_id handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		userRoleParam := c.Param("role")
		role, ok := handler_utils.DetectUserRole(userRoleParam)
		if !ok {
			err := errors.New("invalid user role")
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		usersService := services.UsersListByRole(sp, role)
		if !usersService.Valid {
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			userRoleParam: usersService.UserListDTO,
		})
	}
}
