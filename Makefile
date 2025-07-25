include .env
MIGRATION_PATH=./cmd/migrate/migrations

.PHONY: create-migrate
migration:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@, $(MAKECMDGOALS))

.PHONY: migration-up
migration-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up

.PHONY: migration-down
migration-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filler-out $@, $(MAKECMDGOALS))

# .PHONE: force-clean
# force-clean:
# 	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) force $(filter-out $@, $(MAKECMDGOALS))


.PHONY: force-clean
force-clean:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) force 3

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

.PHONY: gen-docs
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt
