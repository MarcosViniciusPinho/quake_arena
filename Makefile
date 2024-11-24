.PHONY: help build build-linux build-windows build-macos clean test-report

help:
	@echo "Available commands:"
	@echo "  make build               - Builds binaries for Linux, Windows, and macOS and places them in the output directory"
	@echo "  make build-linux         - Compiles the binary for Linux"
	@echo "  make build-windows       - Compiles the binary for Windows"
	@echo "  make build-macos         - Compiles the binary for macOS"
	@echo "  make test-report         - Runs all tests and generates a coverage report"
	@echo "  make clean               - Cleans the output directory"
	@echo "  make help                - Displays this help message"

build: clean build-linux build-windows build-macos
	@echo "\n✅ Build completed successfully."

build-linux:
	@echo "\n🔨 Building Linux executable..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o out/linux/app_quake ./cmd/cli/main.go

build-windows:
	@echo "\n🔨 Building Windows executable..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o out/windows/app_quake.exe ./cmd/cli/main.go

build-macos:
	@echo "\n🔨 Building macOS executable..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o out/macos/app_quake ./cmd/cli/main.go

clean:
	@echo "\n🧹 Cleaning output directory..."
	rm -rf ./out
	@echo "✅ Clean completed."

test-report:
	@echo "\n🧪 Running tests and generating coverage report..."
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	@echo "✅ Test report generated: cover.html"