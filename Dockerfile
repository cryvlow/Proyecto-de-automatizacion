FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o scout-cli ./cmd/scout-cli

FROM alpine:3.20

RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/scout-cli /usr/local/bin/scout-cli

ENTRYPOINT ["scout-cli"]
CMD ["help"]