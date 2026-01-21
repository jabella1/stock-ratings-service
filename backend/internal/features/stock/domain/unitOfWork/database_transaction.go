package unitofwork

import (
	"context"

	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/sqlc"
)

type DatabaseTransaction interface {
	Commit(context context.Context) error
	Rollback(context context.Context) error
	Queries() *sqlc.Queries
}
