package db

import (
	"context"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
)

// CreatePackageTxParams contains the input parameters for the createPackageTx function.
type CreatePackageTxParams struct {
	SourcePackageID     nulls.Int64
	CreatePackageParams CreatePackageParams
}

// CreatePackageTxResult contains the output parameters for the createPackageTx function.
type CreatePackageTxResult struct {
	CreatedPackage                 Package                    `json:"created_package"`
	PackageAdjustment              PackageAdjustment          `json:"package_adjustment"`
	FromPackage                    Package                    `json:"from_package"`
	ToPackage                      Package                    `json:"to_package"`
	FromPackageAdjEntry            PackageAdjEntry            `json:"from_entry"`
	ToPackageAdjEntry              PackageAdjEntry            `json:"to_entry"`
	SourcePackageChildPackageEntry SourcePackagesChildPackage `json:"source_packages_child_package"`
}

// CreatePackageTx creates a new package, and optionally creates a package adjustment and/or a package to package relationship.
func (store *SQLStore) CreatePackageTx(ctx context.Context, arg CreatePackageTxParams) (CreatePackageTxResult, error) {
	var result CreatePackageTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		// Create the package
		pkg, err := q.CreatePackage(ctx, arg.CreatePackageParams)
		if err != nil {
			return err
		}
		result.CreatedPackage = pkg

		// Create a package adjustment
		pkgAdj, err := q.CreatePackageAdjustment(ctx, CreatePackageAdjustmentParams{
			FromPackageID: arg.SourcePackageID.Int64,
			ToPackageID:   pkg.ID,
			Amount:        arg.CreatePackageParams.Quantity,
			UomID:         arg.CreatePackageParams.UomID,
		})
		if err != nil {
			return err
		}
		result.PackageAdjustment = pkgAdj

		// Create package adjustment entries
		fromPckgAdjEntry, err := q.CreatePackageAdjEntry(ctx, CreatePackageAdjEntryParams{
			PackageID: arg.SourcePackageID.Int64,
			Amount:    decimal.NewFromFloat(-1).Mul(arg.CreatePackageParams.Quantity),
			UomID:     arg.CreatePackageParams.UomID,
		})
		if err != nil {
			return err
		}
		result.FromPackageAdjEntry = fromPckgAdjEntry

		toPckgAdjEntry, err := q.CreatePackageAdjEntry(ctx, CreatePackageAdjEntryParams{
			PackageID: pkg.ID,
			Amount:    arg.CreatePackageParams.Quantity,
			UomID:     arg.CreatePackageParams.UomID,
		})
		if err != nil {
			return err
		}
		result.ToPackageAdjEntry = toPckgAdjEntry

		sourcePackageChildPackageEntry, err := q.AssignSourcePackageChildPackage(ctx, AssignSourcePackageChildPackageParams{
			SourcePackageID: arg.SourcePackageID.Int64,
			ChildPackageID:  result.CreatedPackage.ID,
		})
		if err != nil {
			return err
		}
		result.SourcePackageChildPackageEntry = sourcePackageChildPackageEntry

		// Make the quantity transfer
		if arg.SourcePackageID.Int64 < pkg.ID {
			result.FromPackage, result.ToPackage, err = addPckgQty(ctx, q, arg.SourcePackageID.Int64, decimal.NewFromFloat(-1).Mul(arg.CreatePackageParams.Quantity), pkg.ID, arg.CreatePackageParams.Quantity)
		} else {
			result.ToPackage, result.FromPackage, err = addPckgQty(ctx, q, pkg.ID, arg.CreatePackageParams.Quantity, arg.SourcePackageID.Int64, decimal.NewFromFloat(-1).Mul(arg.CreatePackageParams.Quantity))
		}

		// Create Lab Test Package Assignment Entry
		err = q.AssignLabTestToPackage(ctx, AssignLabTestToPackageParams{
			LabTestID: nulls.NewInt64(1), // TODO: Get the lab test ID from the parent package
			PackageID: pkg.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
