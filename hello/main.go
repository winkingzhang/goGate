package main

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
	"math/rand"
	"log"
	"os"
	"fmt"
)

type Runtime struct {
	Name string    `json:"name"`
	OS   string    `json:"os"`
	Arch string    `json:"arch"`
	Rand int       `json:"_r"`
}

var r1 = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("httplog> %v %v %v (%q)", r.RemoteAddr, r.Method, r.Host, r.RequestURI)
		name, _ := os.Hostname()
		rt := Runtime{
			Name: name,
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
	fmt.Println("Server now listen on http://localhost:8081")
	http.ListenAndServe("localhost:8081", nil)
}