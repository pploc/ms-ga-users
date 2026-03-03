FROM golang:1.23-bookworm AS builder

# Install build dependencies for librdkafka (confluent-kafka-go)
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc libc6-dev pkg-config build-essential \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod ./
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application with CGO_ENABLED=1 for confluent-kafka-go
RUN go build -o main ./cmd/api

######## Start a new stage from scratch #######
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates tzdata libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copy the Pre-built binary
COPY --from=builder /app/main .

EXPOSE 8083

CMD ["./main"]
