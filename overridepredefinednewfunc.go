package main

import (
	"fmt"
)

func main() {
	overrideNewWithinFuncScope()
}

type printFunc func(string)()

// overrideNewWithinFuncScope uses the predefined func named "new" to declare
// a variable that itself overrides (shadows) the name of "new" func.
func overrideNewWithinFuncScope() {
	new := *(new(printFunc))
	new = func(s string) { fmt.Println(s) }
	new("int")
}
