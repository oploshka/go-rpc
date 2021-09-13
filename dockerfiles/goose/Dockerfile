FROM golang:1.16-alpine AS builder

MAINTAINER Alexey Kibalnikov <alexey.kibalnikov@gmail.com>

#install GCC (GNU Compiler Collection)
RUN set -ex && \
    apk add --no-cache gcc musl-dev

#install goose CMD
RUN GO111MODULE=on go get -u 'github.com/pressly/goose/cmd/goose'

# copy only goose binary
FROM alpine

WORKDIR /app

COPY --from=builder /go/bin/goose ./goose

ARG GOOSE_DRIVER
ARG GOOSE_DBSTRING

CMD ["/bin/ash"]