
all: build

build: deps test

deps:
	go get ./...

test:
	golint ./...
	go test -cover ./...
	go tool vet **/*.go


