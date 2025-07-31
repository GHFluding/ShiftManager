package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"log/slog"

	"github.com/GHFluding/ShiftManager/SMgrpc/pkg/client"
	entities "github.com/GHFluding/ShiftManager/SMgrpc/pkg/gen"
	logger "github.com/GHFluding/ShiftManager/link/internal/utils"
	"github.com/gin-gonic/gin"
)

const grpcTimeout = 5 * time.Second

type User struct {
	TelegramID string `json:"telegram_id"`
	Name       string `json:"name"`
	BitrixID   int    `json"bitrixid"`
	Role       string `json:"role"`
}

type UserService struct {
	grpcClient *client.Client
	log        *slog.Logger
}

func NewUserService(grpcClient *client.Client, log *slog.Logger) *UserService {
	return &UserService{
		grpcClient: grpcClient,
		log:        log,
	}
}

func (s *UserService) SaveUser(ctx context.Context, user *User) error {
	s.log.Info("Saving user",
		slog.String("telegramID", user.TelegramID),
		slog.Int("bitrixID", user.BitrixID),
		slog.String("name", user.Name),
		slog.String("role", user.Role))

	var bitrixID *int64
	if user.BitrixID != 0 {
		bitrixID = new(int64)
		*bitrixID = int64(user.BitrixID)
	}

	_, err := s.grpcClient.CreateUser(ctx, &entities.CreateUserParams{
		TelegramId: user.TelegramID,
		BitrixId:   bitrixID,
		Name:       user.Name,
		Role:       user.Role,
	})
	if err != nil {
		s.log.Error("Failed to save user", logger.ErrToAttr(err))
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (s *UserService) ProcessTelegramRequest(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Invalid request format", logger.ErrToAttr(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	s.log.Info("Processing Telegram request",
		slog.String("telegramID", req.TelegramID),
		slog.Int("bitrixID", req.BitrixID),
		slog.String("name", req.Name))

	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeout)
	defer cancel()

	if err := s.SaveUser(ctx, &req); err != nil {
		s.log.Error("Failed to process Telegram request", logger.ErrToAttr(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process request"})
		return
	}
}
