package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"pixel-thc-backend-go/util"
	"testing"
	"time"
)

func createRandomPackageTag(t *testing.T) PackageTag {
	arg := CreatePackageTagParams{
		TagNumber:     util.RandomTagNumber(),
		IsAssigned:    util.RandomBool(),
		IsProvisional: util.RandomBool(),
		IsActive:      util.RandomBool(),
		//AssignedPackageID: util.RandomInt(1, 10), // TODO: change this to use a package created as part of the test.
	}

	packagetag, err := testQueries.CreatePackageTag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag)

	require.Equal(t, arg.TagNumber, packagetag.TagNumber)
	require.Equal(t, arg.IsAssigned, packagetag.IsAssigned)
	require.Equal(t, arg.IsProvisional, packagetag.IsProvisional)
	require.Equal(t, arg.IsActive, packagetag.IsActive)
	require.NotZero(t, packagetag.CreatedAt)

	return packagetag
}

func TestCreatePackageTag(t *testing.T) {
	createRandomPackageTag(t)
}

func TestGetPackageTag(t *testing.T) {
	packagetag1 := createRandomPackageTag(t)
	packagetag2, err := testQueries.GetPackageTag(context.Background(), packagetag1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag2)

	require.Equal(t, packagetag1.ID, packagetag2.ID)
	require.Equal(t, packagetag1.TagNumber, packagetag2.TagNumber)
	require.Equal(t, packagetag1.IsAssigned, packagetag2.IsAssigned)
	require.Equal(t, packagetag1.IsProvisional, packagetag2.IsProvisional)
	require.Equal(t, packagetag1.IsActive, packagetag2.IsActive)
	require.WithinDuration(t, packagetag1.CreatedAt, packagetag2.CreatedAt, time.Second)
}

func TestListPackageTags(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPackageTag(t)
	}

	packagetags, err := testQueries.ListPackageTags(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, packagetags)

	for _, packagetag := range packagetags {
		require.NotEmpty(t, packagetag)
	}
}

func TestQueries_GetPackageTagByTagNumber(t *testing.T) {
	packagetag1 := createRandomPackageTag(t)
	packagetag2, err := testQueries.GetPackageTagByTagNumber(context.Background(), packagetag1.TagNumber)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag2)

	require.Equal(t, packagetag1.ID, packagetag2.ID)
	require.Equal(t, packagetag1.TagNumber, packagetag2.TagNumber)
	require.Equal(t, packagetag1.IsAssigned, packagetag2.IsAssigned)
	require.Equal(t, packagetag1.IsProvisional, packagetag2.IsProvisional)
	require.Equal(t, packagetag1.IsActive, packagetag2.IsActive)
	require.WithinDuration(t, packagetag1.CreatedAt, packagetag2.CreatedAt, time.Second)

}

func TestUpdatePackageTag(t *testing.T) {
	packagetag1 := createRandomPackageTag(t)

	arg := UpdatePackageTagParams{
		ID:            packagetag1.ID,
		IsAssigned:    util.RandomBool(),
		IsProvisional: util.RandomBool(),
		IsActive:      util.RandomBool(),
	}

	packagetag2, err := testQueries.UpdatePackageTag(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, packagetag2)

	require.Equal(t, packagetag1.ID, packagetag2.ID)
	require.Equal(t, arg.IsAssigned, packagetag2.IsAssigned)
	require.Equal(t, arg.IsProvisional, packagetag2.IsProvisional)
	require.Equal(t, arg.IsActive, packagetag2.IsActive)
	require.Equal(t, packagetag1.AssignedPackageID, packagetag2.AssignedPackageID)
	require.WithinDuration(t, packagetag1.CreatedAt, packagetag2.CreatedAt, time.Second)
}

func TestDeletePackageTag(t *testing.T) {
	packagetag1 := createRandomPackageTag(t)

	err := testQueries.DeletePackageTag(context.Background(), packagetag1.ID)
	require.NoError(t, err)

	packagetag2, err := testQueries.GetPackageTag(context.Background(), packagetag1.ID)
	require.Error(t, err)
	require.Empty(t, packagetag2)
}
