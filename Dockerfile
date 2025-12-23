# --- 构建阶段 ---
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
# CGO_ENABLED=0 保证生成静态二进制文件
# GOOS=linux 保证在 Mac 上也能打出 Linux 容器可用的包
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# --- 运行阶段 ---
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]