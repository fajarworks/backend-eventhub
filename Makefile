include ./.env

DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATION_PATH=db/migrations
SEEDER_PATH=db/seeds

migrate-create:
	@migrate create -ext sql -dir $(MIGRATION_PATH) -seq create_$(NAME)_table

migrate-up:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) up

migrate-down:
	@migrate -database $(DB_URL) -path $(MIGRATION_PATH) down

seed-list:
	@for %%f in (./db/seeds/*) do @echo %%f

run-seed:
	@for %%f in ($(SEEDER_PATH)\*.sql) do @(psql $(DB_URL) < $(SEEDER_PATH)/%%f)

