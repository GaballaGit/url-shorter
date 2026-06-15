package dbservices

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context, logger slog.Logger) *pgx.Conn {

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("error opening database:" + err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	return conn
}
