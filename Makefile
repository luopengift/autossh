APP = autossh
MAIN = cmd/autossh/main.go
PREFIX = /usr/local
VERSION	= 0.1.3_092918

TIME = $(shell date "+%F %T")
GIT = $(shell git rev-parse HEAD)
PKG = github.com/luopengift/version

FLAG = "-X '${PKG}.VERSION=${VERSION}' -X '${PKG}.APP=${APP}' -X '${PKG}.TIME=${TIME}' -X '${PKG}.GIT=${GIT}'"

build:
	go build -ldflags $(FLAG) -o ${GOPATH}/bin/$(APP) $(MAIN)
install:
	mv $(APP) $(PREFIX)/bin
fmt:
	go fmt ./...
lint:
	go vet ./...
test:
	go test -short ./...
test-all: lint
	go test ./...
clean:
	rm -f $(APP)
	rm -f $(PREFIX)/bin/$(APP)
.PHONY: build install fmt lint test test-all clean
