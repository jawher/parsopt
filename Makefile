test: gotest batstest

gotest:
	go test

batstest: parsopt
	./test/libs/bats/bin/bats test/*.bats

parsopt: *.go
	go build

check: lint vet fmtcheck ineffassign

lint:
	bin/golangci-lint run

vet:
	go vet

fmtcheck:
	@ export output="$$(gofmt -s -d ./*.go)"; \
		[ -n "$${output}" ] && echo "$${output}" && export status=1; \
		exit $${status:-0}

ineffassign:
	ineffassign .

setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.17.1
	go mod download

.PHONY: build gotest batstest check lint vet fmtcheck ineffassign setup
