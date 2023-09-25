build:
	go mod tidy
	go build -o bin/app app/main.go

run:build
	./bin/app