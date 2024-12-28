# 使用するベースイメージ
FROM golang:1.23.4

# 必要なツールやライブラリをインストール
RUN apt-get update && apt-get install -y \
    git \
    bash \
    libc6 \
    curl \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# ワーキングディレクトリの設定
WORKDIR /myapp


RUN go install -v golang.org/x/tools/gopls@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/air-verse/air@latest