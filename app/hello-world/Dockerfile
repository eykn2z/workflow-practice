# Build stage
FROM golang:1.18-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# Go Modulesを使用する場合はmodファイルをコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係を取得
RUN go mod download

# ソースコードをコピー
COPY . .

# Goのビルド
RUN go build -o api-server

# Run stage
FROM alpine:latest

# SSL証明書を追加
RUN apk --no-cache add ca-certificates

# 作業ディレクトリを設定
WORKDIR /root/

# ビルドされたバイナリをコピー
COPY --from=builder /app/api-server .

# コンテナ起動時に実行されるコマンド
CMD ["./api-server"]