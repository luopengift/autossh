APP = autossh
MAIN = cmd/autossh/main.go
PREFIX = /usr/local
PATH = ${GOPATH}/bin
GO = ${GOROOT}/bin/go
VERSION	= $(shell cat version)

TIME = $(shell date "+%F %T")
GIT = $(shell git rev-parse HEAD)
PKG = github.com/luopengift/version

FLAG = -ldflags "-X '${PKG}.VERSION=${VERSION}' -X '${PKG}.APP=${APP}' -X '${PKG}.TIME=${TIME}' -X '${PKG}.GIT=${GIT}'"

output = -o ${PATH}/$(APP)
darwin = -o ${PATH}/$(APP)_darwin
linux = -o ${PATH}/$(APP)_linux
windows = -o ${PATH}/$(APP)_windows

build:
	${GO} build ${FLAG} ${output} $(MAIN)

Darwin:
	${GO} build ${FLAG} ${darwin} $(MAIN)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build ${FLAG} ${linux} $(MAIN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ${GO} build ${FLAG} ${windows} $(MAIN)
Linux:
	${GO} build ${FLAG} ${linux} $(MAIN)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GO} build ${FLAG} ${darwin} $(MAIN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ${GO} build ${FLAG} ${windows} $(MAIN)
Windows:
	${GO} build ${FLAG} ${windows} $(MAIN)
	SET CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	${GO} build ${FLAG} ${darwin} $(MAIN)

	SET CGO_ENABLED=0
	SET GOOS=linux
	SET GOARCH=amd64
	${GO} build ${FLAG} ${linux} $(MAIN)
fmt:
	${GO} fmt ./...
lint:
	${GO} vet ./...
test:
	${GO} test -short ./...
test-all: lint
	${GO} test ./...
.PHONY: build install fmt lint test test-all clean
