FROM golang:1.26.0-alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY dagosy-worker /app/dagosy-worker

ENTRYPOINT ["/app/dagosy-worker"]