package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, I'm running on %s with an %s CPU ",
		runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", indexHandler)
	hostAndPort := ":8080"
	// fix for run without firewall warning in window
	if os.Getenv("GOPRODUCT") != "1" {
		hostAndPort = "localhost:8080"
	}
	http.ListenAndServe(hostAndPort, nil)
}