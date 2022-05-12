package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Abdelrhman-Hosny/go_bank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	acccount, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Owner, acccount.Owner)
	require.Equal(t, arg.Currency, acccount.Currency)
	require.Equal(t, arg.Balance, acccount.Balance)

	require.NotZero(t, acccount.ID)
	require.NotZero(t, acccount.CreatedAt)

	return acccount

}
func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// create account
	account := CreateRandomAccount(t)

	account_get, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account_get)

	require.Equal(t, account.ID, account_get.ID)
	require.Equal(t, account.Owner, account_get.Owner)
	require.Equal(t, account.Currency, account_get.Currency)
	require.Equal(t, account.Balance, account_get.Balance)
	require.WithinDuration(t, account.CreatedAt, account_get.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := UpdateAccountBalanceParams{
		ID:      account.ID,
		Balance: utils.RandomMoney(),
	}

	updated_account, err := testQueries.UpdateAccountBalance(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updated_account)

	require.Equal(t, account.ID, updated_account.ID)
	require.Equal(t, account.Owner, updated_account.Owner)
	require.Equal(t, account.Currency, updated_account.Currency)
	require.Equal(t, arg.Balance, updated_account.Balance)
	require.WithinDuration(t, account.CreatedAt, updated_account.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)

	deleted_account, new_err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, new_err)
	require.EqualError(t, new_err, sql.ErrNoRows.Error())
	require.Empty(t, deleted_account)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
