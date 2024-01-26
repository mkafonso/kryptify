create_migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir database/migrations -seq $$name

sqlc:
	sqlc generate --file sqlc.yaml

.PHONY: create_migration sqlc
