package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    log.Println("Launch server on :8080")
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
