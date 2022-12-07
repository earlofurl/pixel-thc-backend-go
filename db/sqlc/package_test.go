package db

import (
	"context"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"pixel-thc-backend-go/util"
	"testing"
	"time"
)

func createRandomPackage(t *testing.T) Package {
	newRandomPackageTag := createRandomPackageTag(t)
	newRandomUom := createRandomUom(t)

	arg := CreatePackageParams{
		TagID:                             nulls.NewInt64(newRandomPackageTag.ID),
		PackageType:                       util.RandomString(10),
		Quantity:                          decimal.NewFromFloat(util.RandomDecimalTimes100()),
		Notes:                             nulls.NewString(util.RandomString(10)),
		PackagedDateTime:                  time.Now(),
		HarvestDateTime:                   nulls.NewTime(time.Now()),
		LabTestingState:                   nulls.NewString(util.RandomString(10)),
		LabTestingStateDateTime:           nulls.NewTime(time.Now()),
		IsTradeSample:                     util.RandomBool(),
		IsTestingSample:                   util.RandomBool(),
		ProductRequiresRemediation:        util.RandomBool(),
		ContainsRemediatedProduct:         util.RandomBool(),
		RemediationDateTime:               nulls.NewTime(time.Now()),
		ReceivedDateTime:                  nulls.NewTime(time.Now()),
		ReceivedFromManifestNumber:        nulls.NewString(util.RandomString(10)),
		ReceivedFromFacilityLicenseNumber: nulls.NewString(util.RandomString(10)),
		ReceivedFromFacilityName:          nulls.NewString(util.RandomString(10)),
		IsOnHold:                          util.RandomBool(),
		ArchivedDate:                      nulls.NewTime(time.Now()),
		FinishedDate:                      nulls.NewTime(time.Now()),
		ItemID:                            nulls.NewInt64(util.RandomInt(1, 10)),
		ProvisionalLabel:                  nulls.NewString(util.RandomString(10)),
		IsProvisional:                     util.RandomBool(),
		IsSold:                            util.RandomBool(),
		PpuDefault:                        decimal.NewFromFloat(util.RandomPercent()),
		PpuOnOrder:                        decimal.NewFromFloat(util.RandomPercent()),
		TotalPackagePriceOnOrder:          decimal.NewFromFloat(util.RandomPercent()),
		PpuSoldPrice:                      decimal.NewFromFloat(util.RandomPercent()),
		TotalSoldPrice:                    decimal.NewFromFloat(util.RandomPercent()),
		PackagingSuppliesConsumed:         util.RandomBool(),
		IsLineItem:                        util.RandomBool(),
		//OrderID:                           nulls.NewInt64(0), // TODO: add back after Order model is created
		UomID: nulls.NewInt64(newRandomUom.ID),
	}
	pkg, err := testQueries.CreatePackage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pkg)

	// TODO: fix disabled require tests below

	require.Equal(t, arg.TagID, pkg.TagID)
	require.Equal(t, arg.PackageType, pkg.PackageType)
	//require.Equal(t, arg.Quantity, pkg.Quantity)
	require.Equal(t, arg.Notes, pkg.Notes)
	//require.Equal(t, arg.PackagedDateTime, pkg.PackagedDateTime)
	//require.Equal(t, arg.HarvestDateTime, pkg.HarvestDateTime)
	require.Equal(t, arg.LabTestingState, pkg.LabTestingState)
	//require.Equal(t, arg.LabTestingStateDateTime, pkg.LabTestingStateDateTime)
	require.Equal(t, arg.IsTradeSample, pkg.IsTradeSample)
	require.Equal(t, arg.IsTestingSample, pkg.IsTestingSample)
	require.Equal(t, arg.ProductRequiresRemediation, pkg.ProductRequiresRemediation)
	require.Equal(t, arg.ContainsRemediatedProduct, pkg.ContainsRemediatedProduct)
	//require.Equal(t, arg.RemediationDateTime, pkg.RemediationDateTime)
	//require.Equal(t, arg.ReceivedDateTime, pkg.ReceivedDateTime)
	require.Equal(t, arg.ReceivedFromManifestNumber, pkg.ReceivedFromManifestNumber)
	require.Equal(t, arg.ReceivedFromFacilityLicenseNumber, pkg.ReceivedFromFacilityLicenseNumber)
	require.Equal(t, arg.ReceivedFromFacilityName, pkg.ReceivedFromFacilityName)
	require.Equal(t, arg.IsOnHold, pkg.IsOnHold)
	//require.Equal(t, arg.ArchivedDate, pkg.ArchivedDate)
	//require.Equal(t, arg.FinishedDate, pkg.FinishedDate)
	require.Equal(t, arg.ItemID, pkg.ItemID)
	require.Equal(t, arg.ProvisionalLabel, pkg.ProvisionalLabel)
	require.Equal(t, arg.IsProvisional, pkg.IsProvisional)
	require.Equal(t, arg.IsSold, pkg.IsSold)
	//require.Equal(t, arg.PpuDefault, pkg.PpuDefault)
	//require.Equal(t, arg.PpuOnOrder, pkg.PpuOnOrder)
	//require.Equal(t, arg.TotalPackagePriceOnOrder, pkg.TotalPackagePriceOnOrder)
	//require.Equal(t, arg.PpuSoldPrice, pkg.PpuSoldPrice)
	//require.Equal(t, arg.TotalSoldPrice, pkg.TotalSoldPrice)
	require.Equal(t, arg.PackagingSuppliesConsumed, pkg.PackagingSuppliesConsumed)
	require.Equal(t, arg.IsLineItem, pkg.IsLineItem)
	//require.Equal(t, arg.OrderID, pkg.OrderID)
	require.Equal(t, arg.UomID, pkg.UomID)

	return pkg
}

func TestCreatePackage(t *testing.T) {
	createRandomPackage(t)
}

// TODO: Add the rest of the package tests
