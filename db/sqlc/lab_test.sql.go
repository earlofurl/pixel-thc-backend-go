// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: lab_test.sql

package db

import (
	"context"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
)

const assignLabTestToPackage = `-- name: AssignLabTestToPackage :exec
INSERT INTO lab_tests_packages (lab_test_id, package_id) VALUES ($1, $2)
`

type AssignLabTestToPackageParams struct {
	LabTestID nulls.Int64 `json:"lab_test_id"`
	PackageID int64       `json:"package_id"`
}

// description: Assign a lab test to a package via junction table
func (q *Queries) AssignLabTestToPackage(ctx context.Context, arg AssignLabTestToPackageParams) error {
	_, err := q.db.ExecContext(ctx, assignLabTestToPackage, arg.LabTestID, arg.PackageID)
	return err
}

const createLabTest = `-- name: CreateLabTest :one
INSERT INTO lab_tests (test_name,
                       batch_code,
                       test_id_code,
                       lab_facility_name,
                       test_performed_date_time,
                       overall_passed,
                       test_type_name,
                       test_passed,
                       test_comment,
                       thc_total_percent,
                       thc_total_value,
                       cbd_percent,
                       cbd_value,
                       terpene_total_percent,
                       terpene_total_value,
                       thc_a_percent,
                       thc_a_value,
                       delta9_thc_percent,
                       delta9_thc_value,
                       delta8_thc_percent,
                       delta8_thc_value,
                       thc_v_percent,
                       thc_v_value,
                       cbd_a_percent,
                       cbd_a_value,
                       cbn_percent,
                       cbn_value,
                       cbg_a_percent,
                       cbg_a_value,
                       cbg_percent,
                       cbg_value,
                       cbc_percent,
                       cbc_value,
                       total_cannabinoid_percent,
                       total_cannabinoid_value)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24,
        $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35)
RETURNING id, created_at, updated_at, test_name, batch_code, test_id_code, lab_facility_name, test_performed_date_time, test_completed, overall_passed, test_type_name, test_passed, test_comment, thc_total_percent, thc_total_value, cbd_percent, cbd_value, terpene_total_percent, terpene_total_value, thc_a_percent, thc_a_value, delta9_thc_percent, delta9_thc_value, delta8_thc_percent, delta8_thc_value, thc_v_percent, thc_v_value, cbd_a_percent, cbd_a_value, cbn_percent, cbn_value, cbg_a_percent, cbg_a_value, cbg_percent, cbg_value, cbc_percent, cbc_value, total_cannabinoid_percent, total_cannabinoid_value
`

type CreateLabTestParams struct {
	TestName                string          `json:"test_name"`
	BatchCode               string          `json:"batch_code"`
	TestIDCode              string          `json:"test_id_code"`
	LabFacilityName         string          `json:"lab_facility_name"`
	TestPerformedDateTime   time.Time       `json:"test_performed_date_time"`
	OverallPassed           bool            `json:"overall_passed"`
	TestTypeName            string          `json:"test_type_name"`
	TestPassed              bool            `json:"test_passed"`
	TestComment             string          `json:"test_comment"`
	ThcTotalPercent         decimal.Decimal `json:"thc_total_percent"`
	ThcTotalValue           decimal.Decimal `json:"thc_total_value"`
	CbdPercent              decimal.Decimal `json:"cbd_percent"`
	CbdValue                decimal.Decimal `json:"cbd_value"`
	TerpeneTotalPercent     decimal.Decimal `json:"terpene_total_percent"`
	TerpeneTotalValue       decimal.Decimal `json:"terpene_total_value"`
	ThcAPercent             decimal.Decimal `json:"thc_a_percent"`
	ThcAValue               decimal.Decimal `json:"thc_a_value"`
	Delta9ThcPercent        decimal.Decimal `json:"delta9_thc_percent"`
	Delta9ThcValue          decimal.Decimal `json:"delta9_thc_value"`
	Delta8ThcPercent        decimal.Decimal `json:"delta8_thc_percent"`
	Delta8ThcValue          decimal.Decimal `json:"delta8_thc_value"`
	ThcVPercent             decimal.Decimal `json:"thc_v_percent"`
	ThcVValue               decimal.Decimal `json:"thc_v_value"`
	CbdAPercent             decimal.Decimal `json:"cbd_a_percent"`
	CbdAValue               decimal.Decimal `json:"cbd_a_value"`
	CbnPercent              decimal.Decimal `json:"cbn_percent"`
	CbnValue                decimal.Decimal `json:"cbn_value"`
	CbgAPercent             decimal.Decimal `json:"cbg_a_percent"`
	CbgAValue               decimal.Decimal `json:"cbg_a_value"`
	CbgPercent              decimal.Decimal `json:"cbg_percent"`
	CbgValue                decimal.Decimal `json:"cbg_value"`
	CbcPercent              decimal.Decimal `json:"cbc_percent"`
	CbcValue                decimal.Decimal `json:"cbc_value"`
	TotalCannabinoidPercent decimal.Decimal `json:"total_cannabinoid_percent"`
	TotalCannabinoidValue   decimal.Decimal `json:"total_cannabinoid_value"`
}

// description: Create a lab test
func (q *Queries) CreateLabTest(ctx context.Context, arg CreateLabTestParams) (LabTest, error) {
	row := q.db.QueryRowContext(ctx, createLabTest,
		arg.TestName,
		arg.BatchCode,
		arg.TestIDCode,
		arg.LabFacilityName,
		arg.TestPerformedDateTime,
		arg.OverallPassed,
		arg.TestTypeName,
		arg.TestPassed,
		arg.TestComment,
		arg.ThcTotalPercent,
		arg.ThcTotalValue,
		arg.CbdPercent,
		arg.CbdValue,
		arg.TerpeneTotalPercent,
		arg.TerpeneTotalValue,
		arg.ThcAPercent,
		arg.ThcAValue,
		arg.Delta9ThcPercent,
		arg.Delta9ThcValue,
		arg.Delta8ThcPercent,
		arg.Delta8ThcValue,
		arg.ThcVPercent,
		arg.ThcVValue,
		arg.CbdAPercent,
		arg.CbdAValue,
		arg.CbnPercent,
		arg.CbnValue,
		arg.CbgAPercent,
		arg.CbgAValue,
		arg.CbgPercent,
		arg.CbgValue,
		arg.CbcPercent,
		arg.CbcValue,
		arg.TotalCannabinoidPercent,
		arg.TotalCannabinoidValue,
	)
	var i LabTest
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TestName,
		&i.BatchCode,
		&i.TestIDCode,
		&i.LabFacilityName,
		&i.TestPerformedDateTime,
		&i.TestCompleted,
		&i.OverallPassed,
		&i.TestTypeName,
		&i.TestPassed,
		&i.TestComment,
		&i.ThcTotalPercent,
		&i.ThcTotalValue,
		&i.CbdPercent,
		&i.CbdValue,
		&i.TerpeneTotalPercent,
		&i.TerpeneTotalValue,
		&i.ThcAPercent,
		&i.ThcAValue,
		&i.Delta9ThcPercent,
		&i.Delta9ThcValue,
		&i.Delta8ThcPercent,
		&i.Delta8ThcValue,
		&i.ThcVPercent,
		&i.ThcVValue,
		&i.CbdAPercent,
		&i.CbdAValue,
		&i.CbnPercent,
		&i.CbnValue,
		&i.CbgAPercent,
		&i.CbgAValue,
		&i.CbgPercent,
		&i.CbgValue,
		&i.CbcPercent,
		&i.CbcValue,
		&i.TotalCannabinoidPercent,
		&i.TotalCannabinoidValue,
	)
	return i, err
}

const deleteLabTest = `-- name: DeleteLabTest :exec
DELETE
FROM lab_tests
WHERE id = $1
`

// description: Delete a lab test by ID
func (q *Queries) DeleteLabTest(ctx context.Context, id nulls.Int64) error {
	_, err := q.db.ExecContext(ctx, deleteLabTest, id)
	return err
}

const getLabTest = `-- name: GetLabTest :one
SELECT id, created_at, updated_at, test_name, batch_code, test_id_code, lab_facility_name, test_performed_date_time, test_completed, overall_passed, test_type_name, test_passed, test_comment, thc_total_percent, thc_total_value, cbd_percent, cbd_value, terpene_total_percent, terpene_total_value, thc_a_percent, thc_a_value, delta9_thc_percent, delta9_thc_value, delta8_thc_percent, delta8_thc_value, thc_v_percent, thc_v_value, cbd_a_percent, cbd_a_value, cbn_percent, cbn_value, cbg_a_percent, cbg_a_value, cbg_percent, cbg_value, cbc_percent, cbc_value, total_cannabinoid_percent, total_cannabinoid_value
FROM lab_tests
WHERE id = $1
LIMIT 1
`

// description: Get a lab test by ID
func (q *Queries) GetLabTest(ctx context.Context, id nulls.Int64) (LabTest, error) {
	row := q.db.QueryRowContext(ctx, getLabTest, id)
	var i LabTest
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TestName,
		&i.BatchCode,
		&i.TestIDCode,
		&i.LabFacilityName,
		&i.TestPerformedDateTime,
		&i.TestCompleted,
		&i.OverallPassed,
		&i.TestTypeName,
		&i.TestPassed,
		&i.TestComment,
		&i.ThcTotalPercent,
		&i.ThcTotalValue,
		&i.CbdPercent,
		&i.CbdValue,
		&i.TerpeneTotalPercent,
		&i.TerpeneTotalValue,
		&i.ThcAPercent,
		&i.ThcAValue,
		&i.Delta9ThcPercent,
		&i.Delta9ThcValue,
		&i.Delta8ThcPercent,
		&i.Delta8ThcValue,
		&i.ThcVPercent,
		&i.ThcVValue,
		&i.CbdAPercent,
		&i.CbdAValue,
		&i.CbnPercent,
		&i.CbnValue,
		&i.CbgAPercent,
		&i.CbgAValue,
		&i.CbgPercent,
		&i.CbgValue,
		&i.CbcPercent,
		&i.CbcValue,
		&i.TotalCannabinoidPercent,
		&i.TotalCannabinoidValue,
	)
	return i, err
}

const listLabTests = `-- name: ListLabTests :many
SELECT id, created_at, updated_at, test_name, batch_code, test_id_code, lab_facility_name, test_performed_date_time, test_completed, overall_passed, test_type_name, test_passed, test_comment, thc_total_percent, thc_total_value, cbd_percent, cbd_value, terpene_total_percent, terpene_total_value, thc_a_percent, thc_a_value, delta9_thc_percent, delta9_thc_value, delta8_thc_percent, delta8_thc_value, thc_v_percent, thc_v_value, cbd_a_percent, cbd_a_value, cbn_percent, cbn_value, cbg_a_percent, cbg_a_value, cbg_percent, cbg_value, cbc_percent, cbc_value, total_cannabinoid_percent, total_cannabinoid_value
FROM lab_tests
ORDER BY created_at
`

// description: List all lab tests
func (q *Queries) ListLabTests(ctx context.Context) ([]LabTest, error) {
	rows, err := q.db.QueryContext(ctx, listLabTests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []LabTest{}
	for rows.Next() {
		var i LabTest
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TestName,
			&i.BatchCode,
			&i.TestIDCode,
			&i.LabFacilityName,
			&i.TestPerformedDateTime,
			&i.TestCompleted,
			&i.OverallPassed,
			&i.TestTypeName,
			&i.TestPassed,
			&i.TestComment,
			&i.ThcTotalPercent,
			&i.ThcTotalValue,
			&i.CbdPercent,
			&i.CbdValue,
			&i.TerpeneTotalPercent,
			&i.TerpeneTotalValue,
			&i.ThcAPercent,
			&i.ThcAValue,
			&i.Delta9ThcPercent,
			&i.Delta9ThcValue,
			&i.Delta8ThcPercent,
			&i.Delta8ThcValue,
			&i.ThcVPercent,
			&i.ThcVValue,
			&i.CbdAPercent,
			&i.CbdAValue,
			&i.CbnPercent,
			&i.CbnValue,
			&i.CbgAPercent,
			&i.CbgAValue,
			&i.CbgPercent,
			&i.CbgValue,
			&i.CbcPercent,
			&i.CbcValue,
			&i.TotalCannabinoidPercent,
			&i.TotalCannabinoidValue,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLabTest = `-- name: UpdateLabTest :one
UPDATE lab_tests
SET test_name                 = $1,
    batch_code                = $2,
    test_id_code              = $3,
    lab_facility_name         = $4,
    test_performed_date_time  = $5,
    overall_passed            = $6,
    test_type_name            = $7,
    test_passed               = $8,
    test_comment              = $9,
    thc_total_percent         = $10,
    thc_total_value           = $11,
    cbd_percent               = $12,
    cbd_value                 = $13,
    terpene_total_percent     = $14,
    terpene_total_value       = $15,
    thc_a_percent             = $16,
    thc_a_value               = $17,
    delta9_thc_percent        = $18,
    delta9_thc_value          = $19,
    delta8_thc_percent        = $20,
    delta8_thc_value          = $21,
    thc_v_percent             = $22,
    thc_v_value               = $23,
    cbd_a_percent             = $24,
    cbd_a_value               = $25,
    cbn_percent               = $26,
    cbn_value                 = $27,
    cbg_a_percent             = $28,
    cbg_a_value               = $29,
    cbg_percent               = $30,
    cbg_value                 = $31,
    cbc_percent               = $32,
    cbc_value                 = $33,
    total_cannabinoid_percent = $34,
    total_cannabinoid_value   = $35
WHERE id = $36
RETURNING id, created_at, updated_at, test_name, batch_code, test_id_code, lab_facility_name, test_performed_date_time, test_completed, overall_passed, test_type_name, test_passed, test_comment, thc_total_percent, thc_total_value, cbd_percent, cbd_value, terpene_total_percent, terpene_total_value, thc_a_percent, thc_a_value, delta9_thc_percent, delta9_thc_value, delta8_thc_percent, delta8_thc_value, thc_v_percent, thc_v_value, cbd_a_percent, cbd_a_value, cbn_percent, cbn_value, cbg_a_percent, cbg_a_value, cbg_percent, cbg_value, cbc_percent, cbc_value, total_cannabinoid_percent, total_cannabinoid_value
`

type UpdateLabTestParams struct {
	TestName                string          `json:"test_name"`
	BatchCode               string          `json:"batch_code"`
	TestIDCode              string          `json:"test_id_code"`
	LabFacilityName         string          `json:"lab_facility_name"`
	TestPerformedDateTime   time.Time       `json:"test_performed_date_time"`
	OverallPassed           bool            `json:"overall_passed"`
	TestTypeName            string          `json:"test_type_name"`
	TestPassed              bool            `json:"test_passed"`
	TestComment             string          `json:"test_comment"`
	ThcTotalPercent         decimal.Decimal `json:"thc_total_percent"`
	ThcTotalValue           decimal.Decimal `json:"thc_total_value"`
	CbdPercent              decimal.Decimal `json:"cbd_percent"`
	CbdValue                decimal.Decimal `json:"cbd_value"`
	TerpeneTotalPercent     decimal.Decimal `json:"terpene_total_percent"`
	TerpeneTotalValue       decimal.Decimal `json:"terpene_total_value"`
	ThcAPercent             decimal.Decimal `json:"thc_a_percent"`
	ThcAValue               decimal.Decimal `json:"thc_a_value"`
	Delta9ThcPercent        decimal.Decimal `json:"delta9_thc_percent"`
	Delta9ThcValue          decimal.Decimal `json:"delta9_thc_value"`
	Delta8ThcPercent        decimal.Decimal `json:"delta8_thc_percent"`
	Delta8ThcValue          decimal.Decimal `json:"delta8_thc_value"`
	ThcVPercent             decimal.Decimal `json:"thc_v_percent"`
	ThcVValue               decimal.Decimal `json:"thc_v_value"`
	CbdAPercent             decimal.Decimal `json:"cbd_a_percent"`
	CbdAValue               decimal.Decimal `json:"cbd_a_value"`
	CbnPercent              decimal.Decimal `json:"cbn_percent"`
	CbnValue                decimal.Decimal `json:"cbn_value"`
	CbgAPercent             decimal.Decimal `json:"cbg_a_percent"`
	CbgAValue               decimal.Decimal `json:"cbg_a_value"`
	CbgPercent              decimal.Decimal `json:"cbg_percent"`
	CbgValue                decimal.Decimal `json:"cbg_value"`
	CbcPercent              decimal.Decimal `json:"cbc_percent"`
	CbcValue                decimal.Decimal `json:"cbc_value"`
	TotalCannabinoidPercent decimal.Decimal `json:"total_cannabinoid_percent"`
	TotalCannabinoidValue   decimal.Decimal `json:"total_cannabinoid_value"`
	ID                      nulls.Int64     `json:"id"`
}

// description: Update a lab test
func (q *Queries) UpdateLabTest(ctx context.Context, arg UpdateLabTestParams) (LabTest, error) {
	row := q.db.QueryRowContext(ctx, updateLabTest,
		arg.TestName,
		arg.BatchCode,
		arg.TestIDCode,
		arg.LabFacilityName,
		arg.TestPerformedDateTime,
		arg.OverallPassed,
		arg.TestTypeName,
		arg.TestPassed,
		arg.TestComment,
		arg.ThcTotalPercent,
		arg.ThcTotalValue,
		arg.CbdPercent,
		arg.CbdValue,
		arg.TerpeneTotalPercent,
		arg.TerpeneTotalValue,
		arg.ThcAPercent,
		arg.ThcAValue,
		arg.Delta9ThcPercent,
		arg.Delta9ThcValue,
		arg.Delta8ThcPercent,
		arg.Delta8ThcValue,
		arg.ThcVPercent,
		arg.ThcVValue,
		arg.CbdAPercent,
		arg.CbdAValue,
		arg.CbnPercent,
		arg.CbnValue,
		arg.CbgAPercent,
		arg.CbgAValue,
		arg.CbgPercent,
		arg.CbgValue,
		arg.CbcPercent,
		arg.CbcValue,
		arg.TotalCannabinoidPercent,
		arg.TotalCannabinoidValue,
		arg.ID,
	)
	var i LabTest
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TestName,
		&i.BatchCode,
		&i.TestIDCode,
		&i.LabFacilityName,
		&i.TestPerformedDateTime,
		&i.TestCompleted,
		&i.OverallPassed,
		&i.TestTypeName,
		&i.TestPassed,
		&i.TestComment,
		&i.ThcTotalPercent,
		&i.ThcTotalValue,
		&i.CbdPercent,
		&i.CbdValue,
		&i.TerpeneTotalPercent,
		&i.TerpeneTotalValue,
		&i.ThcAPercent,
		&i.ThcAValue,
		&i.Delta9ThcPercent,
		&i.Delta9ThcValue,
		&i.Delta8ThcPercent,
		&i.Delta8ThcValue,
		&i.ThcVPercent,
		&i.ThcVValue,
		&i.CbdAPercent,
		&i.CbdAValue,
		&i.CbnPercent,
		&i.CbnValue,
		&i.CbgAPercent,
		&i.CbgAValue,
		&i.CbgPercent,
		&i.CbgValue,
		&i.CbcPercent,
		&i.CbcValue,
		&i.TotalCannabinoidPercent,
		&i.TotalCannabinoidValue,
	)
	return i, err
}
