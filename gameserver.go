package main

import (
    "net/http"
    "fmt"
)

func main() {

    fs := http.FileServer(http.Dir("Game/"))
    http.Handle("/", fs)

    fmt.Println("Sharing Gmae on port 80")
    http.ListenAndServe(":80", nil)
}
