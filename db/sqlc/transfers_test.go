package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAccountID int64, toAccountID int64) Transfer {

	arg := CreateTransfersParams{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}
