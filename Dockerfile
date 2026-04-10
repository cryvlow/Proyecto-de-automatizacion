FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o scout-cli ./cmd/scout-cli

FROM debian:stable-slim

WORKDIR /root/

COPY --from=builder /app/scout-cli .

CMD ["./scout-cli"]