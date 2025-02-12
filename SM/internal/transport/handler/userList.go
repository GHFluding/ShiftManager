package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserList(p Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := createStartData(c)
		requestLogger(p.log, reqParams, handlerName, "Start", nil)
		users, err := p.db.UsersList(context.Background())
		if err != nil {
			requestLogger(p.log, reqParams, handlerName, "Error", err)
			return
		}
		requestLogger(p.log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func UserListByRole(p Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list_by_id handler"
		reqParams := createStartData(c)
		requestLogger(p.log, reqParams, handlerName, "Start", nil)
		userRoleParam := c.Param("role")
		role, ok := DetectUserRole(userRoleParam)
		if !ok {
			err := errors.New("invalid user role")
			requestLogger(p.log, reqParams, handlerName, "Error", err)
			return
		}
		users, err := p.db.UsersListByRole(context.Background(), role)
		if err != nil {
			requestLogger(p.log, reqParams, handlerName, "Error", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			userRoleParam: users,
		})
	}
}
