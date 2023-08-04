GIT_COMMIT := $(shell git rev-parse --short HEAD)
DEFAULT_TAG := $(shell git tag -l -n1 --sort=-creatordate | head -1 | cut -d' ' -f1)
build_version:
	VERSION ?= $(DEFAULT_TAG)

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/simple-login-linux-x64 bootstrap/*
build-docker:
	mv dist/* docker/
	docker build --platform=linux/amd64 -t simple-login:latest ./docker

