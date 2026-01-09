# --- Build stage ---
FROM golang:1.25 AS builder

WORKDIR /AlertsPoller

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o AlertsPoller .

# --- Run stage ---
FROM alpine:latest

WORKDIR /AlertsPoller

COPY --from=builder /AlertsPoller/AlertsPoller .


CMD ["./AlertsPoller"]
