package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	args := struct {
		port           string
		backendPort    string
		backendBaseURL string
	}{}
	flag.StringVar(&args.port, "port", ":9091", "the port to serve the reverse proxy on")
	flag.StringVar(&args.backendPort, "proxy", "", "the port of the backend serve functioning as a proxy to")
	flag.StringVar(&args.backendBaseURL, "url", "fizbuz.biz", "the URL to serve as a proxy for")
	flag.Parse()

	/*
		TODO(aoeu): can't write to localhost or else errors happen:
		2019/10/11 03:40:31 http: Accept error: accept tcp [::]:9091: accept4: too many open files; retrying in 5ms
		2019/10/11 03:40:31 http: proxy error: dial tcp [::1]:9091: socket: too many open files
		502 Bad Gateway
	*/
	u, err := url.Parse(fmt.Sprintf("http://%s%s", args.backendBaseURL, args.backendPort))
	if err != nil {
		log.Fatal("error parsing URL", err)
	}

	revProxy := httputil.NewSingleHostReverseProxy(u)
	d := revProxy.Director
	revProxy.Director = func(r *http.Request) {
		d(r)
		r.Host = u.Host
	}

	if err := http.ListenAndServe(args.port, revProxy); err != nil {
		log.Fatal(err)
	}
}