SHELL := /bin/bash

run:
	go run main.go

# Building containers

VERSION := 1.0

all: service-api

service-api:
	docker build \
		-f zarf/docker/dockerfile.service-api \
		-t service-api-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# Running from within k8s/kind

KIND_CLUSTER := service-cluster

kind-up:
	kind create cluster \
		--image kindest/node:v1.23.0@sha256:49824ab1727c04e56a21a5d8372a402fcd32ea51ac96a2706a12af38934f81ac \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml
	kubectl config set-context --current --namespace=service-system

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-load:
	cd zarf/k8s/kind/service-pod; kustomize edit set image service-api-image=service-api-amd64:$(VERSION)
	kind load docker-image service-api-amd64:$(VERSION) --name $(KIND_CLUSTER)

kind-apply:
	kustomize build zarf/k8s/kind/service-pod | kubectl apply -f -

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide watch --all-namespaces

kind-status-service:
	kubectl get pods -o wide --watch

kind-logs:
	kubectl logs -l app=service --all-containers=true -f --tail=100

kind-restart:
	kubectl rollout restart deployment service-pod

kind-update: all kind-load kind-restart

kind-update-apply: all kind-load kind-apply

kind-describe:
	kubectl describe pod -l app=service


# Module support

tidy:
	go mod tidy
	go mod vendor


