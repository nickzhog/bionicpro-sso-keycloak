# Build stage
FROM golang:1.23.0-alpine AS builder
WORKDIR /build
COPY . .
COPY resolve_fix.sh /app/resolve_fix.sh
RUN go mod download
RUN CGO_ENABLED=0 go build -a -ldflags="-s -w" -o /app/backend 

# Run stage 
FROM alpine
WORKDIR /app
USER root

RUN chown -R nobody:nogroup /app && adduser nobody nogroup
COPY --from=builder --chown=nobody:nogroup /app/backend /app/backend
COPY --from=builder --chown=nobody:nogroup /app/resolve_fix.sh /app/resolve_fix.sh
RUN chmod +x /app/resolve_fix.sh

# USER nobody
EXPOSE 8000

CMD ["/bin/sh", "-c", "/app/resolve_fix.sh && exec /app/backend"]