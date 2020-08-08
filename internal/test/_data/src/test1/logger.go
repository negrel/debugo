package test1

import (
	// fmt is used in function parameters.
	"fmt"

	// Log & math package are only used in function body,
	// so it should be removed by Debuggo.
	"log"
	"math"
)

// Println calls Output to print to the standard
// logger. Arguments are handled in the manner of
// fmt.Println.
func Println(a ...fmt.Stringer) {
	log.Println(a)
}

func PrintPi() {
	log.Println(math.Pi)
}
