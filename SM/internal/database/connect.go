package database

import (
	"context"
	"fmt"
	"log/slog"
	"sm/internal/config"
	"sm/internal/utils/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg config.Config, log *slog.Logger) (*pgxpool.Pool, error) {
	dbConnect := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		cfg.Storage.User,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.DBName,
	)

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, dbConnect)
	if err != nil {
		log.Info("unable to connect to database", logger.ErrToAttr(err))
		return nil, err
	}

	if err = dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		log.Info("database ping failed", logger.ErrToAttr(err))
		return nil, err
	}

	log.Info("Connected to the database successfully")
	return dbpool, nil
}
