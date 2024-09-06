APP_NAME := secrete

build:
	@templ generate
	@go build -o ./tmp/$(APP_NAME) cmd/secrete/main.go

run: build
	./tmp/$(APP_NAME)

watch:
	@tailwindcss -c ./internal/tailwind.config.js -i ./internal/css/style.css -o ./web/static/css/global.css -w -m