.PHONY: create-cluster
create-cluster:
	kind create cluster --name dem8s --config ./manifests/cluster.yml