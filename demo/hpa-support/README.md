# HPA Overlay Demo

This demo proves `GatewayParameters.spec.kube.horizontalPodAutoscaler` works end-to-end.

## What it validates

1. `GatewayParameters` accepts `horizontalPodAutoscaler` overlay input.
2. The deployer creates an `autoscaling/v2` `HorizontalPodAutoscaler` automatically.
3. `spec.scaleTargetRef` points to the generated `gloo-proxy-<gateway>` deployment.
4. Overlay metadata/spec values are applied to the generated HPA.
5. Patching `GatewayParameters` reconciles the generated HPA.

## Run

```bash
cd demo/hpa-support
./01-setup-cluster.sh
./02-demo-hpa-overlay.sh
```

## Cleanup

```bash
cd demo/hpa-support
./03-cleanup.sh
# or remove cluster too:
./03-cleanup.sh --all
```
