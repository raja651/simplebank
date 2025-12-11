package db

import (
	"context"
	"testing"
	"time"

	"github.com/raja651/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount() (Account, CreateAccountParams, error) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	return account, arg, err

}
func TestCreateAccount(t *testing.T) {
	account, arg, err := createRandomAccount()
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	account1, _, _ := createRandomAccount()
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1, _, _ := createRandomAccount()

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(t.Context(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account1.ID, updatedAccount.ID)
	require.Equal(t, account1.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account1.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func DeleteAccount(t *testing.T) {
	account1, _, _ := createRandomAccount()

	err := testQueries.DeleteAccount(t.Context(), account1.ID)
	require.Error(t, err)
}
