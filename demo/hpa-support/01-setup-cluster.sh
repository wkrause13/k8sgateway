#!/usr/bin/env bash
# 01-setup-cluster.sh — Create a kind cluster, build images, and install gloo.
#
# Usage:
#   ./01-setup-cluster.sh
#
# Prerequisites:
#   - Docker running
#   - kind, kubectl, helm installed
#
# Environment variables (all optional):
#   VERSION        Image tag to use (default: 1.0.0-hpa-dev)
#   CLUSTER_NAME   Kind cluster name (default: kind)

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
VERSION="${VERSION:-1.0.0-hpa-dev}"
CLUSTER_NAME="${CLUSTER_NAME:-kind}"

cd "$REPO_ROOT"

echo "=== Step 1: Create kind cluster ==="
if kind get clusters 2>/dev/null | grep -qx "$CLUSTER_NAME"; then
  echo "Kind cluster '$CLUSTER_NAME' already exists, skipping creation."
else
  JUST_KIND=true make kind-setup
fi

kubectl config use-context "kind-${CLUSTER_NAME}"
kubectl cluster-info

echo ""
echo "=== Step 2: Build and load images ==="
VERSION="$VERSION" make kind-build-and-load

echo ""
echo "=== Step 3: Build test helm chart ==="
VERSION="$VERSION" make build-test-chart

echo ""
echo "=== Step 4: Install Gateway API CRDs ==="
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.2.1/standard-install.yaml

echo ""
echo "=== Step 5: Apply local GatewayParameters CRD ==="
# Helm does not upgrade CRDs in the chart's crds/ directory on helm upgrade.
# Apply the local CRD explicitly so new schema fields are available.
kubectl apply -f "${REPO_ROOT}/install/helm/gloo/crds/gateway.gloo.solo.io_gatewayparameters.yaml"

echo ""
echo "=== Step 6: Install gloo ==="
kubectl create namespace gloo-system 2>/dev/null || true
helm upgrade --install -n gloo-system gloo \
  "${REPO_ROOT}/_test/gloo-${VERSION}.tgz" \
  --set kubeGateway.enabled=true \
  --set discovery.enabled=false \
  --set gloo.deployment.image.tag="${VERSION}" \
  --set global.image.tag="${VERSION}" \
  --set global.image.registry=quay.io/solo-io

echo ""
echo "=== Restarting gloo to pick up the freshly loaded image ==="
kubectl rollout restart deployment/gloo -n gloo-system

echo ""
echo "=== Waiting for gloo pods ==="
kubectl rollout status deployment/gloo -n gloo-system --timeout=120s

echo ""
echo "=== Cluster ready ==="
kubectl get pods -n gloo-system
echo ""
echo "Default GatewayParameters:"
kubectl get gatewayparameters -n gloo-system -o yaml | grep -A5 "deployment:"
echo ""
echo "CRD deployment schema:"
kubectl get crd gatewayparameters.gateway.gloo.solo.io \
  -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema.properties.spec.properties.kube.properties.deployment}' \
  | python3 -m json.tool
echo ""
echo "CRD horizontalPodAutoscaler schema:"
kubectl get crd gatewayparameters.gateway.gloo.solo.io \
  -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema.properties.spec.properties.kube.properties.horizontalPodAutoscaler}' \
  | python3 -m json.tool
