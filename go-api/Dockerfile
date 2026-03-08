# Build stage
FROM golang:1.26.1-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Build the binary with optimization flags:
# -s: Omit symbol table
# -w: Omit DWARF debug information (saves ~20-30% size)
RUN go build -ldflags="-s -w" -o /app/cloud-api ./cmd/server/main.go

# Final stage
FROM scratch

WORKDIR /

COPY --from=builder /app/cloud-api /cloud-api

# Best Practice: Run as a non-root user (security)
# Since scratch has no /etc/passwd, we use a numeric UID
USER 10001

EXPOSE 8080

ENTRYPOINT ["/cloud-api"]