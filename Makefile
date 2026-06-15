MIGRATE_DIR = ./sql/migrations
DATABASE_URL = postgres://localhost:5432/taskdb

.PHONY: migrate-up, migrate-down

migrate-up:
	migrate -source file://$(MIGRATE_DIR) -database $(DATABASE_URL)?sslmode=disable up

migrate-down:
	migrate -source file://$(MIGRATE_DIR) -database $(DATABASE_URL)?sslmode=disable down 

