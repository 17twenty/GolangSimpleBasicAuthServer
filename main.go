package main

import (
    "log"
    "net/http"
)

func main() {
    // Setup routes and serve
    http.HandleFunc("/", HandleIndex)
    http.HandleFunc("/api/submit", PostOnly(BasicAuth(HandlePost)))
    http.HandleFunc("/api/json", GetOnly(BasicAuth(HandleJSON)))

    log.Fatal(http.ListenAndServe(":1380", nil))
}
