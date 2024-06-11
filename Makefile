PHONY: build test clean run-cli

build:
	go build -o log2json cmd/cli/main.go

test:
	go test ./...

clean:
	rm -f log2json

# run-web:
# 	go run cmd/web/main.go

run-cli:
	go run cmd/cli/main.go