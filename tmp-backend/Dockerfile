# 使用官方 Go 镜像作为构建阶段
FROM golang:1.21-alpine AS builder

# 安装 git
RUN apk add --no-cache git

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译 Go 程序为可执行文件
RUN go build -o server main.go

# 使用轻量级镜像作为运行阶段
FROM alpine:latest

# 安装 ca-certificates
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/server .

# 暴露端口
EXPOSE 2121

# 启动程序
CMD ["./server"]
