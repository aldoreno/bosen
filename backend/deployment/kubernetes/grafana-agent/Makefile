.PHONY: all
all: help
help: ## Display help screen
	@echo "Usage:"
	@echo "	make [COMMAND]"
	@echo "	make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Uninstall prev installation and then delete configmap and secret
	helm uninstall grafana-agent-local --namespace monitoring --wait
	kubectl delete --namespace monitoring configmap grafana-agent-local
	kubectl delete --namespace monitoring secret lgtm-credentials

.PHONY: init-configmaps
init-configmaps: agent-config.river ## Initialise grafana-agent ConfigMap
	kubectl create --namespace monitoring configmap grafana-agent-local --from-file=./agent-config.river

.PHONY: init-secrets
init-secrets: lgtm-credentials.properties ## Initialise LGTM credentials Secret
	kubectl create --namespace monitoring secret generic lgtm-credentials --from-env-file=./lgtm-credentials.properties

.PHONY: install
install: init-configmaps init-secrets ## Install grafana-agent using Helm
	helm repo add grafana https://grafana.github.io/helm-charts
	helm repo update
	helm install grafana-agent-local grafana/grafana-agent -f values.yaml --namespace monitoring --wait

.PHONY: upgrade
upgrade: ## Upgrade grafana-agent installation using Helm
	helm repo update
	helm upgrade grafana-agent-local grafana/grafana-agent -f values.yaml --namespace monitoring --wait
