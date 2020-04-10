.PHONY=build
build:
	go build ./cmd -o main
.PHONY=run
run:
	export PORT=8080
	go run ./cmd/main.go
default: build
