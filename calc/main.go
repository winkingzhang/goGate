package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"strconv"
)

func main() {
	http.HandleFunc("/api/calc", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("httplog> %v %v %v (%q)", r.RemoteAddr, r.Method, r.Host, r.RequestURI)
		if (r.Method != http.MethodPost) {
			http.NotFound(w, r)
			return
		}

		decoder := json.NewDecoder(r.Body)
		data := struct {
			Operater string `json:"operator"`
			Left     string `json:"left"`
			Right    string `json:"right"`
			Result   string `json:"result"`
		}{}
		err := decoder.Decode(&data);
		if (err != nil) {
			panic(err);
		}
		defer r.Body.Close()
		left, _ := strconv.Atoi(data.Left)
		right, _ := strconv.Atoi(data.Right)

		switch data.Operater {
		case "+": data.Result = strconv.Itoa(left + right); break;
		case "-": data.Result = strconv.Itoa(left - right); break;
		case "*": data.Result = strconv.Itoa(left * right); break;
		case "/": data.Result = strconv.Itoa(left / right); break;
		case "%": data.Result = strconv.Itoa(left % right); break;
		default: json.NewEncoder(w).Encode(struct{ Error string }{Error: "Not recorgonized operation"}); return;
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	})
	fmt.Println("Server now listen on http://localhost:8082")
	http.ListenAndServe("localhost:8082", nil)
}