.PHONY: run build

run: build
	@./build/main FFFFFF

build:
	@go generate ./cmd/main.go
	@echo "[OK] Files were generated"
	@go build -o ./build/main ./cmd
	@echo "[OK] 'main' binary was created!"

test:
	@go generate ./cmd/main.go
	@go test ./cmd/...
