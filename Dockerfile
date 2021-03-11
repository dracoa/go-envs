# build base stage
FROM golang:alpine AS build_base

WORKDIR /go/src

RUN set -ex; \
    apk add --no-cache curl git gcc musl-dev g++ libc-dev bash ca-certificates

ENV GO111MODULE=on

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

# build stage
FROM build_base AS build-env

WORKDIR /go/src

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app -ldflags="-s -w" ./main/main.go


# final stage
FROM alpine

RUN apk add ca-certificates

WORKDIR /root/

COPY --from=build-env /go/src/app .

ENTRYPOINT ["./app"]
