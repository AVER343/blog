.DEFAULT_GOAL = run

vet:
	go vet ./...

build:
	go build -o bin/main ./cmd/web
 

fmt:
	go fmt ./cmd/...

run: fmt vet build
	bin/blog
