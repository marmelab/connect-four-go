.PHONY: install run test help
.DEFAULT_GOAL := help

help:
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Initialization ===============================================================

install:
	docker build --tag=marmelab-go .

# Deployment ===================================================================

pkg/darwin_amd64/connectfour.a: src/connectfour/renderer/renderer.go src/connectfour/board.go
	@docker run --rm --volume="`pwd`:/srv" -ti marmelab-go bash -c	"cd src/connectfour && go install"

bin/main: src/main/main.go pkg/darwin_amd64/connectfour.a
	@docker run --rm --volume="`pwd`:/srv" -ti marmelab-go bash -c	"cd src/main && go install"

# Development ==================================================================

run:
	@docker run --rm --volume="`pwd`:/srv" -ti marmelab-go bin/main -file=${FILE}

# Tests ========================================================================

test: ## Run all tests
	@docker run --rm --volume="`pwd`:/srv" -ti marmelab-go bash -c	"cd ./src/connectfour/test && go test"
