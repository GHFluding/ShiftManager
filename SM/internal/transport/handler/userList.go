package handler

import (
	"context"
	"errors"
	"net/http"
	"sm/internal/database/postgres"
	"sm/internal/utils/handler_utils"
	handler_output "sm/internal/utils/handler_utils/output"
	"sm/internal/utils/logger"

	"github.com/gin-gonic/gin"
)

// ListStudentsHandler - return list of users
// @Summary Get list of users
// @Description Get all users
// @Produce json
// @Success 200 {array} handler_output.UserOutput "List of users"
// @Failure 500 {object} gin.H "Server error"
// @Router /api/students [get]
func GetUserList(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		users, err := p.DB.UsersList(context.Background())
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		usersOut, err := usersListCovert(users)
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"users": usersOut,
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
// @Router       /api/students/{role} [get]
func GetUserListByRole(p handler_utils.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list_by_id handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(p.Log, reqParams, handlerName, "Start", nil)
		userRoleParam := c.Param("role")
		role, ok := handler_utils.DetectUserRole(userRoleParam)
		if !ok {
			err := errors.New("invalid user role")
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		users, err := p.DB.UsersListByRole(context.Background(), role)
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		usersOut, err := usersListCovert(users)
		if err != nil {
			logger.RequestLogger(p.Log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(p.Log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			userRoleParam: usersOut,
		})
	}
}

func usersListCovert(users []postgres.User) ([]interface{}, error) {
	var usersOut []interface{}
	for i := range users {
		userOut, err := handler_output.ConvertToOut(users[i])
		usersOut = append(usersOut, userOut)
		if err != nil {
			return nil, err
		}
	}
	return usersOut, nil
}
