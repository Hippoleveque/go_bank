package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/hleveque/gobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	id := sql.NullInt64{
		Int64: account.ID,
		Valid: true,
	}
	arg := CreateEntryParams{
		AccountID: id,
		Amount: util.RandomBalance(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, arg.AccountID)
	require.Equal(t, arg.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry 
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.Equal(t, entry.Amount, entry1.Amount)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	entr, _ := testQueries.GetEntry(context.Background(), entry.ID)
	require.Empty(t, entr)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, acc := range entries {
		require.NotEmpty(t, acc)
	}
}
