package db

import "context"

// CreatePckgToPckgTxParams contains the input parameters for the createPckgToPckgTx function.
type CreatePckgToPckgTxParams struct {
	FromPackageID int64 `json:"from_package_id"`
	ToPackageID   int64 `json:"to_package_id"`
	Amount        int64 `json:"amount"`
	UomID         int64 `json:"uom_id"`
}

// CreatePckgToPckgTxResult is the result of the createPckgToPckgTx function.
type CreatePckgToPckgTxResult struct {
	PackageAdjustment PackageAdjustment `json:"package_adjustment"`
	FromPackage       Package           `json:"from_package"`
	ToPackage         Package           `json:"to_package"`
	FromEntry         Entry             `json:"from_entry"`
	ToEntry           Entry             `json:"to_entry"`
}

// CreatePckgToPckgTx performs a package adjustment from one package to the other.
// It creates the package adjustment, add package entries, and update packages' balance within a database transaction
func (store *SQLStore) CreatePckgToPckgTx(ctx context.Context, arg CreatePckgToPckgTxParams) (CreatePckgToPckgTxResult, error) {
	var result CreatePckgToPckgTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.PackageAdjustment, err = q.CreatePackageAdjustment(ctx, CreatePackageAdjustmentParams{
			FromPackageID: arg.FromPackageID,
			ToPackageID:   arg.ToPackageID,
			Amount:        arg.Amount,
			UomID:         arg.UomID,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			PackageID: arg.FromPackageID,
			Amount:    -arg.Amount,
			UomID:     arg.UomID,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			PackageID: arg.ToPackageID,
			Amount:    arg.Amount,
			UomID:     arg.UomID,
		})
		if err != nil {
			return err
		}

		if arg.FromPackageID < arg.ToPackageID {
			result.FromPackage, result.ToPackage, err = addPckgQty(ctx, q, arg.FromPackageID, -arg.Amount, arg.ToPackageID, arg.Amount)
		} else {
			result.ToPackage, result.FromPackage, err = addPckgQty(ctx, q, arg.ToPackageID, arg.Amount, arg.FromPackageID, -arg.Amount)
		}

		return err
	})

	return result, err
}

func addPckgQty(ctx context.Context, q *Queries, fromPckgID, fromAmount, toPckgID, toAmount int64) (fromPckg, toPckg Package, err error) {
	fromPckg, err = q.AddPackageQuantity(ctx, AddPackageQuantityParams{
		ID:     fromPckgID,
		Amount: fromAmount,
	})
	if err != nil {
		return fromPckg, toPckg, err
	}

	toPckg, err = q.AddPackageQuantity(ctx, AddPackageQuantityParams{
		ID:     toPckgID,
		Amount: toAmount,
	})
	return fromPckg, toPckg, err
}
