BINARY_NAME = app
BINARY_DIR = bin
MAIN_PACKAGE = ./cmd/web
# ADDRESS = :4040
# CONNSTR = postgresql://postgres:hello123@localhost:5432/snippetbox?sslmode=disable

.PHONY : default build fmt run 

default : run

fmt: 
	@go fmt ./...

build: fmt
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

run: build
	@$(BINARY_DIR)/$(BINARY_NAME)