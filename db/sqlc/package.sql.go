// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: package.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
)

const createPackage = `-- name: CreatePackage :one
INSERT INTO packages (tag_id,
                      package_type,
                      is_active,
                      quantity,
                      notes,
                      packaged_date_time,
                      harvest_date_time,
                      lab_testing_state,
                      lab_testing_state_date_time,
                      is_trade_sample,
                      is_testing_sample,
                      product_requires_remediation,
                      contains_remediated_product,
                      remediation_date_time,
                      received_date_time,
                      received_from_manifest_number,
                      received_from_facility_license_number,
                      received_from_facility_name,
                      is_on_hold,
                      archived_date,
                      finished_date,
                      item_id,
                      provisional_label,
                      is_provisional,
                      is_sold,
                      ppu_default,
                      ppu_on_order,
                      total_package_price_on_order,
                      ppu_sold_price,
                      total_sold_price,
                      packaging_supplies_consumed,
                      is_line_item,
                      order_id,
                      uom_id)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21,
        $22,
        $23,
        $24,
        $25,
        $26,
        $27,
        $28,
        $29,
        $30,
        $31,
        $32,
        $33,
        $34)
RETURNING id, created_at, updated_at, tag_id, package_type, is_active, quantity, notes, packaged_date_time, harvest_date_time, lab_testing_state, lab_testing_state_date_time, is_trade_sample, is_testing_sample, product_requires_remediation, contains_remediated_product, remediation_date_time, received_date_time, received_from_manifest_number, received_from_facility_license_number, received_from_facility_name, is_on_hold, archived_date, finished_date, item_id, provisional_label, is_provisional, is_sold, ppu_default, ppu_on_order, total_package_price_on_order, ppu_sold_price, total_sold_price, packaging_supplies_consumed, is_line_item, order_id, uom_id
`

type CreatePackageParams struct {
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
	UomID                             nulls.Int64     `json:"uom_id"`
}

// description: Create a package
func (q *Queries) CreatePackage(ctx context.Context, arg CreatePackageParams) (Package, error) {
	row := q.db.QueryRowContext(ctx, createPackage,
		arg.TagID,
		arg.PackageType,
		arg.IsActive,
		arg.Quantity,
		arg.Notes,
		arg.PackagedDateTime,
		arg.HarvestDateTime,
		arg.LabTestingState,
		arg.LabTestingStateDateTime,
		arg.IsTradeSample,
		arg.IsTestingSample,
		arg.ProductRequiresRemediation,
		arg.ContainsRemediatedProduct,
		arg.RemediationDateTime,
		arg.ReceivedDateTime,
		arg.ReceivedFromManifestNumber,
		arg.ReceivedFromFacilityLicenseNumber,
		arg.ReceivedFromFacilityName,
		arg.IsOnHold,
		arg.ArchivedDate,
		arg.FinishedDate,
		arg.ItemID,
		arg.ProvisionalLabel,
		arg.IsProvisional,
		arg.IsSold,
		arg.PpuDefault,
		arg.PpuOnOrder,
		arg.TotalPackagePriceOnOrder,
		arg.PpuSoldPrice,
		arg.TotalSoldPrice,
		arg.PackagingSuppliesConsumed,
		arg.IsLineItem,
		arg.OrderID,
		arg.UomID,
	)
	var i Package
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagID,
		&i.PackageType,
		&i.IsActive,
		&i.Quantity,
		&i.Notes,
		&i.PackagedDateTime,
		&i.HarvestDateTime,
		&i.LabTestingState,
		&i.LabTestingStateDateTime,
		&i.IsTradeSample,
		&i.IsTestingSample,
		&i.ProductRequiresRemediation,
		&i.ContainsRemediatedProduct,
		&i.RemediationDateTime,
		&i.ReceivedDateTime,
		&i.ReceivedFromManifestNumber,
		&i.ReceivedFromFacilityLicenseNumber,
		&i.ReceivedFromFacilityName,
		&i.IsOnHold,
		&i.ArchivedDate,
		&i.FinishedDate,
		&i.ItemID,
		&i.ProvisionalLabel,
		&i.IsProvisional,
		&i.IsSold,
		&i.PpuDefault,
		&i.PpuOnOrder,
		&i.TotalPackagePriceOnOrder,
		&i.PpuSoldPrice,
		&i.TotalSoldPrice,
		&i.PackagingSuppliesConsumed,
		&i.IsLineItem,
		&i.OrderID,
		&i.UomID,
	)
	return i, err
}

const deletePackage = `-- name: DeletePackage :exec
DELETE
FROM packages
WHERE id = $1
`

// description: Delete a package
func (q *Queries) DeletePackage(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePackage, id)
	return err
}

const getPackage = `-- name: GetPackage :one
SELECT id, created_at, updated_at, tag_id, package_type, is_active, quantity, notes, packaged_date_time, harvest_date_time, lab_testing_state, lab_testing_state_date_time, is_trade_sample, is_testing_sample, product_requires_remediation, contains_remediated_product, remediation_date_time, received_date_time, received_from_manifest_number, received_from_facility_license_number, received_from_facility_name, is_on_hold, archived_date, finished_date, item_id, provisional_label, is_provisional, is_sold, ppu_default, ppu_on_order, total_package_price_on_order, ppu_sold_price, total_sold_price, packaging_supplies_consumed, is_line_item, order_id, uom_id
FROM packages
WHERE id = $1
LIMIT 1
`

// description: Get a package
func (q *Queries) GetPackage(ctx context.Context, id int64) (Package, error) {
	row := q.db.QueryRowContext(ctx, getPackage, id)
	var i Package
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagID,
		&i.PackageType,
		&i.IsActive,
		&i.Quantity,
		&i.Notes,
		&i.PackagedDateTime,
		&i.HarvestDateTime,
		&i.LabTestingState,
		&i.LabTestingStateDateTime,
		&i.IsTradeSample,
		&i.IsTestingSample,
		&i.ProductRequiresRemediation,
		&i.ContainsRemediatedProduct,
		&i.RemediationDateTime,
		&i.ReceivedDateTime,
		&i.ReceivedFromManifestNumber,
		&i.ReceivedFromFacilityLicenseNumber,
		&i.ReceivedFromFacilityName,
		&i.IsOnHold,
		&i.ArchivedDate,
		&i.FinishedDate,
		&i.ItemID,
		&i.ProvisionalLabel,
		&i.IsProvisional,
		&i.IsSold,
		&i.PpuDefault,
		&i.PpuOnOrder,
		&i.TotalPackagePriceOnOrder,
		&i.PpuSoldPrice,
		&i.TotalSoldPrice,
		&i.PackagingSuppliesConsumed,
		&i.IsLineItem,
		&i.OrderID,
		&i.UomID,
	)
	return i, err
}

const getPackageByTagID = `-- name: GetPackageByTagID :one
SELECT id, created_at, updated_at, tag_id, package_type, is_active, quantity, notes, packaged_date_time, harvest_date_time, lab_testing_state, lab_testing_state_date_time, is_trade_sample, is_testing_sample, product_requires_remediation, contains_remediated_product, remediation_date_time, received_date_time, received_from_manifest_number, received_from_facility_license_number, received_from_facility_name, is_on_hold, archived_date, finished_date, item_id, provisional_label, is_provisional, is_sold, ppu_default, ppu_on_order, total_package_price_on_order, ppu_sold_price, total_sold_price, packaging_supplies_consumed, is_line_item, order_id, uom_id
FROM packages
WHERE tag_id = $1
LIMIT 1
`

// description: Get a package by tag id
func (q *Queries) GetPackageByTagID(ctx context.Context, tagID nulls.Int64) (Package, error) {
	row := q.db.QueryRowContext(ctx, getPackageByTagID, tagID)
	var i Package
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagID,
		&i.PackageType,
		&i.IsActive,
		&i.Quantity,
		&i.Notes,
		&i.PackagedDateTime,
		&i.HarvestDateTime,
		&i.LabTestingState,
		&i.LabTestingStateDateTime,
		&i.IsTradeSample,
		&i.IsTestingSample,
		&i.ProductRequiresRemediation,
		&i.ContainsRemediatedProduct,
		&i.RemediationDateTime,
		&i.ReceivedDateTime,
		&i.ReceivedFromManifestNumber,
		&i.ReceivedFromFacilityLicenseNumber,
		&i.ReceivedFromFacilityName,
		&i.IsOnHold,
		&i.ArchivedDate,
		&i.FinishedDate,
		&i.ItemID,
		&i.ProvisionalLabel,
		&i.IsProvisional,
		&i.IsSold,
		&i.PpuDefault,
		&i.PpuOnOrder,
		&i.TotalPackagePriceOnOrder,
		&i.PpuSoldPrice,
		&i.TotalSoldPrice,
		&i.PackagingSuppliesConsumed,
		&i.IsLineItem,
		&i.OrderID,
		&i.UomID,
	)
	return i, err
}

const listActivePackages = `-- name: ListActivePackages :many
SELECT p.id, p.created_at, p.updated_at, p.tag_id, p.package_type, p.is_active, p.quantity, p.notes, p.packaged_date_time, p.harvest_date_time, p.lab_testing_state, p.lab_testing_state_date_time, p.is_trade_sample, p.is_testing_sample, p.product_requires_remediation, p.contains_remediated_product, p.remediation_date_time, p.received_date_time, p.received_from_manifest_number, p.received_from_facility_license_number, p.received_from_facility_name, p.is_on_hold, p.archived_date, p.finished_date, p.item_id, p.provisional_label, p.is_provisional, p.is_sold, p.ppu_default, p.ppu_on_order, p.total_package_price_on_order, p.ppu_sold_price, p.total_sold_price, p.packaging_supplies_consumed, p.is_line_item, p.order_id, p.uom_id,
       pt.tag_number,
       u.name         AS uom_name,
       u.abbreviation AS uom_abbreviation,
       i.description,
       it.product_form,
       it.product_modifier,
       s.name         as strain_name,
       s.type         as strain_type,
       lt.id, lt.created_at, lt.updated_at, lt.test_name, lt.batch_code, lt.test_id_code, lt.lab_facility_name, lt.test_performed_date_time, lt.test_completed, lt.overall_passed, lt.test_type_name, lt.test_passed, lt.test_comment, lt.thc_total_percent, lt.thc_total_value, lt.cbd_percent, lt.cbd_value, lt.terpene_total_percent, lt.terpene_total_value, lt.thc_a_percent, lt.thc_a_value, lt.delta9_thc_percent, lt.delta9_thc_value, lt.delta8_thc_percent, lt.delta8_thc_value, lt.thc_v_percent, lt.thc_v_value, lt.cbd_a_percent, lt.cbd_a_value, lt.cbn_percent, lt.cbn_value, lt.cbg_a_percent, lt.cbg_a_value, lt.cbg_percent, lt.cbg_value, lt.cbc_percent, lt.cbc_value, lt.total_cannabinoid_percent, lt.total_cannabinoid_value
FROM packages p
         INNER JOIN package_tags pt ON p.tag_id = pt.id
         INNER JOIN uoms u ON p.uom_id = u.id
         INNER JOIN items i ON p.item_id = i.id
         INNER JOIN item_types it on it.id = i.item_type_id
         INNER JOIN strains s on i.strain_id = s.id
         FULL OUTER JOIN lab_tests_packages ltp on p.id = ltp.package_id
         LEFT JOIN lab_tests lt on lt.id = ltp.lab_test_id
WHERE p.is_active = TRUE
ORDER BY p.created_at DESC
`

type ListActivePackagesRow struct {
	ID                                int64           `json:"id"`
	CreatedAt                         time.Time       `json:"created_at"`
	UpdatedAt                         time.Time       `json:"updated_at"`
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
	UomID                             nulls.Int64     `json:"uom_id"`
	TagNumber                         string          `json:"tag_number"`
	UomName                           string          `json:"uom_name"`
	UomAbbreviation                   string          `json:"uom_abbreviation"`
	Description                       string          `json:"description"`
	ProductForm                       string          `json:"product_form"`
	ProductModifier                   string          `json:"product_modifier"`
	StrainName                        string          `json:"strain_name"`
	StrainType                        string          `json:"strain_type"`
	ID_2                              sql.NullInt64   `json:"id_2"`
	CreatedAt_2                       nulls.Time      `json:"created_at_2"`
	UpdatedAt_2                       nulls.Time      `json:"updated_at_2"`
	TestName                          nulls.String    `json:"test_name"`
	BatchCode                         nulls.String    `json:"batch_code"`
	TestIDCode                        nulls.String    `json:"test_id_code"`
	LabFacilityName                   nulls.String    `json:"lab_facility_name"`
	TestPerformedDateTime             nulls.Time      `json:"test_performed_date_time"`
	TestCompleted                     nulls.Bool      `json:"test_completed"`
	OverallPassed                     nulls.Bool      `json:"overall_passed"`
	TestTypeName                      nulls.String    `json:"test_type_name"`
	TestPassed                        nulls.Bool      `json:"test_passed"`
	TestComment                       nulls.String    `json:"test_comment"`
	ThcTotalPercent                   decimal.Decimal `json:"thc_total_percent"`
	ThcTotalValue                     decimal.Decimal `json:"thc_total_value"`
	CbdPercent                        decimal.Decimal `json:"cbd_percent"`
	CbdValue                          decimal.Decimal `json:"cbd_value"`
	TerpeneTotalPercent               decimal.Decimal `json:"terpene_total_percent"`
	TerpeneTotalValue                 decimal.Decimal `json:"terpene_total_value"`
	ThcAPercent                       decimal.Decimal `json:"thc_a_percent"`
	ThcAValue                         decimal.Decimal `json:"thc_a_value"`
	Delta9ThcPercent                  decimal.Decimal `json:"delta9_thc_percent"`
	Delta9ThcValue                    decimal.Decimal `json:"delta9_thc_value"`
	Delta8ThcPercent                  decimal.Decimal `json:"delta8_thc_percent"`
	Delta8ThcValue                    decimal.Decimal `json:"delta8_thc_value"`
	ThcVPercent                       decimal.Decimal `json:"thc_v_percent"`
	ThcVValue                         decimal.Decimal `json:"thc_v_value"`
	CbdAPercent                       decimal.Decimal `json:"cbd_a_percent"`
	CbdAValue                         decimal.Decimal `json:"cbd_a_value"`
	CbnPercent                        decimal.Decimal `json:"cbn_percent"`
	CbnValue                          decimal.Decimal `json:"cbn_value"`
	CbgAPercent                       decimal.Decimal `json:"cbg_a_percent"`
	CbgAValue                         decimal.Decimal `json:"cbg_a_value"`
	CbgPercent                        decimal.Decimal `json:"cbg_percent"`
	CbgValue                          decimal.Decimal `json:"cbg_value"`
	CbcPercent                        decimal.Decimal `json:"cbc_percent"`
	CbcValue                          decimal.Decimal `json:"cbc_value"`
	TotalCannabinoidPercent           decimal.Decimal `json:"total_cannabinoid_percent"`
	TotalCannabinoidValue             decimal.Decimal `json:"total_cannabinoid_value"`
}

// description: List all ACTIVE packages with related tag_number, uom, item, lab test, and source package
func (q *Queries) ListActivePackages(ctx context.Context) ([]ListActivePackagesRow, error) {
	rows, err := q.db.QueryContext(ctx, listActivePackages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListActivePackagesRow{}
	for rows.Next() {
		var i ListActivePackagesRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TagID,
			&i.PackageType,
			&i.IsActive,
			&i.Quantity,
			&i.Notes,
			&i.PackagedDateTime,
			&i.HarvestDateTime,
			&i.LabTestingState,
			&i.LabTestingStateDateTime,
			&i.IsTradeSample,
			&i.IsTestingSample,
			&i.ProductRequiresRemediation,
			&i.ContainsRemediatedProduct,
			&i.RemediationDateTime,
			&i.ReceivedDateTime,
			&i.ReceivedFromManifestNumber,
			&i.ReceivedFromFacilityLicenseNumber,
			&i.ReceivedFromFacilityName,
			&i.IsOnHold,
			&i.ArchivedDate,
			&i.FinishedDate,
			&i.ItemID,
			&i.ProvisionalLabel,
			&i.IsProvisional,
			&i.IsSold,
			&i.PpuDefault,
			&i.PpuOnOrder,
			&i.TotalPackagePriceOnOrder,
			&i.PpuSoldPrice,
			&i.TotalSoldPrice,
			&i.PackagingSuppliesConsumed,
			&i.IsLineItem,
			&i.OrderID,
			&i.UomID,
			&i.TagNumber,
			&i.UomName,
			&i.UomAbbreviation,
			&i.Description,
			&i.ProductForm,
			&i.ProductModifier,
			&i.StrainName,
			&i.StrainType,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
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

const listPackages = `-- name: ListPackages :many
SELECT p.id, p.created_at, p.updated_at, p.tag_id, p.package_type, p.is_active, p.quantity, p.notes, p.packaged_date_time, p.harvest_date_time, p.lab_testing_state, p.lab_testing_state_date_time, p.is_trade_sample, p.is_testing_sample, p.product_requires_remediation, p.contains_remediated_product, p.remediation_date_time, p.received_date_time, p.received_from_manifest_number, p.received_from_facility_license_number, p.received_from_facility_name, p.is_on_hold, p.archived_date, p.finished_date, p.item_id, p.provisional_label, p.is_provisional, p.is_sold, p.ppu_default, p.ppu_on_order, p.total_package_price_on_order, p.ppu_sold_price, p.total_sold_price, p.packaging_supplies_consumed, p.is_line_item, p.order_id, p.uom_id,
       pt.tag_number,
       u.name         AS uom_name,
       u.abbreviation AS uom_abbreviation,
       i.description,
       it.product_form,
       it.product_modifier,
       s.name         as strain_name,
       s.type         as strain_type,
       lt.id, lt.created_at, lt.updated_at, lt.test_name, lt.batch_code, lt.test_id_code, lt.lab_facility_name, lt.test_performed_date_time, lt.test_completed, lt.overall_passed, lt.test_type_name, lt.test_passed, lt.test_comment, lt.thc_total_percent, lt.thc_total_value, lt.cbd_percent, lt.cbd_value, lt.terpene_total_percent, lt.terpene_total_value, lt.thc_a_percent, lt.thc_a_value, lt.delta9_thc_percent, lt.delta9_thc_value, lt.delta8_thc_percent, lt.delta8_thc_value, lt.thc_v_percent, lt.thc_v_value, lt.cbd_a_percent, lt.cbd_a_value, lt.cbn_percent, lt.cbn_value, lt.cbg_a_percent, lt.cbg_a_value, lt.cbg_percent, lt.cbg_value, lt.cbc_percent, lt.cbc_value, lt.total_cannabinoid_percent, lt.total_cannabinoid_value
FROM packages p
         INNER JOIN package_tags pt ON p.tag_id = pt.id
         INNER JOIN uoms u ON p.uom_id = u.id
         INNER JOIN items i ON p.item_id = i.id
         INNER JOIN item_types it on it.id = i.item_type_id
         INNER JOIN strains s on i.strain_id = s.id
         FULL OUTER JOIN lab_tests_packages ltp on p.id = ltp.package_id
         LEFT JOIN lab_tests lt on lt.id = ltp.lab_test_id
`

type ListPackagesRow struct {
	ID                                int64           `json:"id"`
	CreatedAt                         time.Time       `json:"created_at"`
	UpdatedAt                         time.Time       `json:"updated_at"`
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
	UomID                             nulls.Int64     `json:"uom_id"`
	TagNumber                         string          `json:"tag_number"`
	UomName                           string          `json:"uom_name"`
	UomAbbreviation                   string          `json:"uom_abbreviation"`
	Description                       string          `json:"description"`
	ProductForm                       string          `json:"product_form"`
	ProductModifier                   string          `json:"product_modifier"`
	StrainName                        string          `json:"strain_name"`
	StrainType                        string          `json:"strain_type"`
	ID_2                              sql.NullInt64   `json:"id_2"`
	CreatedAt_2                       nulls.Time      `json:"created_at_2"`
	UpdatedAt_2                       nulls.Time      `json:"updated_at_2"`
	TestName                          nulls.String    `json:"test_name"`
	BatchCode                         nulls.String    `json:"batch_code"`
	TestIDCode                        nulls.String    `json:"test_id_code"`
	LabFacilityName                   nulls.String    `json:"lab_facility_name"`
	TestPerformedDateTime             nulls.Time      `json:"test_performed_date_time"`
	TestCompleted                     nulls.Bool      `json:"test_completed"`
	OverallPassed                     nulls.Bool      `json:"overall_passed"`
	TestTypeName                      nulls.String    `json:"test_type_name"`
	TestPassed                        nulls.Bool      `json:"test_passed"`
	TestComment                       nulls.String    `json:"test_comment"`
	ThcTotalPercent                   decimal.Decimal `json:"thc_total_percent"`
	ThcTotalValue                     decimal.Decimal `json:"thc_total_value"`
	CbdPercent                        decimal.Decimal `json:"cbd_percent"`
	CbdValue                          decimal.Decimal `json:"cbd_value"`
	TerpeneTotalPercent               decimal.Decimal `json:"terpene_total_percent"`
	TerpeneTotalValue                 decimal.Decimal `json:"terpene_total_value"`
	ThcAPercent                       decimal.Decimal `json:"thc_a_percent"`
	ThcAValue                         decimal.Decimal `json:"thc_a_value"`
	Delta9ThcPercent                  decimal.Decimal `json:"delta9_thc_percent"`
	Delta9ThcValue                    decimal.Decimal `json:"delta9_thc_value"`
	Delta8ThcPercent                  decimal.Decimal `json:"delta8_thc_percent"`
	Delta8ThcValue                    decimal.Decimal `json:"delta8_thc_value"`
	ThcVPercent                       decimal.Decimal `json:"thc_v_percent"`
	ThcVValue                         decimal.Decimal `json:"thc_v_value"`
	CbdAPercent                       decimal.Decimal `json:"cbd_a_percent"`
	CbdAValue                         decimal.Decimal `json:"cbd_a_value"`
	CbnPercent                        decimal.Decimal `json:"cbn_percent"`
	CbnValue                          decimal.Decimal `json:"cbn_value"`
	CbgAPercent                       decimal.Decimal `json:"cbg_a_percent"`
	CbgAValue                         decimal.Decimal `json:"cbg_a_value"`
	CbgPercent                        decimal.Decimal `json:"cbg_percent"`
	CbgValue                          decimal.Decimal `json:"cbg_value"`
	CbcPercent                        decimal.Decimal `json:"cbc_percent"`
	CbcValue                          decimal.Decimal `json:"cbc_value"`
	TotalCannabinoidPercent           decimal.Decimal `json:"total_cannabinoid_percent"`
	TotalCannabinoidValue             decimal.Decimal `json:"total_cannabinoid_value"`
}

// description: List all packages with related tag_number, uom, item, lab test, and source package
func (q *Queries) ListPackages(ctx context.Context) ([]ListPackagesRow, error) {
	rows, err := q.db.QueryContext(ctx, listPackages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPackagesRow{}
	for rows.Next() {
		var i ListPackagesRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TagID,
			&i.PackageType,
			&i.IsActive,
			&i.Quantity,
			&i.Notes,
			&i.PackagedDateTime,
			&i.HarvestDateTime,
			&i.LabTestingState,
			&i.LabTestingStateDateTime,
			&i.IsTradeSample,
			&i.IsTestingSample,
			&i.ProductRequiresRemediation,
			&i.ContainsRemediatedProduct,
			&i.RemediationDateTime,
			&i.ReceivedDateTime,
			&i.ReceivedFromManifestNumber,
			&i.ReceivedFromFacilityLicenseNumber,
			&i.ReceivedFromFacilityName,
			&i.IsOnHold,
			&i.ArchivedDate,
			&i.FinishedDate,
			&i.ItemID,
			&i.ProvisionalLabel,
			&i.IsProvisional,
			&i.IsSold,
			&i.PpuDefault,
			&i.PpuOnOrder,
			&i.TotalPackagePriceOnOrder,
			&i.PpuSoldPrice,
			&i.TotalSoldPrice,
			&i.PackagingSuppliesConsumed,
			&i.IsLineItem,
			&i.OrderID,
			&i.UomID,
			&i.TagNumber,
			&i.UomName,
			&i.UomAbbreviation,
			&i.Description,
			&i.ProductForm,
			&i.ProductModifier,
			&i.StrainName,
			&i.StrainType,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
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

const updatePackage = `-- name: UpdatePackage :one
UPDATE packages
SET tag_id                                = $1,
    package_type                          = $2,
    quantity                              = $3,
    notes                                 = $4,
    packaged_date_time                    = $5,
    harvest_date_time                     = $6,
    lab_testing_state                     = $7,
    lab_testing_state_date_time           = $8,
    is_trade_sample                       = $9,
    is_testing_sample                     = $10,
    product_requires_remediation          = $11,
    contains_remediated_product           = $12,
    remediation_date_time                 = $13,
    received_date_time                    = $14,
    received_from_manifest_number         = $15,
    received_from_facility_license_number = $16,
    received_from_facility_name           = $17,
    is_on_hold                            = $18,
    archived_date                         = $19,
    finished_date                         = $20,
    item_id                               = $21,
    provisional_label                     = $22,
    is_provisional                        = $23,
    is_sold                               = $24,
    ppu_default                           = $25,
    ppu_on_order                          = $26,
    total_package_price_on_order          = $27,
    ppu_sold_price                        = $28,
    total_sold_price                      = $29,
    packaging_supplies_consumed           = $30,
    is_line_item                          = $31,
    order_id                              = $32,
    uom_id                                = $33,
    is_active                             = $34
WHERE id = $35
RETURNING id, created_at, updated_at, tag_id, package_type, is_active, quantity, notes, packaged_date_time, harvest_date_time, lab_testing_state, lab_testing_state_date_time, is_trade_sample, is_testing_sample, product_requires_remediation, contains_remediated_product, remediation_date_time, received_date_time, received_from_manifest_number, received_from_facility_license_number, received_from_facility_name, is_on_hold, archived_date, finished_date, item_id, provisional_label, is_provisional, is_sold, ppu_default, ppu_on_order, total_package_price_on_order, ppu_sold_price, total_sold_price, packaging_supplies_consumed, is_line_item, order_id, uom_id
`

type UpdatePackageParams struct {
	TagID                             nulls.Int64     `json:"tag_id"`
	PackageType                       string          `json:"package_type"`
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
	UomID                             nulls.Int64     `json:"uom_id"`
	IsActive                          bool            `json:"is_active"`
	ID                                int64           `json:"id"`
}

// description: Update a package
func (q *Queries) UpdatePackage(ctx context.Context, arg UpdatePackageParams) (Package, error) {
	row := q.db.QueryRowContext(ctx, updatePackage,
		arg.TagID,
		arg.PackageType,
		arg.Quantity,
		arg.Notes,
		arg.PackagedDateTime,
		arg.HarvestDateTime,
		arg.LabTestingState,
		arg.LabTestingStateDateTime,
		arg.IsTradeSample,
		arg.IsTestingSample,
		arg.ProductRequiresRemediation,
		arg.ContainsRemediatedProduct,
		arg.RemediationDateTime,
		arg.ReceivedDateTime,
		arg.ReceivedFromManifestNumber,
		arg.ReceivedFromFacilityLicenseNumber,
		arg.ReceivedFromFacilityName,
		arg.IsOnHold,
		arg.ArchivedDate,
		arg.FinishedDate,
		arg.ItemID,
		arg.ProvisionalLabel,
		arg.IsProvisional,
		arg.IsSold,
		arg.PpuDefault,
		arg.PpuOnOrder,
		arg.TotalPackagePriceOnOrder,
		arg.PpuSoldPrice,
		arg.TotalSoldPrice,
		arg.PackagingSuppliesConsumed,
		arg.IsLineItem,
		arg.OrderID,
		arg.UomID,
		arg.IsActive,
		arg.ID,
	)
	var i Package
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagID,
		&i.PackageType,
		&i.IsActive,
		&i.Quantity,
		&i.Notes,
		&i.PackagedDateTime,
		&i.HarvestDateTime,
		&i.LabTestingState,
		&i.LabTestingStateDateTime,
		&i.IsTradeSample,
		&i.IsTestingSample,
		&i.ProductRequiresRemediation,
		&i.ContainsRemediatedProduct,
		&i.RemediationDateTime,
		&i.ReceivedDateTime,
		&i.ReceivedFromManifestNumber,
		&i.ReceivedFromFacilityLicenseNumber,
		&i.ReceivedFromFacilityName,
		&i.IsOnHold,
		&i.ArchivedDate,
		&i.FinishedDate,
		&i.ItemID,
		&i.ProvisionalLabel,
		&i.IsProvisional,
		&i.IsSold,
		&i.PpuDefault,
		&i.PpuOnOrder,
		&i.TotalPackagePriceOnOrder,
		&i.PpuSoldPrice,
		&i.TotalSoldPrice,
		&i.PackagingSuppliesConsumed,
		&i.IsLineItem,
		&i.OrderID,
		&i.UomID,
	)
	return i, err
}
