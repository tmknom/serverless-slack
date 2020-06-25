# If this variable is not set in your makefile, the program /bin/sh is used as the shell.
# https://www.gnu.org/software/make/manual/html_node/Choosing-the-Shell.html
SHELL := /bin/bash

# This option causes make to display a warning whenever an undefined variable is expanded.
MAKEFLAGS += --warn-undefined-variables

# Disable any builtin pattern rules, then speedup a bit.
MAKEFLAGS += --no-builtin-rules

# Disable any builtin suffix rules, then speedup a bit.
.SUFFIXES:

# The arguments passed to the shell are taken from the variable .SHELLFLAGS.
#
# The -e flag causes bash with qualifications to exit immediately if a command it executes fails.
# The -u flag causes bash to exit with an error message if a variable is accessed without being defined.
# The -o pipefail option causes bash to exit if any of the commands in a pipeline fail.
# The -c flag is in the default value of .SHELLFLAGS and we must preserve it.
# Because it is how make passes the script to be executed to bash.
.SHELLFLAGS := -eu -o pipefail -c

# Sets the default goal to be used if no targets were specified on the command line.
.DEFAULT_GOAL := help

# https://gist.github.com/tadashi-aikawa/da73d277a3c1ec6767ed48d1335900f3
.PHONY: $(shell egrep -oh ^[a-zA-Z0-9][a-zA-Z0-9_-]+: $(MAKEFILE_LIST) | sed 's/://')

# The name defined in the "functions" section of serverless.yml
FUNCTION_NAME := slack

# The path defined in the "handler" section of serverless.yml
HANDLER_PATH := bin/main

clean: ## Clean the binary
	rm -rf bin

deps: ## Install dependencies
	go mod download
	go mod tidy

build: clean deps ## Build the application
	env GOOS=linux go build -ldflags="-s -w" -o $(HANDLER_PATH)

deploy: build ## Deploy a Serverless service
	sls deploy --verbose

generate-event: ## Generate event
	@sls generate-event --type aws:cloudWatch

invoke-local: build ## Invoke function locally
	@$(MAKE) generate-event | sls invoke local --function $(FUNCTION_NAME)

invoke-remote: ## Invoke function remotely
	@$(MAKE) generate-event | sls invoke --function $(FUNCTION_NAME) --log

logs: ## Output the logs of a deployed function
	sls logs --function $(FUNCTION_NAME)

remove: clean ## Remove Serverless service and all resources
	sls remove --verbose

help: ## Show help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
