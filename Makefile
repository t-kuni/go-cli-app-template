build:
	go build -o app

generate_test:
	go generate ./...

test: generate_test
	gotestsum -- ./...