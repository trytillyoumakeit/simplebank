postgres: 
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=td -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
createdb: 
	 docker exec -it postgres16 createdb -U td --owner=td simplebank

migrateup : 
	migrate -path internal/migration -database "postgresql://td:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown :
	migrate -path internal/migration -database "postgresql://td:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

dropdb: 
	 docker exec -it postgres16 dropdb -U td simplebank

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server: 
	go run main.go

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server
