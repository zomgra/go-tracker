build:
	go build -o bin/bitbucket.exe cmd/bitbucket/main.go
run:
	./bin/bitbucket.exe
db:
	docker run -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=postgresDB -d --name postgres-bitbucket postgres