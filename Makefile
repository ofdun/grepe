APP_NAME = grepe
VERSION = 0.1
BUILD_DIR = build
SRC_DIR = ./
BIN_DIR = /usr/local/bin
MAN_DIR = /usr/share/man/man1
GO = go
LINT_UTIL := golangci-lint

all: build

lint:
	@echo "Linting..."
	$(LINT_UTIL) run -v -c ./.golangci-lint.yml

build:
	@echo "Building $(APP_NAME)..."
	$(GO) build -o ./$(BUILD_DIR)/$(APP_NAME) ./$(SRC_DIR)

install: build
	@echo "Installing $(APP_NAME)..."
	sudo cp ./$(BUILD_DIR)/$(APP_NAME) $(BIN_DIR)/$(APP_NAME)
	# sudo cp $(APP_NAME).1 $(MAN_DIR)/$(APP_NAME).1
	@echo "$(APP_NAME) installed successfully."

uninstall:
	@echo "Uninstalling $(APP_NAME)..."
	sudo rm -f $(BIN_DIR)/$(APP_NAME)
	# sudo rm -f $(MAN_DIR)/$(APP_NAME).1
	@echo "$(APP_NAME) uninstalled successfully."

clean:
	@echo "Cleaning build files..."
	rm -rf ./$(BUILD_DIR)
	@echo "Cleaned."

help:
	@echo "Usage:"
	@echo "  make build        - Build the application"
	@echo "  make install      - Install the application"
	@echo "  make uninstall    - Uninstall the application"
	@echo "  make clean        - Clean build files"
	@echo "  make help         - Show this help message"

.PHONY: all build install uninstall clean help lint
