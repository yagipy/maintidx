build:
	go build -o bin/maintidx ./cmd/maintidx

test:
	go test -v ./...
