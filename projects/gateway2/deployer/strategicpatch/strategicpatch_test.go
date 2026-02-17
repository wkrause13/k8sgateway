package strategicpatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	gwv1alpha1 "github.com/solo-io/gloo/projects/gateway2/api/v1alpha1"
)

func jsonPatch(raw string) *apiextensionsv1.JSON {
	return &apiextensionsv1.JSON{Raw: []byte(raw)}
}

func TestOverlayApplier_ApplyOverlays_NilGatewayParameters(t *testing.T) {
	applier := NewOverlayApplierFromGatewayParameters(nil)
	objs := []client.Object{
		&appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
				Kind:       "Deployment",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-deployment",
			},
		},
	}

	result, err := applier.ApplyOverlays(objs)
	require.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestOverlayApplier_ApplyOverlays_DeploymentMetadataAndSpec(t *testing.T) {
	params := &gwv1alpha1.GatewayParameters{
		Spec: gwv1alpha1.GatewayParametersSpec{
			Kube: &gwv1alpha1.KubernetesProxyConfig{
				GatewayParametersOverlays: gwv1alpha1.GatewayParametersOverlays{
					DeploymentOverlay: &gwv1alpha1.KubernetesResourceOverlay{
						Metadata: &gwv1alpha1.ObjectMetadata{
							Labels: map[string]string{
								"custom-label": "custom-value",
							},
							Annotations: map[string]string{
								"custom-annotation": "custom-value",
							},
						},
						Spec: jsonPatch(`{
							"replicas": 3,
							"template": {
								"spec": {
									"containers": [{
										"name": "gateway-proxy",
										"resources": {
											"limits": {
												"memory": "512Mi"
											}
										}
									}]
								}
							}
						}`),
					},
				},
			},
		},
	}

	applier := NewOverlayApplierFromGatewayParameters(params)
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-deployment",
			Namespace: "default",
			Labels: map[string]string{
				"existing-label": "existing-value",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.To[int32](1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "test"},
			},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "gateway-proxy",
							Image: "gloo/gateway:latest",
						},
					},
				},
			},
		},
	}
	objs := []client.Object{deployment}

	objs, err := applier.ApplyOverlays(objs)
	require.NoError(t, err)

	result := objs[0].(*appsv1.Deployment)
	assert.Equal(t, int32(3), *result.Spec.Replicas)
	assert.Equal(t, "custom-value", result.Labels["custom-label"])
	assert.Equal(t, "existing-value", result.Labels["existing-label"])
	assert.Equal(t, "custom-value", result.Annotations["custom-annotation"])
	assert.Equal(t, "gloo/gateway:latest", result.Spec.Template.Spec.Containers[0].Image)
	assert.Equal(t, "512Mi", result.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().String())
}

func TestOverlayApplier_ApplyOverlays_CreatesPDBHpaVPA(t *testing.T) {
	params := &gwv1alpha1.GatewayParameters{
		Spec: gwv1alpha1.GatewayParametersSpec{
			Kube: &gwv1alpha1.KubernetesProxyConfig{
				GatewayParametersOverlays: gwv1alpha1.GatewayParametersOverlays{
					PodDisruptionBudget: &gwv1alpha1.KubernetesResourceOverlay{
						Metadata: &gwv1alpha1.ObjectMetadata{
							Labels: map[string]string{
								"resource-type": "pdb",
							},
						},
						Spec: jsonPatch(`{"minAvailable": 1}`),
					},
					HorizontalPodAutoscaler: &gwv1alpha1.KubernetesResourceOverlay{
						Metadata: &gwv1alpha1.ObjectMetadata{
							Labels: map[string]string{
								"resource-type": "hpa",
							},
						},
						Spec: jsonPatch(`{
							"minReplicas": 2,
							"maxReplicas": 10,
							"metrics": [{
								"type": "Resource",
								"resource": {
									"name": "cpu",
									"target": {
										"type": "Utilization",
										"averageUtilization": 80
									}
								}
							}]
						}`),
					},
					VerticalPodAutoscaler: &gwv1alpha1.KubernetesResourceOverlay{
						Metadata: &gwv1alpha1.ObjectMetadata{
							Labels: map[string]string{
								"resource-type": "vpa",
							},
						},
						Spec: jsonPatch(`{
							"updatePolicy": {
								"updateMode": "Auto"
							}
						}`),
					},
				},
			},
		},
	}

	applier := NewOverlayApplierFromGatewayParameters(params)
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-deployment",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "test"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "test"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "gateway-proxy",
							Image: "gloo/gateway:latest",
						},
					},
				},
			},
		},
	}
	objs := []client.Object{deployment}

	objs, err := applier.ApplyOverlays(objs)
	require.NoError(t, err)
	require.Len(t, objs, 4)

	var pdb *policyv1.PodDisruptionBudget
	var hpa *autoscalingv2.HorizontalPodAutoscaler
	var vpa *unstructured.Unstructured

	for _, obj := range objs {
		switch v := obj.(type) {
		case *policyv1.PodDisruptionBudget:
			pdb = v
		case *autoscalingv2.HorizontalPodAutoscaler:
			hpa = v
		case *unstructured.Unstructured:
			if v.GetKind() == "VerticalPodAutoscaler" {
				vpa = v
			}
		}
	}

	require.NotNil(t, pdb)
	require.NotNil(t, hpa)
	require.NotNil(t, vpa)

	assert.Equal(t, "pdb", pdb.Labels["resource-type"])
	require.NotNil(t, pdb.Spec.MinAvailable)
	assert.Equal(t, int32(1), pdb.Spec.MinAvailable.IntVal)
	assert.Equal(t, "test", pdb.Spec.Selector.MatchLabels["app"])

	assert.Equal(t, "hpa", hpa.Labels["resource-type"])
	assert.Equal(t, int32(2), ptr.Deref(hpa.Spec.MinReplicas, 0))
	assert.Equal(t, int32(10), hpa.Spec.MaxReplicas)
	require.NotEmpty(t, hpa.Spec.Metrics)
	assert.Equal(t, "cpu", string(hpa.Spec.Metrics[0].Resource.Name))
	assert.Equal(t, int32(80), ptr.Deref(hpa.Spec.Metrics[0].Resource.Target.AverageUtilization, 0))
	assert.Equal(t, "test-deployment", hpa.Spec.ScaleTargetRef.Name)

	assert.Equal(t, "vpa", vpa.GetLabels()["resource-type"])
	targetRef, found, err := unstructured.NestedMap(vpa.Object, "spec", "targetRef")
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, "Deployment", targetRef["kind"])
	assert.Equal(t, "test-deployment", targetRef["name"])

	updatePolicy, found, err := unstructured.NestedMap(vpa.Object, "spec", "updatePolicy")
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, "Auto", updatePolicy["updateMode"])
}

func TestOverlayApplier_ApplyOverlays_MergesCreatedAutoscalerResourcesAcrossPasses(t *testing.T) {
	gatewayClassApplier := NewOverlayApplierFromOverlays(&ResourceOverlays{
		PodDisruptionBudget: &gwv1alpha1.KubernetesResourceOverlay{
			Metadata: &gwv1alpha1.ObjectMetadata{
				Labels: map[string]string{
					"shared":   "gatewayclass",
					"gwc-only": "true",
				},
			},
			Spec: jsonPatch(`{"minAvailable": 1}`),
		},
		HorizontalPodAutoscaler: &gwv1alpha1.KubernetesResourceOverlay{
			Metadata: &gwv1alpha1.ObjectMetadata{
				Labels: map[string]string{
					"shared":   "gatewayclass",
					"gwc-only": "true",
				},
			},
			Spec: jsonPatch(`{"minReplicas": 1, "maxReplicas": 5}`),
		},
		VerticalPodAutoscaler: &gwv1alpha1.KubernetesResourceOverlay{
			Metadata: &gwv1alpha1.ObjectMetadata{
				Labels: map[string]string{
					"shared":   "gatewayclass",
					"gwc-only": "true",
				},
			},
			Spec: jsonPatch(`{"updatePolicy": {"updateMode": "Off"}}`),
		},
	})

	gatewayApplier := NewOverlayApplierFromOverlays(&ResourceOverlays{
		PodDisruptionBudget: &gwv1alpha1.KubernetesResourceOverlay{
			Metadata: &gwv1alpha1.ObjectMetadata{
				Labels: map[string]string{
					"shared":  "gateway",
					"gw-only": "true",
				},
			},
			Spec: jsonPatch(`{"minAvailable": 2}`),
		},
		HorizontalPodAutoscaler: &gwv1alpha1.KubernetesResourceOverlay{
			Metadata: &gwv1alpha1.ObjectMetadata{
				Labels: map[string]string{
					"shared":  "gateway",
					"gw-only": "true",
				},
			},
			Spec: jsonPatch(`{"minReplicas": 2, "maxReplicas": 10}`),
		},
		VerticalPodAutoscaler: &gwv1alpha1.KubernetesResourceOverlay{
			Metadata: &gwv1alpha1.ObjectMetadata{
				Labels: map[string]string{
					"shared":  "gateway",
					"gw-only": "true",
				},
			},
			Spec: jsonPatch(`{"updatePolicy": {"updateMode": "Auto"}}`),
		},
	})

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-deployment",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "test"},
			},
		},
	}

	objs, err := gatewayClassApplier.ApplyOverlays([]client.Object{deployment})
	require.NoError(t, err)
	objs, err = gatewayApplier.ApplyOverlays(objs)
	require.NoError(t, err)

	// Deployment + PDB + HPA + VPA (no duplicates from the second pass).
	require.Len(t, objs, 4)

	var pdb *policyv1.PodDisruptionBudget
	var hpa *autoscalingv2.HorizontalPodAutoscaler
	var vpa *unstructured.Unstructured
	for _, obj := range objs {
		switch v := obj.(type) {
		case *policyv1.PodDisruptionBudget:
			pdb = v
		case *autoscalingv2.HorizontalPodAutoscaler:
			hpa = v
		case *unstructured.Unstructured:
			if v.GetKind() == "VerticalPodAutoscaler" {
				vpa = v
			}
		}
	}

	require.NotNil(t, pdb)
	require.NotNil(t, hpa)
	require.NotNil(t, vpa)

	assert.Equal(t, "gateway", pdb.Labels["shared"])
	assert.Equal(t, "true", pdb.Labels["gwc-only"])
	assert.Equal(t, "true", pdb.Labels["gw-only"])
	require.NotNil(t, pdb.Spec.MinAvailable)
	assert.Equal(t, int32(2), pdb.Spec.MinAvailable.IntVal)

	assert.Equal(t, "gateway", hpa.Labels["shared"])
	assert.Equal(t, "true", hpa.Labels["gwc-only"])
	assert.Equal(t, "true", hpa.Labels["gw-only"])
	assert.Equal(t, int32(2), ptr.Deref(hpa.Spec.MinReplicas, 0))
	assert.Equal(t, int32(10), hpa.Spec.MaxReplicas)

	assert.Equal(t, "gateway", vpa.GetLabels()["shared"])
	assert.Equal(t, "true", vpa.GetLabels()["gwc-only"])
	assert.Equal(t, "true", vpa.GetLabels()["gw-only"])
	updatePolicy, found, err := unstructured.NestedMap(vpa.Object, "spec", "updatePolicy")
	require.NoError(t, err)
	require.True(t, found)
	assert.Equal(t, "Auto", updatePolicy["updateMode"])
}
