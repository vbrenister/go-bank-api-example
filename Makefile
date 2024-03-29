postgres:
	@docker run --name=gobank -d --env POSTGRES_DB=gobank --env POSTGRES_PASSWORD=gobank --env POSTGRES_USER=gobank -p 5432:5432 postgres:11.12

dropposgres:
	@docker stop gobank
	@docker rm gobank

migrateup:
	@migrate -path db/migration -database "postgresql://gobank:gobank@localhost:5432/gobank?sslmode=disable" --verbose up

migratedown:
	@migrate -path db/migration -database "postgresql://gobank:gobank@localhost:5432/gobank?sslmode=disable" --verbose down

sqlc:
	@sqlc generate

test:
	@go test -v -cover ./...

fmt:
	@go fmt ./...

server:
	@go run main.go

mock:
	@mockgen -package mockdb -build_flags=--mod=mod -destination db/mock/store.go github.com/vbrenister/go-bank-api-example/db/sqlc Store

proto:
	@rm -f pb/*.go
	@protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

evans: 
	@evans --host localhost --port 4000 -r repl

.PHONY: postgres dropposgres migrateup migratedown sqlc fmt server mock proto evans