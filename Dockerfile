FROM golang:1.21.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy