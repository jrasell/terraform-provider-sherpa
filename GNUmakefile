default: lint build

.PHONY: tools
tools: ## Install the tools used to test and build
	@echo "==> Installing build tools"
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: build
build: ## Build the Sherpa Terraform provider
	@echo "==> Running $@..."
	go build

.PHONY: lint
lint: ## Run golangci-lint
	@echo "==> Running $@..."
	golangci-lint run sherpa/...

release: ## Trigger the release build script
	@echo "==> Running $@..."
	@goreleaser --rm-dist

HELP_FORMAT="    \033[36m%-25s\033[0m %s\n"
.PHONY: help
help: ## Display this usage information
	@echo "Sherpa Terraform provider make commands:"
	@grep -E '^[^ ]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
			{printf $(HELP_FORMAT), $$1, $$2}'
