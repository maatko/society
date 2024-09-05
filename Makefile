APP_NAME := secrete
SOURCE_DIR := .
BUILD_DIR := ./tmp

all: build

clean:
	@rm -rf $(BUILD_DIR)

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/secrete/main.go

run:
	$(BUILD_DIR)/$(APP_NAME)
