ifneq (,$(wildcard ./.env))
	include .env
	export
endif

MIGRATIONS_DIR = ./migrations
DATABASE_URL = postgres://postgres:mypassword@localhost:5432/avitoSegmentsDb?sslmode=disable
UP_STEP =
DOWN_STEP = -all
MAIN_GO = cmd/banner-service/main.go

tidy:
	go mod tidy

migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $(NAME)

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) up $(UP_STEP)

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) down $(DOWN_STEP)