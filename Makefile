.DEFAULT_GOAL := help

# Configuration of different directories
DIST_DIR := dist
BUILD_DIR := build

NAME := github-release-tool
EXECUTABLE := $(NAME)

TAGS ?=

VERSION := $(shell cat VERSION)

# Configuration for golang.mk
include golang.mk

.PHONY: build
build: golang-build ## Wrapper for golang-build

.PHONY: test
test: golang-test ## Wrapper for golang-test

.PHONY: clean
clean: golang-clean ## Wrapper for golang-clean

.PHONY: help
help: ## Prints usage help for this makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(:|: .*?## )"}; {printf "\033[36m%-30s\033[0m %s\n", $$(NF-1), $$(NF)}'

.PHONY:
goals: goals ## Prints all available goals
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(:|: .*?## )"}; {printf "%s\n", $$(NF-1)}'
