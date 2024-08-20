create_migration:
	# create a new migration
	migrate create -ext sql -dir db/migrations -seq $(name)

dropdb:
	migrate -path db/migrations -database $(DATABASE_URL) drop

up:
	migrate -path db/migrations -database $(DATABASE_URL) up

down:
	migrate -path db/migrations -database $(DATABASE_URL) down

sqlc:
	sqlc generate

test:
	go test -v -cover ./db/tests