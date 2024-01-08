FROM golang:1.21.4-alpine

WORKDIR /app

COPY . .

RUN apk update && apk --no-cache add git && \
    go mod tidy && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]