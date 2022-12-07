-- name: CreateItem :one
-- description: Create a new item
INSERT INTO items (description, is_used, item_type_id, strain_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetItem :one
-- description: Get an item by ID
SELECT * FROM items
WHERE id = $1
LIMIT 1;

-- name: ListItems :many
-- description: List all items
SELECT * FROM items ORDER BY strain_id;

-- name: UpdateItem :one
-- description: Update an item by ID
UPDATE items SET
  description = $2,
  is_used = $3,
  item_type_id = $4,
  strain_id = $5
WHERE id = $1
RETURNING *;

-- name: DeleteItem :exec
-- description: Delete an item by ID
DELETE FROM items
WHERE id = $1;