lint-go: lint-install-go
	@golangci-lint run

lint-install-go:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test-go:
	@go test -v -parallel=5 ./...

cov-report:
	@rm -rf reports
	@mkdir reports
	@go test -covermode=count -coverpkg=./... -coverprofile reports/cover.out -v ./...
	@go tool cover -html reports/cover.out -o reports/cover.html

build: doc
	@rm -rf bin
	@mkdir bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o bin/auth cmd/main.go

testing: lint-go test-go cov-report

run-all: lint-go test-go cov-report build

doc:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go get -u github.com/swaggo/echo-swagger
	@go mod tidy
	@swag init -d cmd/web,pkg/app/handlers --parseDependency true --parseInternal true

.PHONY: lint-go test-go cov-report build run-all
