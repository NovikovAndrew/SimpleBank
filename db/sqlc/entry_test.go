package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/NovikovAndrew/SimpleBank/util"
	"github.com/stretchr/testify/require"
)

func testCreateRandomEntry(t *testing.T) Entry {
	account1 := testCreateRandomAccount(t)
	amount := util.RandomMoney()

	entry, err := testQuery.CreateEntry(context.Background(), CreateEntryParams{
		AccountID: account1.ID,
		Amount:    amount,
	})

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account1.ID, entry.AccountID)
	require.Equal(t, amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	testCreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := testCreateRandomEntry(t)
	entry2, err := testQuery.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 20; i++ {
		testCreateRandomEntry(t)
	}

	args := GetEntriesParams{
		Limit:  10,
		Offset: 10,
	}

	entries, err := testQuery.GetEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 10)

	for _, entriy := range entries {
		require.NotEmpty(t, entriy)
	}
}

func TestDeleEntry(t *testing.T) {
	entry1 := testCreateRandomEntry(t)
	err := testQuery.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQuery.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
