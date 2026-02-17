package v1alpha1

import apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// ObjectMetadata contains labels and annotations for metadata overlays.
type ObjectMetadata struct {
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// KubernetesResourceOverlay customizes rendered Kubernetes resources.
// The spec field is applied using strategic merge patch semantics.
type KubernetesResourceOverlay struct {
	// +optional
	Metadata *ObjectMetadata `json:"metadata,omitempty"`

	// +optional
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields
	Spec *apiextensionsv1.JSON `json:"spec,omitempty"`
}
