build:
	go run github.com/99designs/gqlgen

run:
	go run server.go

test:
	go test ./... -v