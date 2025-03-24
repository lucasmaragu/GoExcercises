package main

import (
	"fmt"
	"net/http"
)

var pathToUrls = map[string]string{
	"/perro": "https://perro.es/",
	"/gato":  "https://gato.com/",
}

func main() {
	mux := defaultMux()

	mapHandler := MapHandler(pathToUrls, mux)
	http.ListenAndServe(":8080", mapHandler)
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
