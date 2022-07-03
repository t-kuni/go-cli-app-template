# About

This repository is project skeleton for Go CLI application.

# Features

* Onion Architecture
* Dependency Injection
* Testable

# Requirements

* go 1.18+
* [gotestsum](https://github.com/gotestyourself/gotestsum)
* [gomock](https://github.com/golang/mock)

# Usage

```bash
git clone --depth 1 ssh://git@github.com/t-kuni/go-cli-app-skeleton [ProjectName]
cd [ProjectName]
rm -rf .git 
git init
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
docker build -t go-cli-app-skeleton .
# Run
docker run go-cli-app-skeleton command1
```