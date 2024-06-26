# 使用官方的 Go 镜像作为构建阶段的基础镜像
FROM golang:1.22-alpine AS builder

# 设置环境变量
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# 创建和设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 并下载依赖项
COPY go.mod go.sum ./
RUN go mod download

# 复制项目中的所有代码
COPY . .

# 编译应用程序
RUN go build -o backend ./cmd

# 使用一个轻量级的基础镜像来运行应用程序
FROM alpine:latest

# 创建工作目录
WORKDIR /root/

# 接受项目名作为构建参数
ARG PROJECT_NAME=life-restart-backend

# 复制编译好的二进制文件
COPY --from=builder /app/backend .

# 将配置文件复制到容器中
COPY --from=builder /app/configs ./configs

# 暴露应用程序端口
EXPOSE 8080

# 启动应用程序
CMD ["./backend"]
