// Code generated by Debuggo. DO NOT EDIT.

//go:generate debuggo github.com/negrel/debuggo/examples/helloworld/internal/debug/debuggo/person
//+build !person

package person

const x = math.Pi

func Println(a ...interface{})     {}
func String(a fmt.Stringer) string {}
func PrintPi()                     {}
