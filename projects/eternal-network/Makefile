ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ENTRYPOINT="cmd/core/main.go"
DISTPOINT="dist/main"

dev:
	go run ${ENTRYPOINT}

build:
	go build -o ./dist/ ${ENTRYPOINT}

run:
	./${DISTPOINT}

