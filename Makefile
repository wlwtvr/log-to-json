PHONY: build-cli test clean run-cli

build-cli:
	go build -o log2json cmd/cli/main.go

build-wasm:
	GOOS=js GOARCH=wasm go build -o web/public/wasm/main.wasm wasm/main.go

build-fe:
	make build-wasm

test:
	go test ./...

clean:
	rm -f log2json

run-http:
	go run cmd/http/main.go

run-cli:
	go run cmd/cli/main.go