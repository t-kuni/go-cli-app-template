# ベースイメージを指定
FROM mcr.microsoft.com/devcontainers/go:1.20

# 必要な追加ツールをインストール
RUN apt-get update && apt-get install -y \
    git \
    && apt-get clean && rm -rf /var/lib/apt/lists/*
