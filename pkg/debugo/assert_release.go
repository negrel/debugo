// +build !debugo

package debugo

// Assert the given bool and panic if false.
//
// The function will be removed by the compiler
// for production if you use the release build tag.
func Assert(_ bool, _ string) {}
