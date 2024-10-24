FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o chacha .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/chacha .

ENTRYPOINT ["./chacha"]