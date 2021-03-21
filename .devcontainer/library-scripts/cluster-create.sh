#!/bin/bash
# Unofficial bash strict mode: http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail
IFS=$'\n\t'

# Create the cluster and set $HOME/.kube/config
k3d cluster create test
kubectl config use-context k3d-test

# Update the k3d KubeAPI IP to the host, where the docker containers are running
sed -i -e "s/0.0.0.0/host.docker.internal/g" $HOME/.kube/config