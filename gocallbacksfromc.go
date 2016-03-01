package main

import (
	"fmt"
	"os"
)

/*
extern void go_callback_dispatcher(int callbackID);

// Ideally, C functions or variables would be defined
// in another separate C file to avoid the multiple definition
// errors, however, using "static inline" is a nice workaround
// for simple functions like this one.
static inline void CallMyGoFunction(int callbackID) {
    go_callback_dispatcher(callbackID);
}
*/
import "C"

//export go_callback_dispatcher
func go_callback_dispatcher(id C.int) {
	callback, ok := callbacks[id]
	if !ok {
		fmt.Fprintln(os.Stderr, "Could not find callback function with ID %v\n", id)
		return
	}
	callback(id)
}

func PrintCallbackIDFromGo(id C.int) {
	fmt.Printf("PrintCallbackIDFromGo was called via ID: %v\n", id)
}

// Store the Go callback function in a global map because an unsafe.Pointer
// in Go 1.6 can no longer be passed to and called by C. This is because the Go
// garbage collector may move where Go values are located in memory, but the C
// code with a stored pointer value to the Go function won't update.
//
// This was not necessary in Go 1.5, but the safety concerns are discussed in
// Ian Taylor's design proposal for adding rules for passing pointers between
// C and Go:
// - https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md
//
// A filed issue outlines the implementation change:
// - https://github.com/golang/go/issues/12416
//
// rsc outlines the problems of passing Go Pointers to C:
// -  https://github.com/golang/go/issues/8310
//
// An unsafe.Pointer can still be used to directly call a Go function
// from C by setting GODEBUG=cgocheck=0
// - https://github.com/golang/go/issues/12416#issuecomment-161850713
//
// Keep in mind that those who give up correctness (type-safety)
// for temporary convenience may not have sufficiently felt the
// inevitable inconvenience of incorrectness.
var callbacks = make(map[C.int]func(C.int))

func main() {
	var callbackID C.int = 777
	callbacks[callbackID] = PrintCallbackIDFromGo
	C.CallMyGoFunction(C.int(callbackID))
}
