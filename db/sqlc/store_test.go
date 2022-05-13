package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {

	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	// run 5 concurrent transactions
	n := 5
	amount := int64(100)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	fmt.Println(">>before: ", account1.Balance, account2.Balance)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check that transfer was created properly
		transfer := result.Transfer

		require.NotEmpty(t, transfer)

		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)

		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.CreatedAt)
		require.NotZero(t, transfer.ID)

		_, err = store.GetTranfer(context.Background(), transfer.ID)

		require.NoError(t, err)

		// check that entries were created properly

		// from entry
		fromEntry := result.fromEntry

		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.CreatedAt)
		require.NotZero(t, fromEntry.ID)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// to entry

		toEntry := result.toEntry

		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.CreatedAt)
		require.NotZero(t, toEntry.ID)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check that accounts were updated properly

		// from account
		fromAccount, err := store.GetAccount(context.Background(), account1.ID)

		require.NotEmpty(t, fromAccount)
		require.NoError(t, err)
		require.Equal(t, account1.ID, fromAccount.ID)

		// to account
		toAccount, err := store.GetAccount(context.Background(), account2.ID)

		require.NoError(t, err)
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check the money before and after transfer in both accounts

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := account2.Balance - toAccount.Balance

		require.Equal(t, -diff2, diff1)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)

		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check final balance of two accounts

	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)

	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}
