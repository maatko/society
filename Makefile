build:
	@templ generate
	@go build -o ./tmp/society cmd/society/main.go

run:
	./tmp/society

watch:
	@tailwindcss -c ./internal/tailwind.config.js -i ./internal/css/style.css -o ./web/static/css/global.css -w -m