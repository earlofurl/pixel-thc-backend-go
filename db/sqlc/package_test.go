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
	var lastRandomPackageTag PackageTag
	// create 10 random package tags
	for i := 0; i < 10; i++ {
		lastRandomPackageTag = createRandomPackageTag(t)
	}

	newRandomUom := createRandomUom(t)
	newRandomItem := createRandomItem(t)

	arg := CreatePackageParams{
		TagID:                             nulls.NewInt64(lastRandomPackageTag.ID),
		PackageType:                       util.RandomString(10),
		IsActive:                          util.RandomBool(),
		Quantity:                          decimal.NewFromFloatWithExponent(util.RandomDecimalTimes100(), -6),
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
		ItemID:                            nulls.NewInt64(newRandomItem.ID),
		ProvisionalLabel:                  nulls.NewString(util.RandomString(10)),
		IsProvisional:                     util.RandomBool(),
		IsSold:                            util.RandomBool(),
		PpuDefault:                        decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		PpuOnOrder:                        decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		TotalPackagePriceOnOrder:          decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		PpuSoldPrice:                      decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
		TotalSoldPrice:                    decimal.NewFromFloatWithExponent(util.RandomPercent(), -4),
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
	require.Equal(t, arg.Quantity, pkg.Quantity)
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
	require.Equal(t, arg.PpuDefault, pkg.PpuDefault)
	require.Equal(t, arg.PpuOnOrder, pkg.PpuOnOrder)
	require.Equal(t, arg.TotalPackagePriceOnOrder, pkg.TotalPackagePriceOnOrder)
	require.Equal(t, arg.PpuSoldPrice, pkg.PpuSoldPrice)
	require.Equal(t, arg.TotalSoldPrice, pkg.TotalSoldPrice)
	require.Equal(t, arg.PackagingSuppliesConsumed, pkg.PackagingSuppliesConsumed)
	require.Equal(t, arg.IsLineItem, pkg.IsLineItem)
	//require.Equal(t, arg.OrderID, pkg.OrderID)
	require.Equal(t, arg.UomID, pkg.UomID)

	return pkg
}

func TestQueries_TestCreatePackage(t *testing.T) {
	createRandomPackage(t)
}

func TestQueries_TestGetPackage(t *testing.T) {
	pkg1 := createRandomPackage(t)

	pkg2, err := testQueries.GetPackage(context.Background(), pkg1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pkg2)

	require.Equal(t, pkg1.ID, pkg2.ID)
	require.Equal(t, pkg1.TagID, pkg2.TagID)
	require.Equal(t, pkg1.PackageType, pkg2.PackageType)
	require.Equal(t, pkg1.Quantity, pkg2.Quantity)
	require.Equal(t, pkg1.Notes, pkg2.Notes)
	require.Equal(t, pkg1.PackagedDateTime, pkg2.PackagedDateTime)
	require.Equal(t, pkg1.HarvestDateTime, pkg2.HarvestDateTime)
	require.Equal(t, pkg1.LabTestingState, pkg2.LabTestingState)
	require.Equal(t, pkg1.LabTestingStateDateTime, pkg2.LabTestingStateDateTime)
	require.Equal(t, pkg1.IsTradeSample, pkg2.IsTradeSample)
	require.Equal(t, pkg1.IsTestingSample, pkg2.IsTestingSample)
	require.Equal(t, pkg1.ProductRequiresRemediation, pkg2.ProductRequiresRemediation)
	require.Equal(t, pkg1.ContainsRemediatedProduct, pkg2.ContainsRemediatedProduct)
	require.Equal(t, pkg1.RemediationDateTime, pkg2.RemediationDateTime)
	require.Equal(t, pkg1.ReceivedDateTime, pkg2.ReceivedDateTime)
	require.Equal(t, pkg1.ReceivedFromManifestNumber, pkg2.ReceivedFromManifestNumber)
	require.Equal(t, pkg1.ReceivedFromFacilityLicenseNumber, pkg2.ReceivedFromFacilityLicenseNumber)
	require.Equal(t, pkg1.ReceivedFromFacilityName, pkg2.ReceivedFromFacilityName)
	require.Equal(t, pkg1.IsOnHold, pkg2.IsOnHold)
	require.Equal(t, pkg1.ArchivedDate, pkg2.ArchivedDate)
	require.Equal(t, pkg1.FinishedDate, pkg2.FinishedDate)
	require.Equal(t, pkg1.ItemID, pkg2.ItemID)
	require.Equal(t, pkg1.ProvisionalLabel, pkg2.ProvisionalLabel)
	require.Equal(t, pkg1.IsProvisional, pkg2.IsProvisional)
	require.Equal(t, pkg1.IsSold, pkg2.IsSold)
	require.Equal(t, pkg1.PpuDefault, pkg2.PpuDefault)
	require.Equal(t, pkg1.PpuOnOrder, pkg2.PpuOnOrder)
	require.Equal(t, pkg1.TotalPackagePriceOnOrder, pkg2.TotalPackagePriceOnOrder)
	require.Equal(t, pkg1.PpuSoldPrice, pkg2.PpuSoldPrice)
	require.Equal(t, pkg1.TotalSoldPrice, pkg2.TotalSoldPrice)
	require.Equal(t, pkg1.PackagingSuppliesConsumed, pkg2.PackagingSuppliesConsumed)
	require.Equal(t, pkg1.IsLineItem, pkg2.IsLineItem)
	require.Equal(t, pkg1.OrderID, pkg2.OrderID)
	require.Equal(t, pkg1.UomID, pkg2.UomID)

}

//func TestQueries_ListPackagesWithTestResults(t *testing.T) {
//	var lastPackage Package
//	for i := 0; i < 10; i++ {
//		lastPackage = createRandomPackage(t)
//	}
//
//	var lastTestResult TestResult
//	for i := 0; i < 10; i++ {
//		lastTestResult = createRandomTestResult(t)
//	}
//
//
//	pkg2, err := testQueries.GetPackageWithTestResults(context.Background())
//	require.NoError(t, err)
//	require.NotEmpty(t, pkg2)
//
//	require.Equal(t, lastPackage.ID, pkg2.ID)
//	require.Equal(t, lastPackage.TagID, pkg2.TagID)
//	require.Equal(t, lastPackage.PackageType, pkg2.PackageType)
//	require.Equal(t, lastPackage.Quantity, pkg2.Quantity)
//	require.Equal(t, lastPackage.Notes, pkg2.Notes)
//	require.Equal(t, lastPackage.PackagedDateTime, pkg2.PackagedDateTime)
//	require.Equal(t, lastPackage.HarvestDateTime, pkg2.HarvestDateTime)
//	require.Equal(t, lastPackage.LabTestingState, pkg2.LabTestingState)
//	require.Equal(t, lastPackage.LabTestingStateDateTime, pkg2.LabTestingStateDateTime)
//	require.Equal(t, lastPackage.IsTradeSample, pkg2.IsTradeSample)
//	require.Equal(t, lastPackage.IsTestingSample, pkg2.IsTestingSample)
//	require.Equal(t, lastPackage.ProductRequiresRemediation, pkg2.ProductRequiresRemediation)
//	require.Equal(t, lastPackage.ContainsRemediatedProduct, pkg2.ContainsRemediatedProduct)
//	require.Equal(t, lastPackage.RemediationDateTime, pkg2.RemediationDateTime)
//	require.Equal(t, lastPackage.ReceivedDateTime, pkg2.ReceivedDateTime)
//	require.Equal(t, lastPackage.ReceivedFromManifestNumber, pkg2.ReceivedFromManifestNumber)
//	require.Equal(t, lastPackage.ReceivedFromFacilityLicenseNumber, pkg2.ReceivedFromFacilityLicenseNumber)
//	require.Equal(t, lastPackage.ReceivedFromFacilityName, pkg2.ReceivedFromFacilityName)
//	require.Equal(t, lastPackage.IsOnHold, pkg2.IsOnHold)
//	require.Equal(t, lastPackage.ArchivedDate, pkg2.ArchivedDate)
//	require.Equal(t, lastPackage.FinishedDate, pkg2.FinishedDate)
//	require.Equal(t, lastPackage.ItemID, pkg2.ItemID)
//	require.Equal(t, lastPackage.ProvisionalLabel, pkg2.ProvisionalLabel)
//	require.Equal(t, lastPackage.IsProvisional, pkg2.IsProvisional)
//	require.Equal(t, lastPackage.IsSold, pkg2.IsSold)
//	require.Equal(t, lastPackage.PpuDefault, pkg2.PpuDefault)
//	require.Equal(t, lastPackage.PpuOnOrder, pkg2.PpuOnOrder)
//	require.Equal(t, lastPackage.TotalPackagePriceOnOrder, pkg2.TotalPackagePriceOnOrder)
//	require.Equal(t, lastPackage.PpuSoldPrice, pkg2.PpuSoldPrice)
//	require.Equal(t, lastPackage.TotalSoldPrice, pkg2.TotalSoldPrice)
//	require.Equal(t, lastPackage.PackagingSuppliesConsumed, pkg2.PackagingSuppliesConsumed)
//	require.Equal(t, lastPackage.IsLineItem, pkg2.IsLineItem)
//	require.Equal(t, lastPackage.OrderID, pkg2.OrderID)
//	require.Equal(t, lastPackage.UomID, pkg2.UomID)
//
//}

// TODO: Add the rest of the package tests
