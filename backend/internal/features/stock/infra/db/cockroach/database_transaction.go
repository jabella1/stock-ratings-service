package cockroach

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/sqlc"
	"github.com/jackc/pgx/v5"
)

type DatabaseTransaction struct {
	transaction pgx.Tx
	queries     *sqlc.Queries
}

func (databaseTransaction *DatabaseTransaction) Commit(context context.Context) error {
	return databaseTransaction.transaction.Commit(context)
}

func (databaseTransaction *DatabaseTransaction) Rollback(context context.Context) error {
	return databaseTransaction.transaction.Rollback(context)
}

func (databaseTransaction *DatabaseTransaction) Queries() *sqlc.Queries {
	return databaseTransaction.queries
}
