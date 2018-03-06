.PHONY: \
	build

GO_PATH = $(shell echo $(GOPATH) | awk -F':' '{print $$1}')
GO_SRC = $(shell pwd | xargs dirname | xargs dirname | xargs dirname)
DEPLOY_PATH := /Users/arcila/ebi/dev/bioschemas-govalid/build/
BIN_NAME := bioschema-govalid
VERSION=$(shell git describe --tags)
BUILD=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.buildDate=${BUILD}"

build:
	go build -i -v -o $(DEPLOY_PATH)$(BIN_NAME) main.go

install:
	go install -o $(DEPLOY_PATH) main.go

