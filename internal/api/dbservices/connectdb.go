package dbservices

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context, logger slog.Logger) *pgxpool.Pool {

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("error opening database:" + err.Error())
		os.Exit(1)
	}

	return pool
}
