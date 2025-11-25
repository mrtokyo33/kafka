BIN_NAME=kafka
BIN_DIR=./bin

MAIN_DIR=./cmd
MAIN_FILE=main.go

all:  build

run:
	@go run $(MAIN_DIR)/$(MAIN_FILE)

build:
	@go build -o $(BIN_DIR)/$(BIN_NAME) $(MAIN_DIR)/$(MAIN_FILE)

clean:
	@rm -f $(BIN_DIR)/$(BIN_NAME)
