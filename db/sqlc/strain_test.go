package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"pixel-thc-backend-go/util"
	"testing"
	"time"
)

func createRandomStrain(t *testing.T) Strain {
	strainName := util.RandomString(6)
	strainType := util.RandomString(3)

	arg := CreateStrainParams{
		Name: strainName,
		Type: strainType,
	}

	strain, err := testQueries.CreateStrain(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, strain)

	require.Equal(t, strainName, strain.Name)
	require.Equal(t, strainType, strain.Type)

	require.NotZero(t, strain.ID)
	require.NotZero(t, strain.CreatedAt)

	return strain
}

func TestCreateStrain(t *testing.T) {
	createRandomStrain(t)
}

func TestGetStrain(t *testing.T) {
	strain1 := createRandomStrain(t)
	strain2, err := testQueries.GetStrain(context.Background(), strain1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, strain2)

	require.Equal(t, strain1.ID, strain2.ID)
	require.Equal(t, strain1.Name, strain2.Name)
	require.Equal(t, strain1.Type, strain2.Type)
}

func TestListStrains(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomStrain(t)
	}

	strains, err := testQueries.ListStrains(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, strains)

	for _, strain := range strains {
		require.NotEmpty(t, strain)
	}
}

func TestUpdateStrain(t *testing.T) {
	strain1 := createRandomStrain(t)
	arg := UpdateStrainParams{
		ID:   strain1.ID,
		Name: util.RandomString(6),
		Type: util.RandomString(3),
	}
	strain2, err := testQueries.UpdateStrain(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, strain2)

	require.Equal(t, strain1.ID, strain2.ID)
	require.Equal(t, arg.Name, strain2.Name)
	require.Equal(t, arg.Type, strain2.Type)
	require.WithinDuration(t, strain1.CreatedAt, strain2.CreatedAt, time.Second)
}

func TestDeleteStrain(t *testing.T) {
	strain1 := createRandomStrain(t)
	err := testQueries.DeleteStrain(context.Background(), strain1.ID)
	require.NoError(t, err)

	strain2, err := testQueries.GetStrain(context.Background(), strain1.ID)
	require.Error(t, err)
	require.Empty(t, strain2)
}
