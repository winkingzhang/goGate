package main

import (
	"net/http"
	"net/http/httputil"
	"log"
)

func main() {

	http.ListenAndServe(":8080", &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			log.Printf("httplog> %v %v %v (%q)", r.RemoteAddr, r.Method, r.Host, r.RequestURI)
			r.Header.Set("X-Proxy-Secret", "Secret")
			r.Header.Set("Host", r.Host)
			r.URL.Scheme = "http"
			r.URL.Host = "localhost:8081"
			r.RequestURI = ""
		},
	})
}