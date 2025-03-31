postgres:
	docker run --name postgres-container -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine	
createdb:
	docker exec -it postgres-container createdb --username=root --owner=root alerts
dropdb: 
	docker exec -it postgres-container dropdb alerts
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/alerts?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/alerts?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mockdb:
	mockgen -destination db\mock\store.go -package mockdb github.com.jatin711-debug/simplebank/db/sqlc Store
.PHONY: dropdb createdb postgres migrateup migratedown sqlc test server mockdb