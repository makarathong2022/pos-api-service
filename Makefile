
postgres:
	docker run --name postgres14.5 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Thong@123 -d postgres:14.5-alpine
createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root outlet_0001
dropdb:
	docker exec -it postgres14.5 dropdb outlet_0001

migrateup:
	migrate -path db/migration -database "postgresql://root:Thong@123@localhost:5432/outlet_0001?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:Thong@123@localhost:5432/outlet_0001?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: createdb dropdb postgres migrateup migratedown sqlc server