package main

import (
    "net/http"
)

func GetOnly(h handler) handler {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            h(w, r)
            return
        }
        http.Error(w, "GET requests only", http.StatusMethodNotAllowed)
    }
}

func PostOnly(h handler) handler {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            h(w, r)
            return
        }
        http.Error(w, "POST requests only", http.StatusMethodNotAllowed)
    }
}
