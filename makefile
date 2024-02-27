# This is the name of the binary file that will be created
BIN := server

# This is the directory where the binary will be created
BIN_DIR := .

# This is the directory of the main.go file
SRC_DIR := ./backend/cmd/server
    
# This will create the bin directory if it doesn't exist
$(shell mkdir -p $(BIN_DIR))

# This will build the binary
.PHONY: build
build:
	go build -o $(BIN_DIR)/$(BIN) $(SRC_DIR)

# Use this to run the build
.PHONY: run
run:
	./backend/scripts/run_server.sh

# Use this to clean the build
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

# Use this to build, run and clean
.PHONY: all
all: clean build run