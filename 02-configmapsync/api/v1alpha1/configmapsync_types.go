/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigmapsyncSpec defines the desired state of Configmapsync.
type ConfigmapsyncSpec struct {
	SourceNamespace      string `json "sourceNamespace"`
	DestinationNamespace string `json "destinationNamespace"`
	ConfigmapName        string `json "configmapName"`
}

// ConfigmapsyncStatus defines the observed state of Configmapsync.
type ConfigmapsyncStatus struct {
	LastSyncTime metav1.Time `json:"lastSyncTime"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Configmapsync is the Schema for the configmapsyncs API.
type Configmapsync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigmapsyncSpec   `json:"spec,omitempty"`
	Status ConfigmapsyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigmapsyncList contains a list of Configmapsync.
type ConfigmapsyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configmapsync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Configmapsync{}, &ConfigmapsyncList{})
}
