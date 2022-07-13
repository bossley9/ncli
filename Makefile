PREFIX = /usr/local
BIN = $(PREFIX)/bin
EXE = sn

build: prep
	go build -o $(EXE) ./cmd/main.go

prep:
	sh ./set-token.sh

run:
	go run ./cmd/main.go

clean:
	rm ./$(EXE)

install:
	mkdir -p $(BIN)
	cp -f ./$(EXE) $(BIN)
	chmod 555 $(BIN)/$(EXE)

uninstall:
	rm -f $(BIN)/$(EXE)

.PHONY: build prep run clean install uninstall
