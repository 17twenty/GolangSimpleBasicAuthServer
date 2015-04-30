package main

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello World!\n")
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "post request received\n")
}

type Result struct {
    FirstName string `json:"first"`
    LastName  string `json:"last"`
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    result, _ := json.Marshal(Result{"John", "Smith"})
    io.WriteString(w, string(result))
}
