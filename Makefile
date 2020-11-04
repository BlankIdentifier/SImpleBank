migrateup:
	migrate -path db/migration -database "postgresql://sb:${DB_PASSWORD}@localhost:5432/simplebank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://sb:${DB_PASSWORD}@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./... -count=1

.PHONY: migrateup migratedown sqlc test
