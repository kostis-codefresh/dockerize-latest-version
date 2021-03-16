FROM golang:1.15.1-alpine3.12 AS build-env

WORKDIR /tmp/workdir

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM alpine:3.12

RUN apk add --no-cache jq ca-certificates bash

COPY --from=build-env /tmp/workdir/dockerize-latest-release /app/dockerize-latest-release

WORKDIR /app

CMD ["./dockerize-latest-release"]
