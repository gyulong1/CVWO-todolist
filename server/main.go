package main

import (
    "fmt"
    "log"
    "net/http"
    _ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
    r := router()
    fmt.Println("Starting server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}