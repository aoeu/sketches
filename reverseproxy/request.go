package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	args := struct {
		port string
	}{}
	flag.StringVar(&args.port, "port", ":9090", "the port to make an HTTP request to")
	flag.Parse()

	resp, err := http.Get(fmt.Sprintf("http://%s%s/1", "localhost", args.port))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status, string(b))
}