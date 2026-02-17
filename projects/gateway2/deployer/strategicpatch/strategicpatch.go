package strategicpatch

import (
	"encoding/json"
	"fmt"
	"maps"

	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gwv1alpha1 "github.com/solo-io/gloo/projects/gateway2/api/v1alpha1"
	"github.com/solo-io/gloo/projects/gateway2/wellknown"
)

// ResourceOverlays contains all overlays that can be applied to rendered objects.
type ResourceOverlays struct {
	Deployment              *gwv1alpha1.KubernetesResourceOverlay
	Service                 *gwv1alpha1.KubernetesResourceOverlay
	ServiceAccount          *gwv1alpha1.KubernetesResourceOverlay
	PodDisruptionBudget     *gwv1alpha1.KubernetesResourceOverlay
	HorizontalPodAutoscaler *gwv1alpha1.KubernetesResourceOverlay
	VerticalPodAutoscaler   *gwv1alpha1.KubernetesResourceOverlay
}

// FromGatewayParameters converts GatewayParameters overlays to generic ResourceOverlays.
func FromGatewayParameters(params *gwv1alpha1.GatewayParameters) *ResourceOverlays {
	if params == nil || params.Spec.Kube == nil {
		return nil
	}
	overlays := params.Spec.Kube.GatewayParametersOverlays
	return &ResourceOverlays{
		Deployment:              overlays.DeploymentOverlay,
		Service:                 overlays.ServiceOverlay,
		ServiceAccount:          overlays.ServiceAccountOverlay,
		PodDisruptionBudget:     overlays.PodDisruptionBudget,
		HorizontalPodAutoscaler: overlays.HorizontalPodAutoscaler,
		VerticalPodAutoscaler:   overlays.VerticalPodAutoscaler,
	}
}

// OverlayApplier applies overlays to rendered k8s objects using strategic merge patch semantics.
type OverlayApplier struct {
	overlays *ResourceOverlays
}

// NewOverlayApplierFromGatewayParameters creates a new OverlayApplier from GatewayParameters.
func NewOverlayApplierFromGatewayParameters(params *gwv1alpha1.GatewayParameters) *OverlayApplier {
	return &OverlayApplier{overlays: FromGatewayParameters(params)}
}

// NewOverlayApplierFromOverlays creates a new OverlayApplier from ResourceOverlays directly.
func NewOverlayApplierFromOverlays(overlays *ResourceOverlays) *OverlayApplier {
	return &OverlayApplier{overlays: overlays}
}

// ApplyOverlays applies overlays to rendered objects.
// It mutates objects in-place and may append new objects (PDB, HPA, VPA).
func (a *OverlayApplier) ApplyOverlays(objs []client.Object) ([]client.Object, error) {
	if a.overlays == nil {
		return objs, nil
	}

	// Find the Deployment first - we need it for PDB/HPA/VPA creation.
	var deployment *appsv1.Deployment
	for _, obj := range objs {
		if dep, ok := obj.(*appsv1.Deployment); ok {
			deployment = dep
			break
		}
	}

	for i, obj := range objs {
		var overlay *gwv1alpha1.KubernetesResourceOverlay
		var gvk schema.GroupVersionKind

		switch obj.(type) {
		case *appsv1.Deployment:
			overlay = a.overlays.Deployment
			gvk = wellknown.DeploymentGVK
		case *corev1.Service:
			overlay = a.overlays.Service
			gvk = wellknown.ServiceGVK
		case *corev1.ServiceAccount:
			overlay = a.overlays.ServiceAccount
			gvk = wellknown.ServiceAccountGVK
		default:
			continue
		}

		if overlay == nil {
			continue
		}

		patched, err := applyOverlay(obj, overlay, gvk)
		if err != nil {
			return nil, fmt.Errorf("failed to apply overlay to %s/%s: %w", gvk.Kind, obj.GetName(), err)
		}
		objs[i] = patched
	}

	if a.overlays.PodDisruptionBudget != nil && deployment != nil {
		idx := -1
		var existing *policyv1.PodDisruptionBudget
		for i, obj := range objs {
			if pdb, ok := obj.(*policyv1.PodDisruptionBudget); ok && pdb.Name == deployment.Name && pdb.Namespace == deployment.Namespace {
				idx = i
				existing = pdb
				break
			}
		}

		if existing != nil {
			patched, err := applyOverlay(existing, a.overlays.PodDisruptionBudget, wellknown.PodDisruptionBudgetGVK)
			if err != nil {
				return nil, fmt.Errorf("failed to apply overlay to existing PodDisruptionBudget: %w", err)
			}
			objs[idx] = patched
		} else {
			pdb, err := createPodDisruptionBudget(deployment, a.overlays.PodDisruptionBudget)
			if err != nil {
				return nil, fmt.Errorf("failed to create PodDisruptionBudget: %w", err)
			}
			objs = append(objs, pdb)
		}
	}

	if a.overlays.HorizontalPodAutoscaler != nil && deployment != nil {
		idx := -1
		var existing *autoscalingv2.HorizontalPodAutoscaler
		for i, obj := range objs {
			if hpa, ok := obj.(*autoscalingv2.HorizontalPodAutoscaler); ok && hpa.Name == deployment.Name && hpa.Namespace == deployment.Namespace {
				idx = i
				existing = hpa
				break
			}
		}

		if existing != nil {
			patched, err := applyOverlay(existing, a.overlays.HorizontalPodAutoscaler, wellknown.HorizontalPodAutoscalerGVK)
			if err != nil {
				return nil, fmt.Errorf("failed to apply overlay to existing HorizontalPodAutoscaler: %w", err)
			}
			objs[idx] = patched
		} else {
			hpa, err := createHorizontalPodAutoscaler(deployment, a.overlays.HorizontalPodAutoscaler)
			if err != nil {
				return nil, fmt.Errorf("failed to create HorizontalPodAutoscaler: %w", err)
			}
			objs = append(objs, hpa)
		}
	}

	if a.overlays.VerticalPodAutoscaler != nil && deployment != nil {
		idx := -1
		var existing *unstructured.Unstructured
		for i, obj := range objs {
			if vpa, ok := obj.(*unstructured.Unstructured); ok &&
				vpa.GetKind() == wellknown.VerticalPodAutoscalerGVK.Kind &&
				vpa.GetName() == deployment.Name &&
				vpa.GetNamespace() == deployment.Namespace {
				idx = i
				existing = vpa
				break
			}
		}

		if existing != nil {
			patched, err := applyVerticalPodAutoscalerOverlay(existing, a.overlays.VerticalPodAutoscaler)
			if err != nil {
				return nil, fmt.Errorf("failed to apply overlay to existing VerticalPodAutoscaler: %w", err)
			}
			objs[idx] = patched
		} else {
			vpa, err := createVerticalPodAutoscaler(deployment, a.overlays.VerticalPodAutoscaler)
			if err != nil {
				return nil, fmt.Errorf("failed to create VerticalPodAutoscaler: %w", err)
			}
			objs = append(objs, vpa)
		}
	}

	return objs, nil
}

func applyOverlay(obj client.Object, overlay *gwv1alpha1.KubernetesResourceOverlay, gvk schema.GroupVersionKind) (client.Object, error) {
	if overlay.Metadata != nil {
		if overlay.Metadata.Labels != nil {
			existingLabels := obj.GetLabels()
			if existingLabels == nil {
				existingLabels = make(map[string]string)
			}
			maps.Copy(existingLabels, overlay.Metadata.Labels)
			obj.SetLabels(existingLabels)
		}
		if overlay.Metadata.Annotations != nil {
			existingAnnotations := obj.GetAnnotations()
			if existingAnnotations == nil {
				existingAnnotations = make(map[string]string)
			}
			maps.Copy(existingAnnotations, overlay.Metadata.Annotations)
			obj.SetAnnotations(existingAnnotations)
		}
	}

	if overlay.Spec != nil && len(overlay.Spec.Raw) > 0 {
		return applySpecOverlay(obj, overlay.Spec.Raw, gvk)
	}

	return obj, nil
}

func applySpecOverlay(obj client.Object, patchBytes []byte, gvk schema.GroupVersionKind) (client.Object, error) {
	dataObj, err := getDataObjectForGVK(gvk)
	if err != nil {
		return nil, fmt.Errorf("unsupported kind %s for strategic merge patch: %w", gvk.Kind, err)
	}

	originalBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal original object: %w", err)
	}

	wrappedPatch := map[string]json.RawMessage{
		"spec": patchBytes,
	}
	wrappedPatchBytes, err := json.Marshal(wrappedPatch)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal wrapped patch: %w", err)
	}

	patchedBytes, err := strategicpatch.StrategicMergePatch(originalBytes, wrappedPatchBytes, dataObj)
	if err != nil {
		return nil, fmt.Errorf("failed to apply strategic merge patch: %w", err)
	}

	patchedObj, err := deserializeToObject(patchedBytes, gvk)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize patched object: %w", err)
	}

	return patchedObj, nil
}

func getDataObjectForGVK(gvk schema.GroupVersionKind) (runtime.Object, error) {
	switch gvk.Kind {
	case wellknown.DeploymentGVK.Kind:
		return &appsv1.Deployment{}, nil
	case wellknown.ServiceGVK.Kind:
		return &corev1.Service{}, nil
	case wellknown.ServiceAccountGVK.Kind:
		return &corev1.ServiceAccount{}, nil
	case wellknown.PodDisruptionBudgetGVK.Kind:
		return &policyv1.PodDisruptionBudget{}, nil
	case wellknown.HorizontalPodAutoscalerGVK.Kind:
		return &autoscalingv2.HorizontalPodAutoscaler{}, nil
	case wellknown.VerticalPodAutoscalerGVK.Kind:
		return &unstructured.Unstructured{}, nil
	default:
		return nil, fmt.Errorf("unsupported kind: %s", gvk.Kind)
	}
}

func deserializeToObject(data []byte, gvk schema.GroupVersionKind) (client.Object, error) {
	if gvk.Kind == wellknown.VerticalPodAutoscalerGVK.Kind {
		obj := &unstructured.Unstructured{}
		if err := json.Unmarshal(data, obj); err != nil {
			return nil, fmt.Errorf("failed to unmarshal patched object: %w", err)
		}
		obj.SetGroupVersionKind(gvk)
		return obj, nil
	}

	obj, err := getDataObjectForGVK(gvk)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, obj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal patched object: %w", err)
	}

	clientObj := obj.(client.Object)
	clientObj.GetObjectKind().SetGroupVersionKind(gvk)
	return clientObj, nil
}

func createPodDisruptionBudget(deployment *appsv1.Deployment, overlay *gwv1alpha1.KubernetesResourceOverlay) (client.Object, error) {
	pdb := &policyv1.PodDisruptionBudget{
		TypeMeta: metav1.TypeMeta{
			APIVersion: wellknown.PodDisruptionBudgetGVK.GroupVersion().String(),
			Kind:       wellknown.PodDisruptionBudgetGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
		},
		Spec: policyv1.PodDisruptionBudgetSpec{
			Selector: deployment.Spec.Selector,
		},
	}

	return applyOverlay(pdb, overlay, wellknown.PodDisruptionBudgetGVK)
}

func createHorizontalPodAutoscaler(deployment *appsv1.Deployment, overlay *gwv1alpha1.KubernetesResourceOverlay) (client.Object, error) {
	hpa := &autoscalingv2.HorizontalPodAutoscaler{
		TypeMeta: metav1.TypeMeta{
			APIVersion: wellknown.HorizontalPodAutoscalerGVK.GroupVersion().String(),
			Kind:       wellknown.HorizontalPodAutoscalerGVK.Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
		},
		Spec: autoscalingv2.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: autoscalingv2.CrossVersionObjectReference{
				APIVersion: wellknown.DeploymentGVK.GroupVersion().String(),
				Kind:       wellknown.DeploymentGVK.Kind,
				Name:       deployment.Name,
			},
		},
	}

	return applyOverlay(hpa, overlay, wellknown.HorizontalPodAutoscalerGVK)
}

func createVerticalPodAutoscaler(deployment *appsv1.Deployment, overlay *gwv1alpha1.KubernetesResourceOverlay) (client.Object, error) {
	vpa := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": wellknown.VerticalPodAutoscalerGVK.GroupVersion().String(),
			"kind":       wellknown.VerticalPodAutoscalerGVK.Kind,
			"metadata": map[string]any{
				"name":      deployment.Name,
				"namespace": deployment.Namespace,
			},
			"spec": map[string]any{
				"targetRef": map[string]any{
					"apiVersion": wellknown.DeploymentGVK.GroupVersion().String(),
					"kind":       wellknown.DeploymentGVK.Kind,
					"name":       deployment.Name,
				},
			},
		},
	}
	vpa.SetGroupVersionKind(wellknown.VerticalPodAutoscalerGVK)

	return applyVerticalPodAutoscalerOverlay(vpa, overlay)
}

func applyVerticalPodAutoscalerOverlay(vpa *unstructured.Unstructured, overlay *gwv1alpha1.KubernetesResourceOverlay) (*unstructured.Unstructured, error) {
	if overlay.Metadata != nil {
		if overlay.Metadata.Labels != nil {
			existingLabels := vpa.GetLabels()
			if existingLabels == nil {
				existingLabels = make(map[string]string)
			}
			maps.Copy(existingLabels, overlay.Metadata.Labels)
			vpa.SetLabels(existingLabels)
		}
		if overlay.Metadata.Annotations != nil {
			existingAnnotations := vpa.GetAnnotations()
			if existingAnnotations == nil {
				existingAnnotations = make(map[string]string)
			}
			maps.Copy(existingAnnotations, overlay.Metadata.Annotations)
			vpa.SetAnnotations(existingAnnotations)
		}
	}

	if overlay.Spec != nil && len(overlay.Spec.Raw) > 0 {
		var specPatch map[string]any
		if err := json.Unmarshal(overlay.Spec.Raw, &specPatch); err != nil {
			return nil, fmt.Errorf("failed to unmarshal spec patch: %w", err)
		}
		existingSpec, _, _ := unstructured.NestedMap(vpa.Object, "spec")
		if existingSpec == nil {
			existingSpec = make(map[string]any)
		}
		maps.Copy(existingSpec, specPatch)
		if err := unstructured.SetNestedMap(vpa.Object, existingSpec, "spec"); err != nil {
			return nil, fmt.Errorf("failed to set VPA spec: %w", err)
		}
	}

	return vpa, nil
}
