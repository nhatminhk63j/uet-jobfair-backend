FROM golang:1.17 AS builder

# Set the Current Working Directory inside the container.
WORKDIR /app

COPY go.mod go.sum /app/

# go mod download.
RUN set -eux; \
    GOSUMDB=off go mod download

# Copy source and build project.
COPY . /app/
RUN go build -o jobfair cmd/server/main.go

FROM debian:10-slim

COPY --from=builder /app/jobfair /app/

EXPOSE 8080

RUN chmod +x /app/jobfair
ENTRYPOINT ["/app/jobfair"]
