package handler

import (
	"log/slog"
	"net/http"

	"github.com/GHFluding/ShiftManager/SM/internal/services"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/handler_utils"
	"github.com/GHFluding/ShiftManager/SM/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// ListStudentsHandler - return list of users
// @Summary Get list of users
// @Description Get all users
// @Produce json
// @Success 200 {array} services.User  "List of users"
// @Failure 500 {object} map[string]interface{} "Server error"
// @Router /api/users [get]
func GetUserList(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		usersService, err := services.UsersList(sp)
		if err != nil {
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"users": usersService,
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
// @Success 	 200 {array} services.User "List of users with role"
// @Failure      400  {object} map[string]interface{} "invalid data"
// @Router       /api/users/{role} [get]
func GetUserListByRole(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list_by_id handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		role := c.Param("role")
		users, err := services.UsersListByRole(sp, role)
		if err != nil {
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			role: users,
		})
	}
}
