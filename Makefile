build:
	@templ generate
	@go build -o ./tmp/main cmd/secrete/main.go

run:
	./tmp/main

watch:
	@tailwindcss -c ./internal/tailwind.config.js -i ./internal/css/style.css -o ./web/static/css/global.css -w -m