# syntax=docker/dockerfile:1
## Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs

## Deploy
FROM debian:unstable-slim

WORKDIR /

COPY --from=build /docker-gs /docker-gs

ENTRYPOINT ["/docker-gs"]
