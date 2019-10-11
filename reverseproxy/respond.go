package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	args := struct {
		port string
	}{}
	flag.StringVar(&args.port, "port", ":9090", "the port to serve HTTP requests on")
	flag.Parse()

	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("hello, world\n")); err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/", handleFunc)

	if err := http.ListenAndServe(args.port, nil); err != nil {
		log.Fatal(err)
	}
}