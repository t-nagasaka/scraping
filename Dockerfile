FROM golang:1.23.4

# 必要なツールやライブラリをインストール
RUN apt-get update && apt-get install -y \
    git \
    bash \
    libc6 \
    curl \
    && apt-get clean

# ブラウザに必要なライブラリをインストール
RUN apt-get install -y \
    chromium \
    libglib2.0-0 \
    libgirepository-1.0-1 \
    libnss3 \
    libatk1.0-0 \
    libx11-xcb1 \
    libxcomposite1 \
    libxrandr2 \
    libxss1 \
    libasound2 \
    libpangocairo-1.0-0 \
    libgtk-3-0 \
    libgbm1 \
    xvfb \
    xauth \
    --no-install-recommends && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Chromium用の環境変数を設定
ENV ROD_BROWSER_BIN=/usr/bin/chromium

# ワーキングディレクトリの設定
WORKDIR /myapp

# 必要なGoツールのインストール
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/air-verse/air@latest

CMD ["air","-c",".air.toml"]