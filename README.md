# About

Project template for Go CLI application.

# Features

* Onion Architecture
* Dependency Injection
* Testable

# Usage

1. fork this repository and clone it.
2. create .env file

```
cp .env.example .env
cp .env.test.example .env.test
```

# Build

```bash
make build
```

# Run

```bash
go run ./main.go command1
```

# Run test

```bash
make test
```

# Docker build & run

```bash
# Build
docker build -t go-cli-app-template .
# Run
docker run go-cli-app-template command1
```

# Devcontainer Setup

`devcontainer`を使用して開発環境をセットアップするには、以下の手順を実行してください。

1. VSCodeで当リポジトリを開きます
2. コマンドパレット（`Ctrl+Shift+P`）を開き、`Dev-Containers: Reopen in Container`を選択します。

これで、コンテナ内で開発環境がセットアップされます。

# TODO

* [ ] xxx