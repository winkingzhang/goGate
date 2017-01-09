package main

import (
	"net/http"
	"net/http/httputil"
	"log"
	"fmt"
)

func main() {
	fmt.Println("Server now listen on http://localhost:8080")
	http.ListenAndServe(":8080", &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			log.Printf("httplog> %v %v %v (%q)", r.RemoteAddr, r.Method, r.Host, r.RequestURI)

			// set reverse proxy target
			r.URL.Scheme = "http"
			r.URL.Host = "localhost:8081"
			r.URL.Path = "/api/hello"

			if _, ok := r.Header["User-Agent"]; !ok {
				// explicitly disable User-Agent so it's not set to default value
				r.Header.Set("User-Agent", "")
			}

			// additional header
			r.Header.Set("X-Proxy-Secret", "Secret")
			r.Header.Set("Host", r.Host)
		},
	})
}