all: clean test build

.PHONY: clean
clean:
	go clean -testcache -cache
	rm -f demo/soak

.PHONY: build
build: clean test
	go build ./...
	go build -o demo/soak cli/soak/main.go

.PHONY: test
test:
	go test ./...

.PHONY: bench
bench:
	go test -run=Benchmark -bench=. -benchmem ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: coverage
coverage:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

