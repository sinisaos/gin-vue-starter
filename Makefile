server:
	go run cmd/main.go

generate:
	swag init -g cmd/main.go --output api/docs

format:
	swag fmt
