sqlc:
	sqlc generate

start:
	go run main.go

.PHONY: sqlc start