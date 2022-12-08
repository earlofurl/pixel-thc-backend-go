-- name: CreateLabTest :one
-- description: Create a lab test
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
RETURNING *;

-- name: GetLabTest :one
-- description: Get a lab test by ID
SELECT *
FROM lab_tests
WHERE id = $1
LIMIT 1;

-- name: ListLabTests :many
-- description: List all lab tests
SELECT *
FROM lab_tests
ORDER BY created_at;

-- name: UpdateLabTest :one
-- description: Update a lab test
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
RETURNING *;

-- name: DeleteLabTest :exec
-- description: Delete a lab test by ID
DELETE
FROM lab_tests
WHERE id = $1;

-- name: AssignLabTestToPackage :exec
-- description: Assign a lab test to a package via junction table
INSERT INTO lab_tests_packages (lab_test_id, package_id) VALUES ($1, $2);