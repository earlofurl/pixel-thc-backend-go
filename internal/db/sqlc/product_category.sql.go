// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: product_category.sql

package db

import (
	"context"
)

const createProductCategory = `-- name: CreateProductCategory :one
INSERT INTO product_categories (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

// description: Create a new product category
func (q *Queries) CreateProductCategory(ctx context.Context, name string) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, createProductCategory, name)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProductCategory = `-- name: DeleteProductCategory :exec
DELETE
FROM product_categories
WHERE id = $1
`

// description: Delete a product category
func (q *Queries) DeleteProductCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductCategory, id)
	return err
}

const getProductCategory = `-- name: GetProductCategory :one
SELECT id, name, created_at, updated_at
FROM product_categories
WHERE id = $1
LIMIT 1
`

// description: Get a product category by ID
func (q *Queries) GetProductCategory(ctx context.Context, id int64) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, getProductCategory, id)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductCategoryByName = `-- name: GetProductCategoryByName :one
SELECT id, name, created_at, updated_at
FROM product_categories
WHERE name = $1
LIMIT 1
`

// description: Get a product category by name
func (q *Queries) GetProductCategoryByName(ctx context.Context, name string) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, getProductCategoryByName, name)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProductCategories = `-- name: ListProductCategories :many
SELECT id, name, created_at, updated_at
FROM product_categories
ORDER BY name
`

// description: List all product categories
func (q *Queries) ListProductCategories(ctx context.Context) ([]ProductCategory, error) {
	rows, err := q.db.QueryContext(ctx, listProductCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductCategory{}
	for rows.Next() {
		var i ProductCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateProductCategory = `-- name: UpdateProductCategory :one
UPDATE product_categories
SET name       = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, name, created_at, updated_at
`

type UpdateProductCategoryParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// description: Update a product category
func (q *Queries) UpdateProductCategory(ctx context.Context, arg UpdateProductCategoryParams) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, updateProductCategory, arg.ID, arg.Name)
	var i ProductCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}