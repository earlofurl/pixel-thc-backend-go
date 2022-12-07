-- name: CreateStrain :one
-- description: Create a strain
INSERT INTO strains (name,
                     type,
                     yield_average,
                     terp_average_total,
                     terp_1,
                     terp_1_value,
                     terp_2,
                     terp_2_value,
                     terp_3,
                     terp_3_value,
                     terp_4,
                     terp_4_value,
                     terp_5,
                     terp_5_value,
                     thc_average,
                     total_cannabinoid_average,
                     light_dep_2022,
                     fall_harvest_2022,
                     quantity_available)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
RETURNING *;

-- name: GetStrain :one
-- description: Get a strain by ID
SELECT * FROM strains
WHERE id = $1
LIMIT 1;

-- name: ListStrains :many
-- description: List all strains
SELECT * FROM strains ORDER BY name;

-- name: UpdateStrain :one
-- description: Update a strain
UPDATE strains
SET name = $2,
    type = $3,
    yield_average = $4,
    terp_average_total = $5,
    terp_1 = $6,
    terp_1_value = $7,
    terp_2 = $8,
    terp_2_value = $9,
    terp_3 = $10,
    terp_3_value = $11,
    terp_4 = $12,
    terp_4_value = $13,
    terp_5 = $14,
    terp_5_value = $15,
    thc_average = $16,
    total_cannabinoid_average = $17,
    light_dep_2022 = $18,
    fall_harvest_2022 = $19,
    quantity_available = $20
WHERE id = $1
RETURNING *;

-- name: DeleteStrain :exec
-- description: Delete a strain by ID
DELETE FROM strains
WHERE id = $1;