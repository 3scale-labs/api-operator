package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MAPPINGRULESpec defines the desired state of MAPPINGRULE
type MAPPINGRULESpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// MAPPINGRULEStatus defines the observed state of MAPPINGRULE
type MAPPINGRULEStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MAPPINGRULE is the Schema for the mappingrules API
// +k8s:openapi-gen=true
type MAPPINGRULE struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MAPPINGRULESpec   `json:"spec,omitempty"`
	Status MAPPINGRULEStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MAPPINGRULEList contains a list of MAPPINGRULE
type MAPPINGRULEList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MAPPINGRULE `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MAPPINGRULE{}, &MAPPINGRULEList{})
}
