MIGRATE_DIR = ./sql/migrations
DATABASE_URL = postgres://localhost:5432/taskdb

.PHONY: migrate-up, migrate-down, gen-sqlc, lint-sql

migrate-up:
	migrate -source file://$(MIGRATE_DIR) -database $(DATABASE_URL)?sslmode=disable up

migrate-down:
	migrate -source file://$(MIGRATE_DIR) -database $(DATABASE_URL)?sslmode=disable down 

gen-sqlc:
	sqlc generate

lint-sql:
	sqlc vet
