// Code generated by Debuggo. DO NOT EDIT.

// +build !test1

package test1

import (
	"fmt"
)

// fmt is used in function parameters.

// Log package is only used in function body,
// so it should be removed by debuggo.

// Println calls Output to print to the standard
// logger. Arguments are handled in the manner of
// fmt.Println.
func Println(a ...fmt.Stringer) {

}
