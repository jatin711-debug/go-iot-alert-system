key ?= key
value ?= value
sname ?= secret-name 
nspace ?= default

create-secret:
	kubectl create secret generic $(sname) --from-literal=$(key)=$(value) -n $(nspace)

postgres:
	docker run --name postgres-1 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:17	

createdb:
	docker exec -it docker-postgres-1 createdb --username=root --owner=root root

dropdb: 
	docker exec -it docker-postgres-1 dropdb alerts

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/alerts?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/alerts?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v ${PWD}:/src -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mockdb:
	mockgen -destination db/mock/store.go -package mockdb github.com.jatin711-debug/simplebank/db/sqlc Store

delete-pods:
	kubectl delete all --all -n default
	
.PHONY: create-secret postgres createdb dropdb migrateup migratedown sqlc test server mockdb delete-pods
