package controller

import (
	"github.com/3scale/api-operator/pkg/controller/mappingrule"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mappingrule.Add)
}
