package main

import (
	"html/template"
	"net/http"
	"net/http/httputil"
	"log"
	"fmt"
)

type ProxyMap struct {
	Schema  string              `json:"schema"`
	Host    string              `json:"host"`
	Path    string              `json:"path"`
	Headers *map[string]string  `json:"headers"`
}

func NewProxyMap(schema, host, path string, headers *map[string]string) *ProxyMap {
	return &ProxyMap{
		Schema: schema,
		Host: host,
		Path: path,
		Headers: headers,
	}
}

type Site struct {
	Backend      *ProxyMap
	reverseProxy *httputil.ReverseProxy
}

var sites = make(map[string]*Site)

func (s *Site) writePreLogging(r *http.Request) {
	log.Printf("[reverse proxy] %v %v %v (%q) => [backend] %v://%v%v",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI,
		s.Backend.Schema, s.Backend.Host, s.Backend.Path)
}

func (s *Site) getReverseProxy() *httputil.ReverseProxy {
	if s.reverseProxy != nil {
		return s.reverseProxy
	}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			// set reverse proxy target
			r.URL.Scheme = s.Backend.Schema
			r.URL.Host = s.Backend.Host
			r.URL.Path = s.Backend.Path

			if _, ok := r.Header["User-Agent"]; !ok {
				// explicitly disable User-Agent so it's not set to default value
				r.Header.Set("User-Agent", "")
			}

			// additional header
			//r.Header.Set("X-Proxy-Secret", "Secret")
			//r.Header.Set("Host", r.Host)
			if s.Backend.Headers != nil {
				for h, v := range *s.Backend.Headers {
					r.Header.Set(h, v);
				}
			}

			// prepare ok, write log and send request to backend
			go s.writePreLogging(r)
		},
	}
}

func reverseProxyHandle(w http.ResponseWriter, r *http.Request) {
	// parse requst url and create reverse proxy
	site := sites[r.RequestURI]
	if site != nil {
		// pass to reversed proxy
		site.getReverseProxy().ServeHTTP(w, r);

		//go writePostLogging(w, r);
	} else {
		//make 404 error
		http.NotFound(w, r);
	}
}

func everythingHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[everythingHandle] reserved proxy > %v %v %v (%q)",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<div><strong>Hello golang, this is {{.Name}}</strong></div>
	</body>
</html>`
		t, err := template.New("webpage").Parse(tpl)
		if (err != nil) {
			panic(err)
		}
		data := struct {
			Title string
			Name  string
		}{
			Title: "API Gateway",
			Name: "Gatway Home",
		}
		err = t.Execute(w, data);
		if (err != nil) {
			panic(err)
		}
		return
	default:
		http.NotFound(w, r)
		return
	}
}

func main() {
	sites["/api/hello/v1"] = &Site{
		Backend: NewProxyMap("http", "localhost:8081", "/api/hello", nil),
	}
	sites["/api/calc/v1"] = &Site{
		Backend: NewProxyMap("http", "localhost:8082", "/api/calc", nil),
	}
	http.HandleFunc("/", everythingHandle)
	http.HandleFunc("/api/hello/v1", reverseProxyHandle)
	http.HandleFunc("/api/calc/v1", reverseProxyHandle)

	fmt.Println("Server now listen on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}