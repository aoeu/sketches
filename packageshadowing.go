// This program is a contrived example of how a Go package identifier may
// be shadowed with a local variable that works as a facsimile for the Go
// pacakge itself. The program specifically takes advantage of the named
// wrapper functions of the net/http package, which wrap identically named
// methods of the net/http."Client" instance var net/http."DefaultClient",
// and underhandedly uses an instance of the program's own client type to
// masquerade as the http package.
//
// This approach is discouraged by the author, but shows how a "service"
// can be written and inserted as a placeholder for packages without funcs
// that would satisfy the concept of "a package as a service." Typically,
// this would 3rd party packages written without really understanding or
// examining many of the core concepts, idioms, and examples found within the
// Go standard library. As such, the "net/http" package is not an example of
// such a package, since it offers wrapper funcs over net/http."DefaultClient"
// such that the net/http package truly functions as a service.
//
// A better, recommended approach, would be to model the net/http package
// itself, where a "client" type is defined (and used as as service), but with
// a "DefaultClient" variable instance and wrapper funcs over that variable,
// all at the package-level, such that the package may be used as a service.
// However, package-shadowing may be used as a way of incrementally refactoring
// client code and a package toward the recommended approach, such that the
// package could be usable like a "package as a service."
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type client struct {
	*http.Client
	authToken string
}

func newClient(authToken string) *client {
	c := &client{
		authToken: authToken,
	}
	c.Client = &http.Client{
		Transport: c,
	}
	return c
}

func (c *client) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("masquerading as the http package to automatically add auth headers")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.authToken))
	return http.DefaultTransport.RoundTrip(req)
}

func main() {
	http := newClient("foo") // Shadow the http packgage's identifier with our own.
	resp, err := http.Get("http://fizbuz.biz/15")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
