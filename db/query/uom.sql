-- name: CreateUom :one
-- description: Create a new UOM
INSERT INTO uoms (name, abbreviation)
VALUES ($1, $2)
RETURNING *;

-- name: GetUom :one
-- description: Get a UOM by ID
SELECT *
FROM uoms
WHERE id = $1
LIMIT 1;

-- name: ListUoms :many
-- description: List all UOMs
SELECT *
FROM uoms
ORDER BY name;

-- name: GetUomByName :one
-- description: Get a UOM by name
SELECT *
FROM uoms
WHERE name = $1
LIMIT 1;

-- name: UpdateUom :one
-- description: Update a UOM by ID
UPDATE uoms
SET name         = $2,
    abbreviation = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUom :exec
-- description: Delete a UOM by ID
DELETE
FROM uoms
WHERE id = $1;