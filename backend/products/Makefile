lint-go: lint-install-go
	@golangci-lint run

lint-install-go:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test-go: ginkgo-install-go
	@ginkgo -r --show-node-events --fail-fast --randomize-all --randomize-suites --fail-on-pending --trace \
		--junit-report=junit.xml --output-dir=tests -coverpkg=./... -coverprofile=cover.out

ginkgo-install-go:
	@go install github.com/onsi/ginkgo/v2/ginkgo

cov-report: test-go
	@go test -covermode=count -coverpkg=./... -coverprofile tests/cover.out -v ./...
	@go tool cover -html tests/cover.out -o tests/cover.html

build: doc
	@rm -rf bin
	@mkdir bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o bin/products cmd/main.go

testing: lint-go test-go cov-report

run-all: lint-go test-go cov-report build

doc:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go get -u github.com/swaggo/echo-swagger
	@go mod tidy
	@swag init -d cmd/web,pkg/app/handlers --parseDependency true --parseInternal true

.PHONY: lint-go test-go cov-report build run-all