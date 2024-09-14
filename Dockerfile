# ARM向けUbuntuの最新安定版をベースイメージとして使用
FROM ubuntu:latest

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

# ARM向けGo 1.21.3のインストール
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-arm64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-arm64.tar.gz \
    && rm go${GO_VERSION}.linux-arm64.tar.gz

# ARM向けAWS CLIのインストール
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf aws awscliv2.zip

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