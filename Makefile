create_migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir database/migrations -seq $$name

sqlc:
	sqlc generate --file sqlc.yaml

evans:
	evans --host localhost --port 8081 -r repl

.PHONY: create_migration sqlc evans
