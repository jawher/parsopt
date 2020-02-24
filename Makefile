test: gotest batstest

gotest:
	go test

batstest: parsopt
	./test/libs/bats/bin/bats test/*.bats

check:
	bin/golangci-lint run

parsopt:
	go build

.PHONY: build gotest batstest check setup
