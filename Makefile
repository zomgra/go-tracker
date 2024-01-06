.PHONY: test

build:
	go build -o bin/main main.go

run:
	go run .\cmd\tracker\main.go --config="configs/app.env"
db:
	docker run -p 5432:5432 --rm -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=postgresDB -d --name postgres-tracker postgres
test_db:
	docker exec -it postgres-tracker createdb -U postgres postgresTestDB

test:
	go test -v  -cover ./... --config="../configs/test.env"

migrate_up:
	docker run --rm -v $(CURDIR)/internal/db/postgres/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://postgres:password@localhost:5432/postgresDB?sslmode=disable up
migrate_up_test:
	docker run --rm -v $(CURDIR)/internal/db/postgres/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://postgres:password@localhost:5432/postgresTestDB?sslmode=disable up