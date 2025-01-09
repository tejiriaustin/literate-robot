# Stage 1: Build
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-service .

# Stage 2: Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/order-service .

EXPOSE 8082

CMD ["./order-service api"]
