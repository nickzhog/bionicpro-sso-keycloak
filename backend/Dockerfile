# Build stage
FROM golang:1.23.0-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -ldflags="-s -w" -o /app/backend

# Run stage
FROM alpine
WORKDIR /app

COPY --from=builder /app/backend /app/backend

EXPOSE 8000

CMD ["/app/backend"]