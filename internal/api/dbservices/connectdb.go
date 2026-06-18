package dbservices

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
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

func ConnectDBTX(ctx context.Context, logger slog.Logger) *pgx.Tx {

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("error opening database:" + err.Error())
		os.Exit(1)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		logger.Error("error begining tx:" + err.Error())
		os.Exit(1)
	}

	return &tx
}
