// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: facility.sql

package db

import (
	"context"
)

const createFacility = `-- name: CreateFacility :one
INSERT INTO facilities (name, license_number)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, name, license_number
`

type CreateFacilityParams struct {
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
}

// description: Create a new facility
func (q *Queries) CreateFacility(ctx context.Context, arg CreateFacilityParams) (Facility, error) {
	row := q.db.QueryRowContext(ctx, createFacility, arg.Name, arg.LicenseNumber)
	var i Facility
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.LicenseNumber,
	)
	return i, err
}

const getFacility = `-- name: GetFacility :one
SELECT id, created_at, updated_at, name, license_number
FROM facilities
WHERE id = $1
`

// description: Get a facility by ID
func (q *Queries) GetFacility(ctx context.Context, id int64) (Facility, error) {
	row := q.db.QueryRowContext(ctx, getFacility, id)
	var i Facility
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.LicenseNumber,
	)
	return i, err
}

const listFacilities = `-- name: ListFacilities :many
SELECT id, created_at, updated_at, name, license_number
FROM facilities
`

// description: List all facilities
func (q *Queries) ListFacilities(ctx context.Context) ([]Facility, error) {
	rows, err := q.db.QueryContext(ctx, listFacilities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Facility{}
	for rows.Next() {
		var i Facility
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.LicenseNumber,
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