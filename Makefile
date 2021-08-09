default: build

build:
	go build -o bin/main cmd/ova-place-api/main.go

run:
	go run cmd/ova-place-api/main.go
