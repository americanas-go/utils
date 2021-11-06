#!make

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' | \
	sort

tests: ## Run all tests
	@CGO_ENABLED=0 GOFLAGS="-count=1" go test  ./...

.PHONY: help tests
