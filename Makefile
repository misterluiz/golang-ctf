createdb:
	createdb --username=postgres --owner=postgres go-ctf

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:14-alpine

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/go-ctf?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/go-ctfsslmode=disable" -verbose down

test:
	go test -v -cover ./...

.PHONY: createdb postgres dropdb migrateup migrationdrop test 