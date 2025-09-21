# ---- Build stage ----
FROM golang:1.25.1 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN go build -o server ./main.go

# ---- Run stage ----
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy only the built binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
