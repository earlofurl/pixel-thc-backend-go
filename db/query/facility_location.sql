-- name: CreateFacilityLocation :one
-- description: Create a new location within a facility
INSERT INTO facility_locations (name, facility_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetFacilityLocation :one
-- description: Get a location within a facility by ID
SELECT *
FROM facility_locations
WHERE id = $1;

-- name: ListFacilityLocations :many
-- description: List all locations within facilities
SELECT *
FROM facilities;

-- name: ListFacilityLocationsByFacility :many
-- description: List all locations within a facility
SELECT *
FROM facility_locations
WHERE facility_id = $1;