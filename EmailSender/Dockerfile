# --- Build stage ---
FROM golang:1.25 AS builder

WORKDIR /EmailSender

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o EmailSender .

# --- Run stage ---
FROM alpine:latest

WORKDIR /EmailSender

COPY --from=builder /EmailSender/EmailSender .


CMD ["./EmailSender"]
