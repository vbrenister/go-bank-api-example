postgres:
	@docker run --name=gobank -d --env POSTGRES_DB=gobank --env POSTGRES_PASSWORD=gobank --env POSTGRES_USER=gobank -p 5432:5432 postgres:11.12

dropposgres:
	@docker stop gobank
	@docker rm gobank

migrateup:
	@migrate -path migration -database "postgresql://gobank:gobank@localhost:5432/gobank?sslmode=disable" --verbose up

migratedown:
	@migrate -path migration -database "postgresql://gobank:gobank@localhost:5432/gobank?sslmode=disable" --verbose down

sqlc:
	@sqlc generate

test:
	@go test -v -cover ./...

.PHONY: postgres dropposgres migrateup migratedown sqlc