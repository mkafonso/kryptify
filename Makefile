create_migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir database/migrations -seq $$name

sqlc:
	sqlc generate --file sqlc.yaml

evans:
	evans --host localhost --port 8081 -r repl

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: create_migration sqlc evans proto
