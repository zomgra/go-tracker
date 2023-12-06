build:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-freebsd-386 main.go
run:
	go run cmd/bitbucket/main.go
execute:
	./bin/bitbucket.exe
db:
	docker run -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=postgresDB -d --name postgres-bitbucket postgres
migrate:
	migrate -source file://./pkg/db/migrations -database postgresql://postgres:password@localhost:5432/postgresDB?sslmode=disable up 