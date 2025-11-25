# Simple Dockerfile to build and run the CLI habit tracker
# Builds using Debian-based Go image to avoid CGO issues with sqlite3
FROM golang:1.21-buster AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o /habit ./cmd/habits

# Minimal runtime image
FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /habit /habit
WORKDIR /data
ENTRYPOINT ["/habit"]
