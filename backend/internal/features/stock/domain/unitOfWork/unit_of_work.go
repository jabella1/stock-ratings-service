package unitofwork

import "context"

type UnitOfWork interface {
	BeginTransaction(context context.Context) (DatabaseTransaction, error)
}
