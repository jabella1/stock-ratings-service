package cockroach

import (
	"context"

	unitofwork "github.com/jabella1/stock-ratings-service/internal/features/stock/domain/unitOfWork"
	"github.com/jabella1/stock-ratings-service/internal/features/stock/infra/db/sqlc"
	"github.com/jackc/pgx/v5"
)

type UnitOfWork struct {
	connection *pgx.Conn
}

func CreateUnitOfWork(connection *pgx.Conn) unitofwork.UnitOfWork {
	return &UnitOfWork{
		connection: connection,
	}
}

func (unitOfWork *UnitOfWork) BeginTransaction(context context.Context) (unitofwork.DatabaseTransaction, error) {
	transaction, err := unitOfWork.connection.Begin(context)
	if err != nil {
		return nil, err
	}

	queries := sqlc.New(transaction)

	return &DatabaseTransaction{
		transaction: transaction,
		queries:     queries,
	}, nil
}
