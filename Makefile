setup:
	@go mod download

build:
	@go build ./...

run: build
	@go install github.com/allanassis/todo-list
	@todo-list

services-up:
	@docker-compose up -d