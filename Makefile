DB_URL=postgresql://root:secret@localhost:5432/pixel_thc_dev?sslmode=disable

network:
	docker network create pixel-thc-network

postgres:
	docker run --name postgres --network pixel-thc-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root pixel_thc_dev

dropdb:
	docker exec -it postgres dropdb pixel_thc_dev

migrateup:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down 1

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/pixelthc-entrypoint/main.go

mock:
	mockgen -package mockdb -destination internal/db/mock/store.go pixel-thc-backend-go/internal/db/sqlc Store

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock redis
