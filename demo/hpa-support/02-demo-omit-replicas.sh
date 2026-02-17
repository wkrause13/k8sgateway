#!/usr/bin/env bash
# 02-demo-omit-replicas.sh — Demonstrate GatewayParameters HPA overlay support.
#
# Verifies:
#   1. GatewayParameters with horizontalPodAutoscaler overlay is accepted
#   2. Gateway deployer creates an HPA automatically
#   3. HPA scaleTargetRef points to the generated gateway proxy Deployment
#   4. Overlay metadata/spec values are applied
#   5. Updating GatewayParameters updates the generated HPA
#
# Usage:
#   ./02-demo-omit-replicas.sh

set -euo pipefail

VERSION="${VERSION:-1.0.0-hpa-dev}"
NAMESPACE="${NAMESPACE:-gloo-system}"
GWP_NAME="${GWP_NAME:-hpa-overlay-gwp}"
GW_NAME="${GW_NAME:-hpa-overlay-gateway}"
DEPLOYMENT_NAME="gloo-proxy-${GW_NAME}"
HPA_NAME="${DEPLOYMENT_NAME}"
PASS=0
FAIL=0

pass() { PASS=$((PASS + 1)); echo "  PASS: $1"; }
fail() { FAIL=$((FAIL + 1)); echo "  FAIL: $1"; }

wait_for_exists() {
  local resource="$1"
  local timeout="${2:-90}"
  for i in $(seq 1 "$timeout"); do
    if kubectl get "$resource" -n "$NAMESPACE" >/dev/null 2>&1; then
      return 0
    fi
    sleep 1
  done
  return 1
}

wait_for_not_exists() {
  local resource="$1"
  local timeout="${2:-90}"
  for i in $(seq 1 "$timeout"); do
    if ! kubectl get "$resource" -n "$NAMESPACE" >/dev/null 2>&1; then
      return 0
    fi
    sleep 1
  done
  return 1
}

wait_for_jsonpath_value() {
  local resource="$1"
  local path="$2"
  local expected="$3"
  local timeout="${4:-90}"
  for i in $(seq 1 "$timeout"); do
    local actual
    actual="$(kubectl get "$resource" -n "$NAMESPACE" -o "jsonpath=${path}" 2>/dev/null || true)"
    if [ "$actual" = "$expected" ]; then
      return 0
    fi
    sleep 1
  done
  return 1
}

check_sa_can_i() {
  local verb="$1"
  local resource="$2"
  local scope="$3"
  if [ "$scope" = "cluster" ]; then
    kubectl auth can-i --as=system:serviceaccount:gloo-system:gloo "$verb" "$resource" --all-namespaces >/dev/null 2>&1
  else
    kubectl auth can-i --as=system:serviceaccount:gloo-system:gloo "$verb" "$resource" -n "$NAMESPACE" >/dev/null 2>&1
  fi
}

echo "=== Cleaning up prior demo resources (if any) ==="
kubectl delete gateway "$GW_NAME" -n "$NAMESPACE" --ignore-not-found >/dev/null 2>&1 || true
kubectl delete gatewayparameters "$GWP_NAME" -n "$NAMESPACE" --ignore-not-found >/dev/null 2>&1 || true
kubectl delete hpa "$HPA_NAME" -n "$NAMESPACE" --ignore-not-found >/dev/null 2>&1 || true
kubectl delete deployment "$DEPLOYMENT_NAME" -n "$NAMESPACE" --ignore-not-found >/dev/null 2>&1 || true

if wait_for_not_exists "gateway/${GW_NAME}" 60; then
  pass "Previous Gateway ${GW_NAME} removed"
else
  fail "Previous Gateway ${GW_NAME} still exists"
fi

if wait_for_not_exists "gatewayparameters/${GWP_NAME}" 60; then
  pass "Previous GatewayParameters ${GWP_NAME} removed"
else
  fail "Previous GatewayParameters ${GWP_NAME} still exists"
fi

if wait_for_not_exists "hpa/${HPA_NAME}" 60; then
  pass "Previous HPA ${HPA_NAME} removed"
else
  fail "Previous HPA ${HPA_NAME} still exists"
fi

if wait_for_not_exists "deployment/${DEPLOYMENT_NAME}" 90; then
  pass "Previous Deployment ${DEPLOYMENT_NAME} removed"
else
  fail "Previous Deployment ${DEPLOYMENT_NAME} still exists"
fi

echo ""
echo "=== Preflight: Verify controller RBAC for HPA/PDB ==="
if check_sa_can_i list horizontalpodautoscalers.autoscaling cluster; then
  pass "gloo SA can list HPAs cluster-wide"
else
  fail "gloo SA cannot list HPAs cluster-wide"
fi

if check_sa_can_i watch horizontalpodautoscalers.autoscaling cluster; then
  pass "gloo SA can watch HPAs cluster-wide"
else
  fail "gloo SA cannot watch HPAs cluster-wide"
fi

if check_sa_can_i create horizontalpodautoscalers.autoscaling namespaced; then
  pass "gloo SA can create HPAs"
else
  fail "gloo SA cannot create HPAs"
fi

if check_sa_can_i patch horizontalpodautoscalers.autoscaling namespaced; then
  pass "gloo SA can patch HPAs"
else
  fail "gloo SA cannot patch HPAs"
fi

if check_sa_can_i list poddisruptionbudgets.policy cluster; then
  pass "gloo SA can list PDBs cluster-wide"
else
  fail "gloo SA cannot list PDBs cluster-wide"
fi

if check_sa_can_i watch poddisruptionbudgets.policy cluster; then
  pass "gloo SA can watch PDBs cluster-wide"
else
  fail "gloo SA cannot watch PDBs cluster-wide"
fi

if check_sa_can_i create poddisruptionbudgets.policy namespaced; then
  pass "gloo SA can create PDBs"
else
  fail "gloo SA cannot create PDBs"
fi

if check_sa_can_i patch poddisruptionbudgets.policy namespaced; then
  pass "gloo SA can patch PDBs"
else
  fail "gloo SA cannot patch PDBs"
fi

echo ""
echo "=== Test 1: Apply GatewayParameters with HPA overlay ==="
kubectl apply -f - <<EOF
apiVersion: gateway.gloo.solo.io/v1alpha1
kind: GatewayParameters
metadata:
  name: $GWP_NAME
  namespace: $NAMESPACE
spec:
  kube:
    deployment:
      replicas: 1
    horizontalPodAutoscaler:
      metadata:
        labels:
          demo.hpa/source: initial
        annotations:
          demo.hpa/managed: "true"
      spec:
        minReplicas: 2
        maxReplicas: 5
        metrics:
        - type: Resource
          resource:
            name: cpu
            target:
              type: Utilization
              averageUtilization: 80
    envoyContainer:
      image:
        registry: quay.io/solo-io
        repository: gloo-envoy-wrapper
        tag: $VERSION
        pullPolicy: IfNotPresent
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
    service:
      type: ClusterIP
EOF
pass "Applied GatewayParameters ${GWP_NAME}"

if wait_for_jsonpath_value "gatewayparameters/${GWP_NAME}" '{.spec.kube.horizontalPodAutoscaler.spec.maxReplicas}' '5' 60; then
  pass "GatewayParameters maxReplicas starts at 5"
else
  fail "GatewayParameters maxReplicas did not start at 5"
fi

if wait_for_jsonpath_value "gatewayparameters/${GWP_NAME}" '{.spec.kube.horizontalPodAutoscaler.metadata.labels.demo\.hpa/source}' 'initial' 60; then
  pass "GatewayParameters overlay label starts at initial"
else
  fail "GatewayParameters overlay label did not start at initial"
fi

# Give controller caches a moment to observe the recreated GatewayParameters before Gateway creation.
sleep 3

echo ""
echo "=== Test 2: Create Gateway using that GatewayParameters ==="
kubectl apply -f - <<EOF
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: $GW_NAME
  namespace: $NAMESPACE
  annotations:
    gateway.gloo.solo.io/gateway-parameters-name: $GWP_NAME
spec:
  gatewayClassName: gloo-gateway
  listeners:
  - name: http
    port: 8080
    protocol: HTTP
    allowedRoutes:
      namespaces:
        from: Same
EOF
pass "Applied Gateway ${GW_NAME}"

echo ""
echo "=== Test 3: Verify Deployment and generated HPA exist ==="
if wait_for_exists "deployment/${DEPLOYMENT_NAME}" 120; then
  if kubectl rollout status "deployment/${DEPLOYMENT_NAME}" -n "$NAMESPACE" --timeout=180s >/dev/null; then
    pass "Deployment ${DEPLOYMENT_NAME} created and rolled out"
  else
    fail "Deployment ${DEPLOYMENT_NAME} rollout did not complete"
  fi
else
  fail "Deployment ${DEPLOYMENT_NAME} was not created"
fi

if wait_for_exists "hpa/${HPA_NAME}" 120; then
  pass "HPA ${HPA_NAME} was auto-created by the deployer"
else
  fail "HPA ${HPA_NAME} was not created"
fi

echo ""
echo "=== Test 4: Verify generated HPA wiring and overlay values ==="
if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.spec.scaleTargetRef.kind}' 'Deployment' 60; then
  pass "HPA scaleTargetRef.kind is Deployment"
else
  fail "HPA scaleTargetRef.kind is not Deployment"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.spec.scaleTargetRef.name}' "${DEPLOYMENT_NAME}" 60; then
  pass "HPA scaleTargetRef.name points to ${DEPLOYMENT_NAME}"
else
  fail "HPA scaleTargetRef.name does not point to ${DEPLOYMENT_NAME}"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.spec.minReplicas}' '2' 60; then
  pass "HPA minReplicas is 2"
else
  fail "HPA minReplicas is not 2"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.spec.maxReplicas}' '5' 60; then
  pass "HPA maxReplicas is 5"
else
  fail "HPA maxReplicas is not 5"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.metadata.labels.demo\.hpa/source}' 'initial' 60; then
  pass "HPA overlay label was applied"
else
  fail "HPA overlay label was not applied"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.metadata.annotations.demo\.hpa/managed}' 'true' 60; then
  pass "HPA overlay annotation was applied"
else
  fail "HPA overlay annotation was not applied"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.metadata.ownerReferences[0].kind}' 'Gateway' 60; then
  pass "HPA ownerReference kind is Gateway"
else
  fail "HPA ownerReference kind is not Gateway"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.metadata.ownerReferences[0].name}' "${GW_NAME}" 60; then
  pass "HPA ownerReference points to ${GW_NAME}"
else
  fail "HPA ownerReference does not point to ${GW_NAME}"
fi

echo ""
echo "=== Test 5: Update GatewayParameters overlay and verify HPA reconciles ==="
kubectl patch gatewayparameters "$GWP_NAME" -n "$NAMESPACE" --type merge -p '{
  "spec": {
    "kube": {
      "horizontalPodAutoscaler": {
        "metadata": {
          "labels": {
            "demo.hpa/source": "updated"
          }
        },
        "spec": {
          "maxReplicas": 7
        }
      }
    }
  }
}'

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.spec.maxReplicas}' '7' 90; then
  pass "HPA maxReplicas updated to 7 after GatewayParameters patch"
else
  fail "HPA maxReplicas did not update to 7"
fi

if wait_for_jsonpath_value "hpa/${HPA_NAME}" '{.metadata.labels.demo\.hpa/source}' 'updated' 90; then
  pass "HPA metadata labels reconciled after GatewayParameters patch"
else
  fail "HPA metadata labels did not reconcile after patch"
fi

echo ""
echo "=== Rendered HPA ==="
kubectl get hpa "$HPA_NAME" -n "$NAMESPACE" -o yaml

echo ""
echo "=== Summary ==="
kubectl get deployment "${DEPLOYMENT_NAME}" -n "$NAMESPACE"
kubectl get hpa "${HPA_NAME}" -n "$NAMESPACE"
echo ""
echo "Results: $PASS passed, $FAIL failed"
if [ "$FAIL" -gt 0 ]; then
  exit 1
fi
