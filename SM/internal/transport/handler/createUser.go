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

type createUserDTO struct {
	ID       int64  `json:"id"`
	Bitrixid int64  `json:"bitrixid"`
	Name     string `json:"name" `
	Role     string `json:"role"`
}

func CreateUser(log *slog.Logger, sp *services.ServicesParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		const handlerName = "get request with create_user handler"
		reqParams := handler_utils.CreateStartData(c)
		logger.RequestLogger(log, reqParams, handlerName, "Start", nil)
		req, err := parseCreateUserRequest(c, log)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		ok := CheckUserRole(req.Role)
		if !ok {
			err := errors.New("Invalid userrole")
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
		}
		userParams := convertUserForServices(req)
		err = services.CreateUser(sp, userParams)
		if err != nil {
			logger.RequestLogger(log, reqParams, handlerName, "Error", err)
			return
		}
		logger.RequestLogger(log, reqParams, handlerName, "Successfully", nil)
		c.JSON(http.StatusOK, gin.H{})
	}
}

func parseCreateUserRequest(c *gin.Context, log *slog.Logger) (createUserDTO, error) {
	var req createUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request payload", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return req, err
	}
	return req, nil
}

func CheckUserRole(sRole string) bool {
	switch sRole {
	case "engineer":
		return true
	case "worker":
		return true
	case "master":
		return true
	case "manager":
		return true
	case "admin":
		return true
	default:
		return false
	}
}

func convertUserForServices(req createUserDTO) services.User {
	var user services.User
	user.ID = req.ID
	user.Bitrixid = req.Bitrixid
	user.Name = req.Name
	user.Role = req.Role
	return user
}
