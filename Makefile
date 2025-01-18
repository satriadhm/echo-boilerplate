DB_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable

migrate-up:
	goose -dir ./migrations postgres $(DB_URL) up

migrate-down:
	goose -dir ./migrations postgres $(DB_URL) down

create-migration:
	goose create $(name) sql
