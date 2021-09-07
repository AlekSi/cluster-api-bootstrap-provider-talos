// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConfigFinalizer allows us to clean up resources before deletion
	ConfigFinalizer = "talosconfig.bootstrap.cluster.x-k8s.io"
)

// TalosConfigSpec defines the desired state of TalosConfig
type TalosConfigSpec struct {
	TalosVersion  string          `json:"talosVersion,omitempty"` //talos version formatted like v0.8. used for backwards compatibility
	GenerateType  string          `json:"generateType"`           //none,init,controlplane,worker mutually exclusive w/ data
	Data          string          `json:"data,omitempty"`
	ConfigPatches []ConfigPatches `json:"configPatches,omitempty"`
	// Important: Run "make" to regenerate code after modifying this file
}

// TalosConfigStatus defines the observed state of TalosConfig
type TalosConfigStatus struct {
	// Ready indicates the BootstrapData field is ready to be consumed
	Ready bool `json:"ready,omitempty"`

	// DataSecretName is the name of the secret that stores the bootstrap data script.
	// +optional
	DataSecretName *string `json:"dataSecretName,omitempty"`

	// Talos config will be a string containing the config for download
	// +optional
	TalosConfig string `json:"talosConfig,omitempty"`

	// FailureReason will be set on non-retryable errors
	// +optional
	FailureReason string `json:"failureReason,omitempty"`

	// FailureMessage will be set on non-retryable errors
	// +optional
	FailureMessage string `json:"failureMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=talosconfigs,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// TalosConfig is the Schema for the talosconfigs API
type TalosConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TalosConfigSpec   `json:"spec,omitempty"`
	Status TalosConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TalosConfigList contains a list of TalosConfig
type TalosConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TalosConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TalosConfig{}, &TalosConfigList{})
}
