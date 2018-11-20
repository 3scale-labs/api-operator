package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// APPLICATIONPLANSpec defines the desired state of APPLICATIONPLAN
type APPLICATIONPLANSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// APPLICATIONPLANStatus defines the observed state of APPLICATIONPLAN
type APPLICATIONPLANStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APPLICATIONPLAN is the Schema for the applicationplans API
// +k8s:openapi-gen=true
type APPLICATIONPLAN struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APPLICATIONPLANSpec   `json:"spec,omitempty"`
	Status APPLICATIONPLANStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APPLICATIONPLANList contains a list of APPLICATIONPLAN
type APPLICATIONPLANList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APPLICATIONPLAN `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APPLICATIONPLAN{}, &APPLICATIONPLANList{})
}
