-- name: CreatePackage :one
-- description: Create a package
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
                      uom_id,
                      facility_location_id)
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
        $34,
        $35)
RETURNING *;

-- name: GetPackage :one
-- description: Get a package
SELECT *
FROM packages
WHERE id = $1
LIMIT 1;

-- name: GetPackageByTagID :one
-- description: Get a package by tag id
SELECT *
FROM packages
WHERE tag_id = $1
LIMIT 1;

-- name: ListActivePackages :many
-- description: List all ACTIVE packages with related tag_number, uom, item, lab test, and source package
SELECT p.*,
       pt.tag_number,
       u.name         AS uom_name,
       u.abbreviation AS uom_abbreviation,
       i.description,
       it.product_form,
       it.product_modifier,
       s.name         as strain_name,
       s.type         as strain_type,
       lt.id          as lab_test_id,
       lt.*
FROM packages p
         INNER JOIN package_tags pt ON p.tag_id = pt.id
         INNER JOIN uoms u ON p.uom_id = u.id
         INNER JOIN items i ON p.item_id = i.id
         INNER JOIN item_types it on it.id = i.item_type_id
         INNER JOIN strains s on i.strain_id = s.id
         FULL OUTER JOIN lab_tests_packages ltp on p.id = ltp.package_id
         LEFT JOIN lab_tests lt on lt.id = ltp.lab_test_id
WHERE p.is_active = TRUE
ORDER BY p.created_at DESC;

-- name: ListPackages :many
-- description: List all packages with related tag_number, uom, item, lab test, and source package
SELECT p.*,
       pt.tag_number,
       u.name         AS uom_name,
       u.abbreviation AS uom_abbreviation,
       i.description,
       it.product_form,
       it.product_modifier,
       s.name         as strain_name,
       s.type         as strain_type,
       lt.id          as lab_test_id,
       lt.*
FROM packages p
         INNER JOIN package_tags pt ON p.tag_id = pt.id
         INNER JOIN uoms u ON p.uom_id = u.id
         INNER JOIN items i ON p.item_id = i.id
         INNER JOIN item_types it on it.id = i.item_type_id
         INNER JOIN strains s on i.strain_id = s.id
         FULL OUTER JOIN lab_tests_packages ltp on p.id = ltp.package_id
         LEFT JOIN lab_tests lt on lt.id = ltp.lab_test_id;

-- name: UpdatePackage :one
-- description: Update a package
UPDATE packages
SET tag_id                                = COALESCE(sqlc.narg(tag_id), tag_id),
    package_type                          = COALESCE(sqlc.narg(package_type), package_type),
    quantity                              = COALESCE(sqlc.narg(quantity), quantity),
    notes                                 = COALESCE(sqlc.narg(notes), notes),
    packaged_date_time                    = COALESCE(sqlc.narg(packaged_date_time), packaged_date_time),
    harvest_date_time                     = COALESCE(sqlc.narg(harvest_date_time), harvest_date_time),
    lab_testing_state                     = COALESCE(sqlc.narg(lab_testing_state), lab_testing_state),
    lab_testing_state_date_time           = COALESCE(sqlc.narg(lab_testing_state_date_time),
                                                     lab_testing_state_date_time),
    is_trade_sample                       = COALESCE(sqlc.narg(is_trade_sample), is_trade_sample),
    is_testing_sample                     = COALESCE(sqlc.narg(is_testing_sample), is_testing_sample),
    product_requires_remediation          = COALESCE(sqlc.narg(product_requires_remediation),
                                                     product_requires_remediation),
    contains_remediated_product           = COALESCE(sqlc.narg(contains_remediated_product),
                                                     contains_remediated_product),
    remediation_date_time                 = COALESCE(sqlc.narg(remediation_date_time), remediation_date_time),
    received_date_time                    = COALESCE(sqlc.narg(received_date_time), received_date_time),
    received_from_manifest_number         = COALESCE(sqlc.narg(received_from_manifest_number),
                                                     received_from_manifest_number),
    received_from_facility_license_number = COALESCE(sqlc.narg(received_from_facility_license_number),
                                                     received_from_facility_license_number),
    received_from_facility_name           = COALESCE(sqlc.narg(received_from_facility_name),
                                                     received_from_facility_name),
    is_on_hold                            = COALESCE(sqlc.narg(is_on_hold), is_on_hold),
    archived_date                         = COALESCE(sqlc.narg(archived_date), archived_date),
    finished_date                         = COALESCE(sqlc.narg(finished_date), finished_date),
    item_id                               = COALESCE(sqlc.narg(item_id), item_id),
    provisional_label                     = COALESCE(sqlc.narg(provisional_label), provisional_label),
    is_provisional                        = COALESCE(sqlc.narg(is_provisional), is_provisional),
    is_sold                               = COALESCE(sqlc.narg(is_sold), is_sold),
    ppu_default                           = COALESCE(sqlc.narg(ppu_default), ppu_default),
    ppu_on_order                          = COALESCE(sqlc.narg(ppu_on_order), ppu_on_order),
    total_package_price_on_order          = COALESCE(sqlc.narg(total_package_price_on_order),
                                                     total_package_price_on_order),
    ppu_sold_price                        = COALESCE(sqlc.narg(ppu_sold_price), ppu_sold_price),
    total_sold_price                      = COALESCE(sqlc.narg(total_sold_price), total_sold_price),
    packaging_supplies_consumed           = COALESCE(sqlc.narg(packaging_supplies_consumed),
                                                     packaging_supplies_consumed),
    is_line_item                          = COALESCE(sqlc.narg(is_line_item), is_line_item),
    order_id                              = COALESCE(sqlc.narg(order_id), order_id),
    uom_id                                = COALESCE(sqlc.narg(uom_id), uom_id),
    is_active                             = COALESCE(sqlc.narg(is_active), is_active)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: AddPackageQuantity :one
-- description: Add quantity to a package (can be negative to subtract)
-- arguments: package_id int, quantity float
UPDATE packages
SET quantity = quantity + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SubtractPackageQuantity :one
-- description: Subtract quantity from a package (can be negative to add)
-- arguments: package_id int, quantity float
UPDATE packages
SET quantity = quantity - sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeletePackage :exec
-- description: Delete a package
DELETE
FROM packages
WHERE id = $1;

-- name: AssignSourcePackageChildPackage :one
-- description: Assign a source package child package relationship on junction table
INSERT INTO source_packages_child_packages (source_package_id, child_package_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetLabTestByPackageID :one
-- description: Get a lab test connected to package by package id in lab_tests_packages junction table
SELECT *
FROM lab_tests_packages
WHERE package_id = $1
LIMIT 1;
