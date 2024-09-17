# マルチアーキテクチャ対応のUbuntuベースイメージを使用
FROM --platform=$BUILDPLATFORM ubuntu:latest

# ARGを使用してビルド時のアーキテクチャを取得
ARG TARGETARCH

# 環境変数を設定
ENV DEBIAN_FRONTEND=noninteractive
ENV GO_VERSION=1.21.3
ENV PATH="/usr/local/go/bin:${PATH}"

# 必要なパッケージをインストール
RUN apt-get update && apt-get install -y \
    wget \
    git \
    curl \
    unzip \
    python3 \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*

# アーキテクチャに応じてGoをインストール
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        wget https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz -O go.tar.gz; \
    else \
        wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -O go.tar.gz; \
    fi && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

# アーキテクチャに応じてAWS CLIをインストール
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip"; \
    else \
        curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"; \
    fi && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf aws awscliv2.zip

# AWS SAM CLIのインストール
RUN pip3 install aws-sam-cli

# 作業ディレクトリを設定
WORKDIR /app

# Go Modulesを有効化
ENV GO111MODULE=on

# ソースコードをコピー
COPY . .

# デフォルトコマンドを設定（シェルを起動）
CMD ["/bin/bash"]