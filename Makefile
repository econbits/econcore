SHELL := zsh
ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

.PHONY: clean help vet test cov test-all gen
.DEFAULT_GOAL := help

help:
> @grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## remove object files and cached files
> go clean

vet: ## report likely mistakes in packages
> go vet github.com/econbits/econcore/...

gen: ## auto generates ISO4217 country list
> go generate github.com/econbits/econcore/...

test: ## run tests
> go test -v github.com/econbits/econcore/...

cov: ## gets the test coverage for the code and fails if minimum level is not reached
> go test -coverprofile=coverage.out github.com/econbits/econcore/...
> go tool cover -html=coverage.out

test-all: gen test vet cov ## runs all the checks to run before pushing to github

docs: ## generate documentation
> go doc -all
