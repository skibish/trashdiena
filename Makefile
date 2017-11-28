# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

VERSION = 1.0.0

.PHONY: test build
.DEFAULT_GOAL := help

test: ## run tests
	go test -v -cover -race `go list ./... | grep -v /vendor/`

build: ## build binaries for distribution
	docker build -t trashdiena:latest -t trashdiena:${VERSION} .

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
