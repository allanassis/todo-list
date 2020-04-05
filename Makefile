setup:
	@go mod download

build:
	@go build ./...

run:
	@go install github.com/allanassis/todo-list
	@todo-list