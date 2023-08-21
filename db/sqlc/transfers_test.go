package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/hleveque/gobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: sql.NullInt64{Int64: account1.ID, Valid: true},
		ToAccountID:   sql.NullInt64{Int64: account2.ID, Valid: true},
		Amount:        util.RandomBalance(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, arg.FromAccountID)
	require.Equal(t, arg.ToAccountID, arg.ToAccountID)
	require.Equal(t, arg.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, transfer.Amount, transfer1.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	trans, _ := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Empty(t, trans)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, acc := range transfers {
		require.NotEmpty(t, acc)
	}
}
