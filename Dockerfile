FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o scout-cli ./cmd/scout-cli

FROM debian:stable-slim

WORKDIR /root/

COPY --from=builder /app/scout-cli .

CMD ["./scout-cli"]