package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"runtime"
)

type Runtime struct {
	OS   string    `json:"os"`
	Arch string   `json:"arch"`
}

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		rt := Runtime{
			OS: runtime.GOOS,
			Arch:runtime.GOARCH,
		};
		body, err := json.Marshal(rt)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(body));
	})
	hostAndPort := "localhost:8081"
	http.ListenAndServe(hostAndPort, nil)
}