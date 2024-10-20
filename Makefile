# Makefile for user-service
# vim: set ft=make ts=8 noet
# Licence MIT

# Variables
# UNAME		:= $(shell uname -s)
PWD = $(shell pwd)
PROJECT_NAME = user-service
API := api
DOCKER_COMPOSE=docker-compose -p ${PROJECT_NAME} -f ${PWD}/ops/docker/docker-compose.yml
GREEN=\033[0;32m
RESET=\033[0m

.EXPORT_ALL_VARIABLES:

# this is godly
# https://news.ycombinator.com/item?id=11939200
.PHONY: help
help:
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@# this is not tested, but prepared in advance for you, Mac drivers
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

tests: ### Run tests
	go test -v ./...

start: build run ### Start the application


restart : stop start ### Restart the application

build:
	@${DOCKER_COMPOSE} build

run:
	@${DOCKER_COMPOSE} up -d

stop: ### Stop the docker containers
	@${DOCKER_COMPOSE} down --remove-orphans

analysis: ### Run static analysis and linter
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:latest golangci-lint run

refresh-openapi: ### Generate openapi docs
	swag init -g cmd/api/main.go -g internal/platform/server/routes.go -o docs

open-docs:
	open http://localhost:9091/swagger/index.html

generate-mocks:
	docker run -v "$PWD":/src -w /src vektra/mockery --all