FROM golang:alpine AS builder

WORKDIR /app
COPY . .

RUN cd backend && GOOS=linux go build cmd/server/main.go

# Run the binary
ENTRYPOINT ["/app/backend/main"]