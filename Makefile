test: gotest batstest

gotest:
	go test

batstest: parsopt
	./test/libs/bats/bin/bats test/*.bats

golint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.50 golangci-lint run -v

parsopt:
	go build

.PHONY: build gotest batstest check setup
