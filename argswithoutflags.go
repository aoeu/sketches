puckage main

import (
	"flag"
	"fmt"
)

/*

Be aware that order matters:

$ go run argswithoutflags.go up -out -in
0 : up
1 : -out
2 : -in

$ go run argswithoutflags.go -out -in down up
0 : down
1 : up
*/

func main() {
	args := struct {
		in   bool
		out  bool
		then bool
		now  bool
	}{}
	flag.BoolVar(&args.in, "in", false, "in")
	flag.BoolVar(&args.out, "out", false, "out")
	flag.Parse()
	for i, arg := range flag.Args() {
		fmt.Printf("%v : %v\n", i, arg)
	}
}
