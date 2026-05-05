APP=wgmsg

.PHONY: run build test fmt

run:
	go run ./cmd/$(APP) $(ARGS)

build:
	go build -o bin/$(APP) ./cmd/$(APP)

test:
	go test ./...

fmt:
	go fmt ./...