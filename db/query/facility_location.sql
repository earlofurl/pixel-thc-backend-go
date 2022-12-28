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
FROM facility_locations;

-- name: ListFacilityLocationsByFacility :many
-- description: List all locations within a facility
SELECT *
FROM facility_locations
WHERE facility_id = $1;

-- name: UpdateFacilityLocation :one
-- description: Update a location within a facility
UPDATE facility_locations
SET name        = $1,
    facility_id = $2
WHERE id = $3
RETURNING *;

-- name: DeleteFacilityLocation :exec
-- description: Delete a location within a facility
DELETE
FROM facility_locations
WHERE id = $1;
