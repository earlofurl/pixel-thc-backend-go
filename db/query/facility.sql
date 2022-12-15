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