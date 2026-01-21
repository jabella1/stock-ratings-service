package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUnitOfWork_BeginTransaction_CommitRollback(t *testing.T) {
	err := godotenv.Load("../../../../../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	context := context.Background()
	connectionString := os.Getenv("CONNECTION_STRING")

	if connectionString == "" {
		t.Fatalf("environment variable CONNECTION_STRING is not set")
	}

	conn, err := NewConnection(context, connectionString)
	assert.NoError(t, err)
	if err == nil {
		defer conn.Close(context)
	}

	unitOfWork := CreateUnitOfWork(conn)

	transaction, err := unitOfWork.BeginTransaction(context)
	assert.NoError(t, err)
	assert.NotNil(t, transaction)
	assert.NotNil(t, transaction.Queries())

	err = transaction.Commit(context)
	assert.NoError(t, err)

	transaction2, err := unitOfWork.BeginTransaction(context)
	assert.NoError(t, err)
	assert.NotNil(t, transaction2)

	err = transaction2.Rollback(context)
	assert.NoError(t, err)
}
