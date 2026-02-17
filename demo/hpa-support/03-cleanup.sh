#!/usr/bin/env bash
# 03-cleanup.sh — Remove demo resources and optionally the kind cluster.
#
# Usage:
#   ./03-cleanup.sh           # Remove demo resources only
#   ./03-cleanup.sh --all     # Also delete the kind cluster

set -euo pipefail

NAMESPACE="gloo-system"
CLUSTER_NAME="${CLUSTER_NAME:-kind}"

echo "=== Removing demo Gateways ==="
kubectl delete gateway hpa-overlay-gateway \
  -n "$NAMESPACE" --ignore-not-found 2>&1

echo ""
echo "=== Removing demo GatewayParameters ==="
kubectl delete gatewayparameters hpa-overlay-gwp \
  -n "$NAMESPACE" --ignore-not-found 2>&1

echo ""
echo "=== Removing generated HPA (if still present) ==="
kubectl delete hpa gloo-proxy-hpa-overlay-gateway \
  -n "$NAMESPACE" --ignore-not-found 2>&1

echo ""
echo "=== Waiting for proxy deployment to be cleaned up ==="
for dep in gloo-proxy-hpa-overlay-gateway; do
  for i in $(seq 1 30); do
    if ! kubectl get deployment "$dep" -n "$NAMESPACE" &>/dev/null; then
      echo "  $dep removed"
      break
    fi
    sleep 1
  done
done

if [ "${1:-}" = "--all" ]; then
  echo ""
  echo "=== Uninstalling gloo helm release ==="
  helm uninstall gloo -n "$NAMESPACE" 2>/dev/null || true
  kubectl delete namespace "$NAMESPACE" --ignore-not-found 2>/dev/null || true

  echo ""
  echo "=== Deleting kind cluster '$CLUSTER_NAME' ==="
  kind delete cluster --name "$CLUSTER_NAME"
  echo "Cluster deleted."
else
  echo ""
  echo "Remaining resources:"
  kubectl get deployments -n "$NAMESPACE"
  echo ""
  echo "To also delete the kind cluster, run: $0 --all"
fi
