test: gotest batstest

gotest:
	go test

batstest: parsopt
	./test/libs/bats/bin/bats test/*.bats

parsopt: *.go
	go build

check: lint vet fmtcheck ineffassign

lint:
	golint -set_exit_status .

vet:
	go vet

fmtcheck:
	@ export output="$$(gofmt -s -d ./*.go)"; \
		[ -n "$${output}" ] && echo "$${output}" && export status=1; \
		exit $${status:-0}

ineffassign:
	ineffassign .

setup:
	go get github.com/Masterminds/glide
	go get github.com/gordonklaus/ineffassign
	go get github.com/golang/lint/golint
	glide up

.PHONY: build gotest batstest check lint vet fmtcheck ineffassign setup
