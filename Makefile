all: clean build test bench coverage

clean:
	go clean -testcache -cache
	rm -f demo/soak

build:
	go build
	go build -o demo/soak cli/soak/main.go

test:
	go test ./...

bench:
	go test -run=Benchmark -bench=. -benchmem

coverage:
	go test -cover -coverprofile=coverage.out
	go tool cover -func=coverage.out

