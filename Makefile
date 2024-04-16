build:
	@go build -o bin/fs

run: build
	@go run main.go