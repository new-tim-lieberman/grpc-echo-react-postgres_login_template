# Build Stage
FROM golang:1.26.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o api-service \
    ./api/cmd/api

# Runtime Stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=builder /app/api-service .

EXPOSE 8080

CMD ["./api-service"]