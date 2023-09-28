FROM golang:1.21.1-alpine AS builder

WORKDIR /myapp
COPY . /myapp

RUN go mod download
RUN go get github.com/gofiber/fiber/v2


CMD ["go", "run", "/myapp/main.go"]