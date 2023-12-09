.PHONY: test build docker run

PROTO_BUILD_DIR = intermediate

# Where binary are put
TARGET_DIR ?= ./

DOCKER_OPTS ?= --rm

# TEST_ARGS = -v | grep -c RUN
VERSION := $(shell git describe --tags --abbrev=0)

DOCKER_TAG ?= $(VERSION)

SVC_NAME = github.com/grippenet/webhook-service

help:
	@echo "Service building targets"
	@echo "  build: build services (env TARGET_DIR to define binary location)"
	@echo "  docker: build docker container service service"

	@echo "Env:"
	@echo "  DOCKER_OPTS : default docker build options (default : $(DOCKER_OPTS))"
	@echo "  TEST_ARGS : Arguments to pass to go test call"

docker:
	docker build -t  $(SVC_NAME):$(DOCKER_TAG)  -f build/docker/Dockerfile $(DOCKER_OPTS) .

build:
	go build -o $(TARGET_DIR) ./cmd/webhook-service

run:
	go run ./cmd/webhook-service