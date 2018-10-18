package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/", basic)
        http.ListenAndServe(":8080", nil)
}

func basic(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Oi, vc requisitou: %s\n", r.URL.Path)
}
