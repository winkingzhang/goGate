package main

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
	"math/rand"
)

type Runtime struct {
	OS   string    `json:"os"`
	Arch string    `json:"arch"`
	Rand int       `json:"_r"`
}

var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		rt := Runtime{
			OS: runtime.GOOS,
			Arch:runtime.GOARCH,
			Rand: r1.Intn(100),
		};
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(rt); err != nil {
			panic(err)
		}
	})
	http.ListenAndServe("localhost:8081", nil)
}