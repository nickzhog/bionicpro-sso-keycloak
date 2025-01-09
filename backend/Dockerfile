
# Build stage
FROM golang:latest AS builder

LABEL stage=gobuilder
ENV GOOS linux
WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/backend

# Run stage
FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates
WORKDIR /app

EXPOSE 8000

COPY --from=builder /app/backend .

CMD ["./backend"]