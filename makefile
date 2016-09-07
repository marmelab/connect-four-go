.PHONY: run test help
.DEFAULT_GOAL := help

help:
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Initialization ===============================================================

# Deployment ===================================================================

pkg/darwin_amd64/connectfour.a: src/connectfour/renderer/renderer.go src/connectfour/board.go
	cd src/connectfour && go install

bin/main: src/main/main.go pkg/darwin_amd64/connectfour.a
	cd src/main && go install

# Development ==================================================================

run: bin/main ## Run executable (usage: make run FILE=myfilegame)
	@./bin/main -file=${FILE}

# Tests ========================================================================
test: ## Run all tests
	@cd ./src/connectfour/test && go test
