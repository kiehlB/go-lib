APP_NAME=todo
SRC_PATH=.
BIN_PATH=./bin
GOFLAGS=-ldflags "-s -w"

build:
	@echo "Building $(APP_NAME)..."
	@go build $(GOFLAGS) -o $(BIN_PATH)/$(APP_NAME) $(SRC_PATH)/main.go

run:
	@echo "Running $(APP_NAME)..."
	@./$(BIN_PATH)/$(APP_NAME)

test:
	@echo "Testing $(APP_NAME)..."
	@go test $(SRC_PATH)/...

install:
	@go mod tidy
	
help:
	@echo "Usage: make [command]"
	@echo "Commands:"
	@echo "  build    - Build $(APP_NAME)"
	@echo "  run      - Run $(APP_NAME)"
	@echo "  test     - Test $(APP_NAME)"
	@echo "  install     - Install $(APP_NAME)"
	@echo "  help     - Show help message"
	 

default: build