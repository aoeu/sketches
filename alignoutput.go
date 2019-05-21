package main

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

func main() {
	var sb strings.Builder

	w := tabwriter.NewWriter(&sb, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "aaaaaaaaaaaaaaaaaaaaaaaaaa\to\te\t")
	fmt.Fprintln(w, "aa\too\tee\t")
	fmt.Fprintln(w, "aaaa\toooo\teeee\tuuuu\t")
	fmt.Fprintln(w, "a\t")
	w.Flush()

	fmt.Println(sb.String())
}
