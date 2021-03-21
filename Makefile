# Assumes you are working in the .devcontainer

.DEFAULT_GOAL := help

cluster_create: ## Create a k3d cluster named `test`
	k3d cluster create test
	kubectl config use-context k3d-test
	# Update the k3d localhost KubeAPI IP to the host IP where the docker containers are running
	sed -i -e "s/0.0.0.0/host.docker.internal/g" ${HOME}/.kube/config

cluster_delete: ## Delete the k3d `test` cluster
	k3d cluster delete test

help: ## List the targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run the Go unit tests
	cd test && go mod tidy &&	go test -v

terraform_docs: ## Generate README docs
	terraform-docs markdown .

terraform_lint: ## Terraform lint the module
	tflint .

.PHONY: \
	cluster_create \
	cluster_delete \
	help \
	test \
	terraform_lint