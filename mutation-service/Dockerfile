FROM golang:1.19.0 AS base
WORKDIR /usr/src/app

COPY . .
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

EXPOSE 5100

ENTRYPOINT [ "air", "./main.go", "-b", "0.0.0.0"]
