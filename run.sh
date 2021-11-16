#!/usr/bin/env sh

docker build -t kvmayer/go-users:latest .
docker push kvmayer/go-users:latest
kubectl delete -f deploy.yaml
kubectl apply -f deploy.yaml