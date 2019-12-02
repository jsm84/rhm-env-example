package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UBINoOpSpec defines the desired state of UBINoOp
// +k8s:openapi-gen=true
type UBINoOpSpec struct {
	Size int32 `json:"size"`
}

// UBINoOpStatus defines the observed state of UBINoOp
// +k8s:openapi-gen=true
type UBINoOpStatus struct {
	Deployed bool `json:"deployed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UBINoOp is the Schema for the ubinoops API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ubinoops,scope=Namespaced
type UBINoOp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UBINoOpSpec   `json:"spec,omitempty"`
	Status UBINoOpStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UBINoOpList contains a list of UBINoOp
type UBINoOpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UBINoOp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UBINoOp{}, &UBINoOpList{})
}
