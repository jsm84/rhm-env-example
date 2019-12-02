package controller

import (
	"github.com/jsm84/om-env-example/pkg/controller/ubinoop"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, ubinoop.Add)
}
