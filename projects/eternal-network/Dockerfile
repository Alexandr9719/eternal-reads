# syntax=docker/dockerfile:1

FROM golang:1.21.5

ARG APP_EXPOSE_PORT=8443

WORKDIR /usr/lib/app
EXPOSE ${APP_EXPOSE_PORT}

COPY Makefile go.mod go.sum ./
COPY dist ./dist

CMD ["make", "run"]
