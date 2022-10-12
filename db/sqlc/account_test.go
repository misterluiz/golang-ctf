package db

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/misterluiz/golang-ctf/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	category := createRandomCategory(t)
	arg := CreateAccountParams{
		UserID:      category.UserID,
		CategoryID:  category.ID,
		Tytle:       util.RandomString(12),
		Type:        category.Type,
		Description: util.RandomString(20),
		Value:       10,
		Date:        time.Now(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.CategoryID, account.CategoryID)
	require.Equal(t, arg.Value, account.Value)
	require.Equal(t, arg.Tytle, account.Tytle)
	require.Equal(t, arg.Type, account.Type)
	require.Equal(t, arg.Description, account.Description)

	require.NotEmpty(t, account.CreatedAt)
	require.NotEmpty(t, account.Date)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.CategoryID, account2.CategoryID)
	require.Equal(t, account1.Value, account2.Value)
	require.Equal(t, account1.Tytle, account2.Tytle)
	require.Equal(t, account1.Type, account2.Type)
	require.NotEmpty(t, account2.Date)
	require.NotEmpty(t, account2.CreatedAt)

}
func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:          account1.ID,
		Tytle:       util.RandomString(12),
		Description: util.RandomString(20),
		Value:       15,
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Tytle, account2.Tytle)
	require.Equal(t, arg.Description, account2.Description)
	require.Equal(t, arg.Value, account2.Value)
	require.NotEmpty(t, account1.CreatedAt, account2.CreatedAt)

}
func TestListAccounts(t *testing.T) {

	lastAccount := createRandomAccount(t)

	arg := GetAccountsParams{
		UserID:      lastAccount.UserID,
		Tytle:       lastAccount.Tytle,
		CategoryID:  lastAccount.CategoryID,
		Date:        lastAccount.Date,
		Type:        lastAccount.Type,
		Description: lastAccount.Description,
	}
	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.Equal(t, lastAccount.UserID, account.UserID)
		require.Equal(t, lastAccount.ID, account.ID)
		require.Equal(t, lastAccount.Tytle, account.Tytle)
		require.Equal(t, lastAccount.Value, account.Value)
		require.Equal(t, lastAccount.Description, account.Description)
		require.NotEmpty(t, lastAccount.Date)
		require.NotEmpty(t, lastAccount.CreatedAt)
		log.Fatal("account category title", account.CategoryTytle)
	}

}

func TestListGetReports(t *testing.T) {

	lastAccount := createRandomAccount(t)

	arg := GetAccountsReportsParams{
		UserID: lastAccount.UserID,
		Type:   lastAccount.Type,
	}
	reportValues, err := testQueries.GetAccountsReports(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reportValues)

}

func TestListGetGraph(t *testing.T) {

	lastAccount := createRandomAccount(t)

	arg := GetAccountsGraphParams{
		UserID: lastAccount.UserID,
		Type:   lastAccount.Type,
	}
	graphValue, err := testQueries.GetAccountsGraph(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, graphValue)

}
