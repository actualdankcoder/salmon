TARGET_DIR := /usr/local/bin
BIN:= salmon

build:
	@go mod tidy
	@go build

static:
	@go mod tidy
	@go build -ldflags "-linkmode external -extldflags -static"

install:
	@mkdir -p $(TARGET_DIR)
	@cp ./$(BIN) $(TARGET_DIR)
	@chmod +x $(TARGET_DIR)/$(BIN)

check:
	@gofmt -w -s *.go
	@go test

uninstall:
	@rm $(TARGET_DIR)/$(BIN)

clean:
	@go clean

.PHONY: build install check uninstall clean all
