GOOS=js
GOARCH=wasm
OUTPUT=main.wasm
MAIN_PATH=main
MAIN_SRC=$(MAIN_PATH)/main.go
SERVER_PATH=server
SERVER_SRC=$(SERVER_PATH)/server.go

compile:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT) $(MAIN_SRC)

run:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go run $(MAIN_SRC)

clean:
	rm -f $(OUTPUT)

tinygo:
	tinygo build -o $(OUTPUT) -target wasm $(MAIN_SRC)

server-start:
	go run $(SERVER_SRC)

go-start:
	$(MAKE) build
	$(MAKE) server-start

tinygo-start:
	$(MAKE) tinygo
	$(MAKE) server-start

.PHONY: compile build run clean tinygo server-start go-start tinygo-start