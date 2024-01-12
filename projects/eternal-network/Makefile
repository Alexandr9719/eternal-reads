ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DIST="dist"
ENTRYPOINT="cmd/core/main.go"
DISTPOINT="dist/main"

dev:
	go run ${ENTRYPOINT}

build:
	go build -o ./${DIST}/ ${ENTRYPOINT}

build-ci:
	CGO_ENABLED=0 GOOS=linux go build -o ./${DIST}/ ${ENTRYPOINT}

run:
	./${DISTPOINT}

test:
	go test ./... $(OPTS)

