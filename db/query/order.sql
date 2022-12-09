-- name: CreateOrder :one
-- description: Create a new order
INSERT INTO orders (scheduled_pack_date_time, scheduled_ship_date_time, scheduled_delivery_date_time,
                    actual_pack_date_time, actual_ship_date_time, actual_delivery_date_time, order_total, notes, status,
                    customer_name)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetOrder :one
-- description: Get a single order by id
SELECT *
FROM orders
WHERE id = $1
LIMIT 1;

-- name: ListOrders :many
-- description: List all orders
SELECT *
FROM orders
ORDER BY created_at DESC;

-- name: UpdateOrder :one
-- description: Update a single order by id
UPDATE orders
SET scheduled_pack_date_time     = $1,
    scheduled_ship_date_time     = $2,
    scheduled_delivery_date_time = $3,
    actual_pack_date_time        = $4,
    actual_ship_date_time        = $5,
    actual_delivery_date_time    = $6,
    order_total                  = $7,
    notes                        = $8,
    status                       = $9,
    customer_name                = $10
WHERE id = $11
RETURNING *;

-- name: DeleteOrder :exec
-- description: Delete a single order by id
DELETE
FROM orders
WHERE id = $1;