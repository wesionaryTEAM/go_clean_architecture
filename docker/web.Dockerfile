FROM golang:alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add --no-cache git

RUN apk add inotify-tools

RUN echo $GOPATH

COPY . /clean_web

RUN go install github.com/rubenv/sql-migrate/...@latest

WORKDIR /clean_web

RUN go mod download

RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD sh /clean_web/docker/run.sh