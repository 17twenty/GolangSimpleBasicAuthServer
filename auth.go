package main

import (
    "encoding/base64"
    "net/http"
    "strings"
)

type handler func(w http.ResponseWriter, r *http.Request)

func BasicAuth(pass handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {

        auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

        if len(auth) != 2 {
            w.Header().Set("WWW-Authenticate", "basic realm=\"Please enter your credentials\"")
            http.Error(w, "Authorisation Required", http.StatusUnauthorized)
            return
        }

        payload, _ := base64.StdEncoding.DecodeString(auth[1])
        pair := strings.SplitN(string(payload), ":", 2)
        if len(pair) != 2 || !Validate(pair[0], pair[1]) {
            http.Error(w, "authorization failed", http.StatusUnauthorized)
            return
        }

        pass(w, r)
    }
}

func Validate(username, password string) bool {
    if username == "foo" && password == "bar" {
        return true
    }
    return false
}
