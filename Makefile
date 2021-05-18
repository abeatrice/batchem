# DEFAUL VALUES
# Choose in case you have multiple profiles on your workstation
aws_profile ?= default
# Select application environment (dev, stage, prod)
app_env ?= dev

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

null:
	help

# Makefile Command Line Arguments!
ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

CMD :=
DOCKER_LOGIN := aws ecr get-login-password | \
	docker login -u AWS --password-stdin \
	$$(aws sts get-caller-identity --query 'Account' --output text).dkr.ecr.us-east-1.amazonaws.com

.PHONY: build
build: ## Build the app
	docker-compose -f docker-compose.yml build

.PHONY: run
run: ## Run the app locally
	docker-compose -f docker-compose.yml up

%:    # Supresses the below 
	@:  # make: *** No rule to make target `...'.  Stop.
