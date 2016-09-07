.PHONY: build test help
.DEFAULT_GOAL := help

help:
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Initialization ===============================================================

# Deployment ===================================================================

build: ## Build all
	@cd ./src/connectfour && go install
	@cd ./src/main && go install

# Development ==================================================================

# Tests ========================================================================
test: ## Run all tests
	@cd ./src/connectfour && go test
