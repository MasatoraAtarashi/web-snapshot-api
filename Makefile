.PHONY: install-golangci-lint
install-golangci-lint:
	which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v1.43.0

.PHONY: test-all
test-all:
	make lint
	make test

.PHONY: test
test:
	go clean ./... && go test ./...

.PHONY: lint
lint:
	make install-golangci-lint
	go clean ./... && golangci-lint cache clean
	golangci-lint run
