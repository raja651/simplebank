package db

import (
	"testing"

	"github.com/raja651/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) (Entry, error) {
	account, err := testQueries.GetAccount(t.Context(), util.RandomInt(1, 6))

	if err != nil {
		return Entry{}, err
	}

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(t.Context(), arg)

	return entry, err
}

func TestCreateEntry(t *testing.T) {
	entry, err := createRandomEntry(t)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
}

func TestGetEntry(t *testing.T) {
	entry, err := testQueries.GetEntry(t.Context(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
}
