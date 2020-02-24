test: gotest batstest

gotest:
	go test

batstest: parsopt
	./test/libs/bats/bin/bats test/*.bats

parsopt: *.go
	go build

check:
	bin/golangci-lint run

lint:
	golint -set_exit_status .

fmtcheck:
	@ export output="$$(gofmt -s -d ./*.go)"; \
		[ -n "$${output}" ] && echo "$${output}" && export status=1; \
		exit $${status:-0}

setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.23.6

.PHONY: build gotest batstest check setup
