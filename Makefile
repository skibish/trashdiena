# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

VERSION = 1.1.0

.PHONY: test build push
.DEFAULT_GOAL := help

test: ## run tests
	go test -cover -race `go list ./... | grep -v /vendor/ | grep -v /cmd/`

build: ## build binaries for distribution
	docker build -t skibish/trashdiena:latest -t skibish/trashdiena:${VERSION} .

push: ## push images to the registry
	docker push skibish/trashdiena:latest
	docker push skibish/trashdiena:${VERSION}

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
