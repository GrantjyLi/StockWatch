# --- Build stage ---
FROM golang:1.25 AS builder

WORKDIR /AlertsEvaluator

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o AlertsEvaluator .

# --- Run stage ---
FROM alpine:latest

WORKDIR /AlertsEvaluator

COPY --from=builder /AlertsEvaluator/AlertsEvaluator .


CMD ["./AlertsEvaluator"]
