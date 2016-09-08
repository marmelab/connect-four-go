.PHONY: install run test benchmark help
.DEFAULT_GOAL := help

GO_BIN := docker run \
	--interactive \
	--rm \
	--tty \
	--volume="$(CURDIR):/srv" \
	marmelab-go

help:
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Initialization ===============================================================

install:
	docker build --tag=marmelab-go .

# Deployment ===================================================================

pkg/darwin_amd64/connectfour.a: src/connectfour/renderer/renderer.go src/connectfour/board.go
	$(GO_BIN) bash -c "cd src/connectfour && go install"

bin/main: src/main/main.go pkg/darwin_amd64/connectfour.a
	$(GO_BIN) bash -c "cd src/main && go install"

# Development ==================================================================

run:
	$(GO_BIN) bin/main -file=${FILE}

# Tests ========================================================================

test: ## Run all tests
	$(GO_BIN) bash -c "cd src/connectfour && go test"
	$(GO_BIN) bash -c "cd src/connectfour/renderer && go test"

benchmark: ## Run all the benchmarks
	$(GO_BIN) bash -c "cd src/connectfour && go test -run=XXX -bench=."
