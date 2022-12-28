-- name: CreateFacility :one
-- description: Create a new facility
INSERT INTO facilities (name, license_number)
VALUES ($1, $2)
RETURNING *;

-- name: GetFacility :one
-- description: Get a facility by ID
SELECT *
FROM facilities
WHERE id = $1;

-- name: ListFacilities :many
-- description: List all facilities
SELECT *
FROM facilities;

-- name: UpdateFacility :one
-- description: Update a facility by ID
UPDATE facilities
SET name = $2, license_number = $3
WHERE id = $1
RETURNING *;

-- name: DeleteFacility :exec
-- description: Delete a facility by ID
DELETE FROM facilities
WHERE id = $1;
