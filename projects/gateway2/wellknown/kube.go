package wellknown

import (
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	DeploymentGVK              = appsv1.SchemeGroupVersion.WithKind("Deployment")
	ServiceGVK                 = corev1.SchemeGroupVersion.WithKind("Service")
	ServiceAccountGVK          = corev1.SchemeGroupVersion.WithKind("ServiceAccount")
	PodDisruptionBudgetGVK     = policyv1.SchemeGroupVersion.WithKind("PodDisruptionBudget")
	HorizontalPodAutoscalerGVK = autoscalingv2.SchemeGroupVersion.WithKind("HorizontalPodAutoscaler")
	VerticalPodAutoscalerGVK   = schema.GroupVersionKind{Group: "autoscaling.k8s.io", Version: "v1", Kind: "VerticalPodAutoscaler"}
)
