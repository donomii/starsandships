package main

import (
    "net/http"
)

func main() {

    fs := http.FileServer(http.Dir("Game/"))
    http.Handle("/", fs)

    http.ListenAndServe(":80", nil)
}
