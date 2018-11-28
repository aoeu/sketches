package main

import (
	"flag"
	"fmt"
	"os"
)

var usageMessage string = `
		This command-line-program is a tool to do the thing.
		
		This text prints out as help or to tell you that you're doing it wrong.
`

func main() {

	args := struct {
		everythingIsFine  bool
		contrivedArgument string
	}{}

	flag.BoolVar(&args.everythingIsFine, "ok", false, "A command line argument to specify everything is OK.")
	flag.StringVar(&args.contrivedArgument, "foo", "giberrish", "A contrived command land argument to supply nonsense with.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageMessage)
	}

	flag.Parse()
	// The Usage function will trigger automatically if things didn't parse correctly.

	// But we can also trigger the Usage function ourselves if we don't like some provided (or not provided) arguments.
	if !args.everythingIsFine {
		flag.Usage()
	}

}
