# --- Build stage ---
FROM golang:1.25 AS builder

WORKDIR /WatchlistAPI

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o WatchlistAPI .

# --- Run stage ---
FROM alpine:latest

WORKDIR /WatchlistAPI

# Copy the binary from the builder stage
COPY --from=builder /WatchlistAPI/WatchlistAPI .

EXPOSE 8080

# Run the app
CMD ["./WatchlistAPI"]
