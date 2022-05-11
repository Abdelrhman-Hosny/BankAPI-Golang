DB_URL=postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgress_instance -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it postgress_instance createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgress_instance dropdb --username=postgres simple_bank

migrate_up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate