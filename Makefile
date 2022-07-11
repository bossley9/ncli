EXE = sn

build: prep
	go build -o $(EXE) ./cmd/main.go

prep:
	sh ./set-token.sh

run:
	go run ./cmd/main.go


.PHONY: build run
