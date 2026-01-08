package postgres

import (
	"context"

	db "github.com/jabella1/stock-ratings-service/internal/app/features/stock/infra/db/sqlc"
	"github.com/jackc/pgx/v5"
)

func NewConnection(ctx context.Context, dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewQueries(conn *pgx.Conn) *db.Queries {
	return db.New(conn)
}
