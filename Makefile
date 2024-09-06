run:
	@templ generate
	@go run cmd/secrete/main.go

watch:
	@tailwindcss -c ./internal/tailwind.config.js -i ./internal/css/style.css -o ./web/static/css/global.css -w -m