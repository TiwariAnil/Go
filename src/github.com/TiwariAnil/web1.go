package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func data_base(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let me check if its working"))
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/db", data_base)
    http.ListenAndServe(":8080", nil)
}