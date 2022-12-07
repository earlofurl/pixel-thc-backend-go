-- name: CreateItemType :one
-- description: Create a new item type
INSERT INTO item_types (product_form,
                        product_modifier,
                        uom_default,
                        product_category_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetItemType :one
-- description: Get an item type by id
SELECT *
FROM item_types
WHERE id = $1
LIMIT 1;

-- name: ListItemTypes :many
-- description: List all item types
SELECT *
FROM item_types
ORDER BY id;

-- name: UpdateItemType :one
-- description: Update an item type by ID
UPDATE item_types
SET product_form        = $1,
    product_modifier    = $2,
    uom_default         = $3,
    product_category_id = $4,
    updated_at          = NOW()
WHERE id = $5
RETURNING *;

-- name: DeleteItemType :exec
-- description: Delete an item type by ID
DELETE
FROM item_types
WHERE id = $1;