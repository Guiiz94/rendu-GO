FROM golang:1.17-alpine AS builder

WORKDIR /myapp
COPY . /myapp

RUN go mod download


CMD ["go", "run", "/myapp/main.go"]