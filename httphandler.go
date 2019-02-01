// httphandler demonstrates how to create a basic HTTP server in Go
// for educational purposes, leveraging the standard library packages.
//
// Functions are commented with proper 'go doc' comments for the sake
// of explanation.
//
// Functions are intentionally limited to one level of abstraction, both for
// the sake of self-documenting intention with the function name itself,
// and to highlight the variable types in use from the standard library.
//
// Usage:
// 		httphandler
//
// Examples:
//		$ go run httphandler.go
//		$ curl 'localhost:9001/hello'
//		Hello, world!
//		$ curl 'localhost:9001/hello?name=foo'
//		Hello, foo!
//
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// main runs a default HTTP server on port 9001 and servers a greeting
// to HTTP requests made to the URI "hello".
func main() {
	http.HandleFunc("/hello", sayHello)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatalf("error when attempting to listen and serve HTTP: %v", err)
	}
}

// sayHello writes a greeting as an HTTP response to an HTTP request.
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write(createGreeting(r.URL.Query()))
}

// createGreeting creates a greeting message using a named URL query
// parameter or else returns a default greeting if the query parameter
// or its value are ommitted.
func createGreeting(v url.Values) []byte {
	s := fmt.Sprintf("Hello, %v!", lookupParameter(v, "name", "world"))
	return []byte(s)
}

// lookupParameter determines if a URL's query parameter or form parameter
// values have a parameter with the given name and, if so, return it,
// otherwise the provided fallback value is returned.
func lookupParameter(v url.Values, name, fallback string) (value string) {
	if s, ok := v[name]; ok && len(s) == 1 && s[0] != "" {
		value = s[0]
	} else {
		value = fallback
	}
	return value
}
