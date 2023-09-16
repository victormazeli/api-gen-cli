# Variables
APP_NAME := api-gen-cli
BUILD_DIR := build

.PHONY: release-local
release-local:
	goreleaser release --snapshot --skip-publish --clean


.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: mod-verify
mod-verify:
	go mod verify

.PHONY: mod-download
mod-download:
	go mod download


