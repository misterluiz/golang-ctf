package db

import (
	"context"
	"testing"

	"github.com/misterluiz/golang-ctf/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Tytle:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}
	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Tytle, category.Tytle)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Tytle, category2.Tytle)
	require.Equal(t, category1.Type, category2.Type)
	require.NotEmpty(t, category2.CreatedAt)

}
func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	err := testQueries.DeleteCategories(context.Background(), category.ID)
	require.NoError(t, err)

}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	arg := UpdateCategoriesParams{
		ID:          category1.ID,
		Tytle:       util.RandomString(12),
		Description: util.RandomString(20),
	}
	category2, err := testQueries.UpdateCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)
	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Tytle, category2.Tytle)
	require.Equal(t, arg.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)

}
func TestListCategories(t *testing.T) {

	lastCategory := createRandomCategory(t)

	arg := GetCategoriesParams{
		UserID:      lastCategory.UserID,
		Tytle:       lastCategory.Tytle,
		Type:        lastCategory.Type,
		Description: lastCategory.Description,
	}
	categorys, err := testQueries.GetCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categorys)

	for _, category := range categorys {
		require.Equal(t, lastCategory.UserID, category.UserID)
		require.Equal(t, lastCategory.ID, category.ID)
		require.Equal(t, lastCategory.Tytle, category.Tytle)
		require.Equal(t, lastCategory.Description, category.Description)
		require.NotEmpty(t, lastCategory.CreatedAt)
	}

}
