# About

This repository is project skeleton for TypeScript CLI application.

# Usage

```
git clone --depth 1 ssh://git@github.com/t-kuni/go-cli-app-skeleton [ProjectName]
cd [ProjectName]
rm -rf .git 
```

# Generate Mock

```
mockgen -source=clock.go -package=main -destination=clock_mock.go
```

# Build

```
docker build --tag example-container .
```

# Run

```
docker run example-container example
```