# Build Stage
FROM golang:1.26.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o auth-service \
    ./auth/cmd/auth

# Runtime Stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=builder /app/auth-service .

EXPOSE 50052

CMD ["./auth-service"]