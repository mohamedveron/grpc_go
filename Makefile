NAME = $(notdir $(PWD))

VERSION = $(shell printf "%s.%s" \
		$$(git rev-list --count HEAD) \
		$$(git rev-parse --short HEAD))

BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

generate:
	@echo :: getting generator
	go get -v -d
	go run github.com/99designs/gqlgen generate


build:  $(OUTPUT)
	CGO_ENABLED=0 GOOS=linux go build -o bin/app \
		-ldflags "-X main.version=$(VERSION)" \
		-gcflags "-trimpath $(GOPATH)/src"


run:
	@echo :: start http server at port 9090
	go run server.go

all: generate build run


$(OUTPUT):
	mkdir -p $(OUTPUT)