package db

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	// run n concurrent transfer transaction
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check balances
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}

func TestQueries_PackageToPackageTx(t *testing.T) {
	store := NewStore(testDB)

	package1 := createRandomPackage(t)
	package2 := createRandomPackage(t)
	labTest1 := createRandomLabTest(t)
	fmt.Println(">> before:", package1.Quantity, package2.Quantity)

	n := 5
	amount := decimal.NewFromFloatWithExponent(10.0, -6)

	errs := make(chan error)
	results := make(chan CreatePckgToPckgTxResult)

	// run n concurrent package to package transfer transaction
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreatePckgToPckgTx(context.Background(), CreatePckgToPckgTxParams{
				FromPackageID: package1.ID,
				ToPackageID:   package2.ID,
				Amount:        amount,
				UomID:         package1.UomID,
				LabTestID:     labTest1.ID,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check package adjustment
		pckgAdjustment := result.PackageAdjustment
		require.NotEmpty(t, pckgAdjustment)
		require.Equal(t, package1.ID, pckgAdjustment.FromPackageID)
		require.Equal(t, package2.ID, pckgAdjustment.ToPackageID)
		// TODO: Fix the decimal types here.
		//require.Equal(t, amount, pckgAdjustment.Amount)
		require.NotZero(t, pckgAdjustment.ID)
		require.NotZero(t, pckgAdjustment.CreatedAt)

		_, err = store.GetTransfer(context.Background(), pckgAdjustment.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromPackageAdjEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, package1.ID, fromEntry.PackageID)
		//require.Equal(t, amount.Neg(), fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToPackageAdjEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, package2.ID, toEntry.PackageID)
		//require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetPackageAdjEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check packages
		fromPackage := result.FromPackage
		require.NotEmpty(t, fromPackage)
		require.Equal(t, package1.ID, fromPackage.ID)

		toPackage := result.ToPackage
		require.NotEmpty(t, toPackage)
		require.Equal(t, package2.ID, toPackage.ID)

		// check the diff of balance
		diff1 := package1.Quantity.Sub(fromPackage.Quantity)
		diff2 := toPackage.Quantity.Sub(package2.Quantity)
		require.Equal(t, diff1, diff2)
		//require.Equal(t, amount, diff1)

		k := diff1.Div(amount)
		require.True(t, k.GreaterThanOrEqual(decimal.NewFromInt(1)) && k.LessThanOrEqual(decimal.NewFromInt(int64(n))))
		require.NotContains(t, existed, k)
		existed[int(k.IntPart())] = true
	}

	// check the final updated balance
	updatedPackage1, err := store.GetPackage(context.Background(), package1.ID)
	require.NoError(t, err)

	updatedPackage2, err := store.GetPackage(context.Background(), package2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedPackage1.Quantity, updatedPackage2.Quantity)
	require.Equal(t, package1.Quantity.Sub(amount.Mul(decimal.NewFromInt(int64(n)))), updatedPackage1.Quantity)
	require.Equal(t, package2.Quantity.Add(amount.Mul(decimal.NewFromInt(int64(n)))), updatedPackage2.Quantity)
}
