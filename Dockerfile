# Build stage
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Copy go mod và tải dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ mã nguồn
COPY . .

# Build ứng dụng
RUN go build -o main .

# Run stage
FROM debian:bookworm-slim

WORKDIR /app

# Cài đặt thư viện C cơ bản (nếu cần cho runtime)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary từ stage build
COPY --from=builder /app/main .

# Copy uploads (nếu dùng upload ảnh)
COPY --from=builder /app/uploads ./uploads

# Expose port
EXPOSE 8080

CMD ["./main"]
