migrateup:
	migrate -path db/migration -database "postgres://maxxue@localhost:5432/go-server-template?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://maxxue@localhost:5432/go-server-template?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
