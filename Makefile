MIGRATE_DIR = ./sql/migrations

.PHONY: migrate-up

migrate-up:
	migrate -source file://$(MIGRATE_DIR) -database postgres://localhost:5432/taskdb?sslmode=disable up 2

