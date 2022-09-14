FROM golang:1.18-alpine as build

ARG CGO_ENABLE=0
ARG GO111MODULE=on

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY . .

RUN go install ./...

FROM alpine

WORKDIR /app

COPY --from=build /go/bin/go-challenege .

RUN mkdir -p deployment  && mkdir -p message &&  mkdir -p dbmigration && mkdir -p docs && apk --no-cache add tzdata

