FROM golang:alpine

COPY .env_example .env

# Copy dir to workdir
COPY . /go/src/bookstore

# Install git & bash
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/bookstore/

# Copy go mod and sum files
COPY go.mod go.sum /go/src/bookstore/

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Put all dependencies into a vendor folder
RUN go mod vendor
