package main

import (
	"encoding/json"
	"net/http"
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(rt); err != nil {
			panic(err)
		}
	})
	hostAndPort := "localhost:8081"
	http.ListenAndServe(hostAndPort, nil)
}