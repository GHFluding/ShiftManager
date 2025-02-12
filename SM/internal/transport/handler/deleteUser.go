package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteUser(p Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with user_list handler"
		reqParams := createStartData(c)
		requestLogger(p.log, reqParams, handlerName, "Start", nil)
		userIdFromPath := c.Param("id")
		userId, err := strconv.Atoi(userIdFromPath)
		if err != nil {
			err := errors.New("invalid user id")
			requestLogger(p.log, reqParams, handlerName, "Error", err)
			return
		}
		err = p.db.DeleteUser(context.Background(), int64(userId))
		if err != nil {
			requestLogger(p.log, reqParams, handlerName, "Error", err)
			return
		}
		requestLogger(p.log, reqParams, handlerName, "Successfully", nil)
	}
}
