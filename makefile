postgres:
	docker run --name postgres-latest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres

createdb:
	docker exec -it postgres-latest createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-latest simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:1234@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc