package db

import (
	"context"
	"testing"
	"time"

	"github.com/Abdelrhman-Hosny/go_bank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T, account Account) Entry {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomEntry(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {

	account := CreateRandomAccount(t)
	CreateRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {

	account := CreateRandomAccount(t)
	entry := CreateRandomEntry(t, account)

	entry_get, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry_get)

	require.Equal(t, entry.ID, entry_get.ID)
	require.Equal(t, entry.AccountID, entry_get.AccountID)
	require.Equal(t, entry.Amount, entry_get.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry_get.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {

	account := CreateRandomAccount(t)
	entry := CreateRandomEntry(t, account)

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: utils.RandomEntry(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Amount, entry2.Amount)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {

	account := CreateRandomAccount(t)
	entry := CreateRandomEntry(t, account)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	entry_get, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err)
	require.Empty(t, entry_get)
}

func TestListEntries(t *testing.T) {

	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, CreateRandomAccount(t))
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, entries, 5)

	for i := 0; i < 5; i++ {
		require.NotEmpty(t, entries[i])
	}
}
