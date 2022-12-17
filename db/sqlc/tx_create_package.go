package db

import (
	"context"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
	"time"
)

// CreatePackageTxParams contains the input parameters for the createPackageTx function.
type CreatePackageTxParams struct {
	SourcePackageID                   nulls.Int64     `json:"source_package_id"`
	TagID                             nulls.Int64     `json:"tag_id"`
	PackageType                       string          `json:"package_type"`
	IsActive                          bool            `json:"is_active"`
	Quantity                          decimal.Decimal `json:"quantity"`
	Notes                             nulls.String    `json:"notes"`
	PackagedDateTime                  time.Time       `json:"packaged_date_time"`
	HarvestDateTime                   nulls.Time      `json:"harvest_date_time"`
	LabTestingState                   nulls.String    `json:"lab_testing_state"`
	LabTestingStateDateTime           nulls.Time      `json:"lab_testing_state_date_time"`
	IsTradeSample                     bool            `json:"is_trade_sample"`
	IsTestingSample                   bool            `json:"is_testing_sample"`
	ProductRequiresRemediation        bool            `json:"product_requires_remediation"`
	ContainsRemediatedProduct         bool            `json:"contains_remediated_product"`
	RemediationDateTime               nulls.Time      `json:"remediation_date_time"`
	ReceivedDateTime                  nulls.Time      `json:"received_date_time"`
	ReceivedFromManifestNumber        nulls.String    `json:"received_from_manifest_number"`
	ReceivedFromFacilityLicenseNumber nulls.String    `json:"received_from_facility_license_number"`
	ReceivedFromFacilityName          nulls.String    `json:"received_from_facility_name"`
	IsOnHold                          bool            `json:"is_on_hold"`
	ArchivedDate                      nulls.Time      `json:"archived_date"`
	FinishedDate                      nulls.Time      `json:"finished_date"`
	ItemID                            nulls.Int64     `json:"item_id"`
	ProvisionalLabel                  nulls.String    `json:"provisional_label"`
	IsProvisional                     bool            `json:"is_provisional"`
	IsSold                            bool            `json:"is_sold"`
	PpuDefault                        decimal.Decimal `json:"ppu_default"`
	PpuOnOrder                        decimal.Decimal `json:"ppu_on_order"`
	TotalPackagePriceOnOrder          decimal.Decimal `json:"total_package_price_on_order"`
	PpuSoldPrice                      decimal.Decimal `json:"ppu_sold_price"`
	TotalSoldPrice                    decimal.Decimal `json:"total_sold_price"`
	PackagingSuppliesConsumed         bool            `json:"packaging_supplies_consumed"`
	IsLineItem                        bool            `json:"is_line_item"`
	OrderID                           nulls.Int64     `json:"order_id"`
	UomID                             int64           `json:"uom_id"`
}

// CreatePackageTxResult contains the output parameters for the createPackageTx function.
type CreatePackageTxResult struct {
	CreatedPackage           Package                  `json:"created_package"`
	CreatePckgToPckgTxResult CreatePckgToPckgTxResult `json:"create_pckg_to_pckg_tx_result"`
}

// CreatePackageTx creates a new package, connects it to the source package and lab tests, and returns the new package ID.
func (store *SQLStore) CreatePackageTx(ctx context.Context, arg CreatePackageTxParams) (CreatePackageTxResult, error) {
	var result CreatePackageTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.CreatedPackage, err = q.CreatePackage(ctx, CreatePackageParams{
			TagID:                             arg.TagID,
			PackageType:                       arg.PackageType,
			IsActive:                          arg.IsActive,
			Quantity:                          arg.Quantity,
			Notes:                             arg.Notes,
			PackagedDateTime:                  arg.PackagedDateTime,
			HarvestDateTime:                   arg.HarvestDateTime,
			LabTestingState:                   arg.LabTestingState,
			LabTestingStateDateTime:           arg.LabTestingStateDateTime,
			IsTradeSample:                     arg.IsTradeSample,
			IsTestingSample:                   arg.IsTestingSample,
			ProductRequiresRemediation:        arg.ProductRequiresRemediation,
			ContainsRemediatedProduct:         arg.ContainsRemediatedProduct,
			RemediationDateTime:               arg.RemediationDateTime,
			ReceivedDateTime:                  arg.ReceivedDateTime,
			ReceivedFromManifestNumber:        arg.ReceivedFromManifestNumber,
			ReceivedFromFacilityLicenseNumber: arg.ReceivedFromFacilityLicenseNumber,
			ReceivedFromFacilityName:          arg.ReceivedFromFacilityName,
			IsOnHold:                          arg.IsOnHold,
			ArchivedDate:                      arg.ArchivedDate,
			FinishedDate:                      arg.FinishedDate,
			ItemID:                            arg.ItemID,
			ProvisionalLabel:                  arg.ProvisionalLabel,
			IsProvisional:                     arg.IsProvisional,
			IsSold:                            arg.IsSold,
			PpuDefault:                        arg.PpuDefault,
			PpuOnOrder:                        arg.PpuOnOrder,
			TotalPackagePriceOnOrder:          arg.TotalPackagePriceOnOrder,
			PpuSoldPrice:                      arg.PpuSoldPrice,
			TotalSoldPrice:                    arg.TotalSoldPrice,
			PackagingSuppliesConsumed:         arg.PackagingSuppliesConsumed,
			IsLineItem:                        arg.IsLineItem,
			OrderID:                           arg.OrderID,
			UomID:                             arg.UomID,
		})
		if err != nil {
			return err
		}
		return err
	})

	if arg.SourcePackageID != (nulls.Int64{Valid: false}) {
		err := store.execTx(ctx, func(q *Queries) error {
			var err error

			println("FromPackageID: ", arg.SourcePackageID.Int64)
			println("ToPackageID: ", result.CreatedPackage.ID)
			println("Quantity: ", arg.Quantity.String())
			println("UomID: ", arg.UomID)

			// nested tx_package_to_package transaction to connect the new package to the source package
			result.CreatePckgToPckgTxResult, err = store.CreatePckgToPckgTx(ctx, CreatePckgToPckgTxParams{
				FromPackageID: arg.SourcePackageID.Int64,
				ToPackageID:   result.CreatedPackage.ID,
				Amount:        arg.Quantity,
				UomID:         arg.UomID,
			})
			if err != nil {
				return err
			}
			return err
		})
		return result, err
	}

	return result, err

}
