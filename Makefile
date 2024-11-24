.PHONY: help build build-linux build-windows build-macos clean test-report

help:
	@echo "Choose one of the following commands:"
	@echo "  -> make build               - Creates binaries for Linux, Windows, and macOS and places them in the output directory"
	@echo "  -> make build-linux         - Compiles the binary for Linux"
	@echo "  -> make build-windows       - Compiles the binary for Windows"
	@echo "  -> make build-macos         - Compiles the binary for macOS"
	@echo "  -> make test-report         - Runs all tests and generates a coverage report"
	@echo "  -> make clean               - Cleans the output directory"
	@echo "  -> make help                - Displays this help message"

build: clean build-linux build-windows build-macos
	@echo "\nBuild completed successfully."

build-linux:
	@echo "\nCreating an executable artifact for the Linux environment..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o out/linux/app_quake ./cmd/cli/main.go

build-windows:
	@echo "\nCreating an executable artifact for the Windows environment..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o out/windows/app_quake.exe ./cmd/cli/main.go

build-macos:
	@echo "\nCreating an executable artifact for the macOS environment..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o out/macos/app_quake ./cmd/cli/main.go

clean:
	rm -rf ./out

test-report:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html