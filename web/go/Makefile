.PHONY: run build test lint sec secrets check clean

run:
	go run .

build:
	go build -o bin/webapp .

test:
	go test -race -cover ./...

lint:
	go vet ./...
	go tool staticcheck ./...

sec:
	go tool gosec ./...
	go tool govulncheck ./...

secrets:
	go tool gitleaks dir .

check: test lint sec secrets

clean:
	rm -rf bin
