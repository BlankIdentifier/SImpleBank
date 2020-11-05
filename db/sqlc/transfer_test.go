package db

import (
	"context"
	"testing"

	"github.com/BlankIdentifier/SImpleBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountFrom, accountTo Account) Transfer{

	arg := CreateTransferParams {
		FromAccountID: accountFrom.ID,
		ToAccountID: accountTo.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, account1, account2)

	res, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, transfer.ID, res.ID)
	require.Equal(t, transfer.FromAccountID, res.FromAccountID)
	require.Equal(t, transfer.ToAccountID, res.ToAccountID)
	require.Equal(t, transfer.CreatedAt, res.CreatedAt)

}

func TestListTransfer(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i:=0; i <10 ; i++ {
		createRandomTransfer(t, account1, account2)
	}

	arg := ListTransfersParams {
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)


	for _, transfer := range transfers{
		require.NotEmpty(t, transfer)
		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	}
}