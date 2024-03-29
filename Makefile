build: install-templ generate
install-templ:
	@echo "Installing templ..."
	@go install github.com/a-h/templ/cmd/templ@v0.2.598
generate:
	@echo "Generating..."
	@go run main.go -gen -ver v1.3.0
format-templ:
	@echo "Formatting templ files..."
	@templ fmt .
generate-templ:
	@echo "Generating templ files..."
	@templ generate -path ./components
generate-templ-watch:
	@echo "Generating templ files..."
	@templ generate -path ./components -watch
generate-tailwind:
	@echo "Generating tailwind files..."
	@tailwindcss -i ./styles/input.css -o ./dist/assets/css/output@dev.css
generate-tailwind-watch:
	@echo "Generating tailwind files..."
	@tailwindcss -i ./styles/input.css -o ./dist/assets/css/output@dev.css --watch
run:
	@echo "Running..."
	@go run main.go 
build:
	@echo "Building..."
	@go build -o ./app -ldflags="-s -w -X version.Value=1.0.0"
