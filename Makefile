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

pkg/linux_amd64/connectfour/parser.a: src/connectfour/parser/parser.go
	$(GO_BIN) bash -c "cd src/connectfour/parser && go install"

pkg/linux_amd64/connectfour/renderer.a: src/connectfour/renderer/renderer.go
	$(GO_BIN) bash -c "cd src/connectfour/renderer && go install"

pkg/linux_amd64/connectfour/ai.a: src/connectfour/ai/ai.go
	$(GO_BIN) bash -c "cd src/connectfour/ai && go install"

pkg/linux_amd64/connectfour/board.a: src/connectfour/board/board.go
	$(GO_BIN) bash -c "cd src/connectfour/board && go install"

pkg/linux_amd64/connectfour.a: src/connectfour/game.go
	$(GO_BIN) bash -c "cd src/connectfour && go install"

bin/main: src/main/main.go pkg/linux_amd64/connectfour.a pkg/linux_amd64/connectfour/ai.a pkg/linux_amd64/connectfour/renderer.a pkg/linux_amd64/connectfour/parser.a pkg/linux_amd64/connectfour/board.a
	$(GO_BIN) bash -c "cd src/main && go install"

# Development ==================================================================

run: bin/main
	$(GO_BIN) bin/main -file=${FILE}

# Tests ========================================================================

test: ## Run all tests
	$(GO_BIN) bash -c "cd src/connectfour && go test"
	$(GO_BIN) bash -c "cd src/connectfour/ai && go test"
	$(GO_BIN) bash -c "cd src/connectfour/board && go test"
	$(GO_BIN) bash -c "cd src/connectfour/renderer && go test"

benchmark: ## Run all the benchmarks
	$(GO_BIN) bash -c "cd src/connectfour/ai && go test -run=XXX -bench=."
